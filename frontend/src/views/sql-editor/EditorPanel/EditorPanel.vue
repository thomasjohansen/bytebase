<template>
  <div class="flex h-full w-full flex-col justify-start items-start">
    <template v-if="tab.editMode === 'SQL-EDITOR'">
      <EditorAction @execute="handleExecute" />

      <ConnectionPathBar />

      <Suspense>
        <SQLEditor @execute="handleExecute" />
        <template #fallback>
          <div
            class="w-full h-auto flex-grow flex flex-col items-center justify-center"
          >
            <BBSpin />
          </div>
        </template>
      </Suspense>
    </template>

    <Suspense>
      <AIChatToSQL
        v-if="!tabStore.isDisconnected && showAIChatBox"
        :allow-config="pageMode === 'BUNDLED'"
        @apply-statement="handleApplyStatement"
      />
    </Suspense>

    <ExecutingHintModal />

    <SaveSheetModal />
  </div>
</template>

<script lang="ts" setup>
import { defineAsyncComponent } from "vue";
import { useExecuteSQL } from "@/composables/useExecuteSQL";
import { AIChatToSQL } from "@/plugins/ai";
import {
  useCurrentTab,
  useInstanceV1Store,
  useTabStore,
  usePageMode,
} from "@/store";
import type { Connection, ExecuteConfig, ExecuteOption } from "@/types";
import { formatEngineV1 } from "@/utils";
import {
  EditorAction,
  ConnectionPathBar,
  ExecutingHintModal,
  SaveSheetModal,
} from "../EditorCommon";
import { useSQLEditorContext } from "../context";

const SQLEditor = defineAsyncComponent(() => import("./SQLEditor.vue"));

const tabStore = useTabStore();
const tab = useCurrentTab();
const { showAIChatBox } = useSQLEditorContext();
const pageMode = usePageMode();

const { executeReadonly } = useExecuteSQL();

const handleExecute = (
  query: string,
  config: ExecuteConfig,
  option?: ExecuteOption
) => {
  executeReadonly(query, config, option);
};

// const trySaveSheet = (sheetName?: string) => {
//   saveSheetModal.value?.trySaveSheet(sheetName);
// };

const handleApplyStatement = async (
  statement: string,
  conn: Connection,
  run: boolean
) => {
  tab.value.statement = statement;
  if (run) {
    const instanceStore = useInstanceV1Store();
    const instance = await instanceStore.getOrFetchInstanceByUID(
      conn.instanceId
    );
    handleExecute(statement, {
      databaseType: formatEngineV1(instance),
    });
  }
};
</script>
