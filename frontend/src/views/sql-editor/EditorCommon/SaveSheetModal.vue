<template>
  <BBModal
    v-if="state.showModal"
    :title="$t('sql-editor.save-sheet')"
    @close="close"
  >
    <SaveSheetForm @close="close" @confirm="doSaveSheet" />
  </BBModal>
</template>

<script lang="ts" setup>
import { computed, reactive } from "vue";
import { useEmitteryEventListener } from "@/composables/useEmitteryEventListener";
import { useDatabaseV1Store, useWorkSheetStore, useTabStore } from "@/store";
import { UNKNOWN_ID } from "@/types";
import {
  Worksheet,
  Worksheet_Visibility,
} from "@/types/proto/v1/worksheet_service";
import { extractWorksheetUID } from "@/utils";
import { useSheetContext } from "../Sheet";
import { useSQLEditorContext } from "../context";
import SaveSheetForm from "./SaveSheetForm.vue";

type LocalState = {
  showModal: boolean;
};

const tabStore = useTabStore();
const databaseStore = useDatabaseV1Store();
const worksheetV1Store = useWorkSheetStore();
const { events: sheetEvents } = useSheetContext();
const { events: editorEvents } = useSQLEditorContext();

const state = reactive<LocalState>({
  showModal: false,
});

const allowSave = computed((): boolean => {
  const tab = tabStore.currentTab;
  if (tab.statement === "") {
    return false;
  }
  if (tab.isSaved) {
    return false;
  }
  // Temporarily disable saving and sharing if we are connected to an instance
  // but not a database.
  if (tab.connection.databaseId === String(UNKNOWN_ID)) {
    return false;
  }
  return true;
});

const doSaveSheet = async (title: string) => {
  const { name, statement, sheetName } = tabStore.currentTab;
  title = title || name;

  const sheetId = Number(extractWorksheetUID(sheetName ?? ""));

  const conn = tabStore.currentTab.connection;
  const database = await databaseStore.getOrFetchDatabaseByUID(
    conn.databaseId,
    true /* silent */
  );

  let sheet: Worksheet | undefined;
  if (sheetId !== UNKNOWN_ID) {
    sheet = await worksheetV1Store.patchSheet(
      {
        name: sheetName,
        title,
        content: new TextEncoder().encode(statement),
      },
      ["title", "content"]
    );
  } else {
    sheet = await worksheetV1Store.createSheet(
      Worksheet.fromPartial({
        title: title,
        project: database.project,
        content: new TextEncoder().encode(statement),
        database: database.name,
        visibility: Worksheet_Visibility.VISIBILITY_PRIVATE,
      })
    );
  }

  if (sheet) {
    tabStore.updateCurrentTab({
      sheetName: sheet.name,
      isSaved: true,
      name: title,
    });

    // Refresh "my" sheet list.
    sheetEvents.emit("refresh", { views: ["my"] });
  }
  state.showModal = false;
};

const needSheetTitle = (title: string) => {
  const tab = tabStore.currentTab;
  if (tab.sheetName) {
    // If the sheet is saved, we don't need to show the name popup.
    return false;
  }
  return true;
};

const trySaveSheet = (title: string) => {
  if (!allowSave.value) {
    return;
  }

  if (needSheetTitle(title)) {
    state.showModal = true;
    return;
  }
  state.showModal = false;

  doSaveSheet(title);
};

const close = () => {
  state.showModal = false;
};

useEmitteryEventListener(editorEvents, "save-sheet", ({ title }) => {
  trySaveSheet(title);
});
</script>
