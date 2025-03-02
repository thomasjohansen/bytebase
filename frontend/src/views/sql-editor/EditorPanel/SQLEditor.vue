<template>
  <div
    class="w-full h-auto flex-grow flex flex-col justify-start items-start overflow-auto"
  >
    <MonacoEditor
      class="w-full h-full"
      :filename="filename"
      :content="content"
      :language="language"
      :dialect="dialect"
      :readonly="readonly"
      :advices="advices"
      :auto-complete-context="{
        instance: instance.name,
        database: database.name,
      }"
      @update:content="handleChange"
      @select-content="handleChangeSelection"
      @ready="handleEditorReady"
    />
  </div>
</template>

<script lang="ts" setup>
import { v1 as uuidv1 } from "uuid";
import { computed, nextTick, ref, watch } from "vue";
import type {
  AdviceOption,
  IStandaloneCodeEditor,
  MonacoModule,
} from "@/components/MonacoEditor";
import MonacoEditor from "@/components/MonacoEditor/MonacoEditor.vue";
import {
  extensionNameOfLanguage,
  formatEditorContent,
} from "@/components/MonacoEditor/utils";
import {
  useTabStore,
  useSQLEditorStore,
  useUIStateStore,
  useInstanceV1ByUID,
  useWorkSheetAndTabStore,
  useDatabaseV1ByUID,
} from "@/store";
import {
  dialectOfEngineV1,
  ExecuteConfig,
  ExecuteOption,
  SQLDialect,
} from "@/types";
import { formatEngineV1, useInstanceV1EditorLanguage } from "@/utils";
import { useSQLEditorContext } from "../context";

const emit = defineEmits<{
  (
    e: "execute",
    sql: string,
    config: ExecuteConfig,
    option?: ExecuteOption
  ): void;
}>();

const tabStore = useTabStore();
const sqlEditorStore = useSQLEditorStore();
const sheetAndTabStore = useWorkSheetAndTabStore();
const uiStateStore = useUIStateStore();
const { events: editorEvents } = useSQLEditorContext();

const content = computed(() => tabStore.currentTab.statement);
const advices = computed((): AdviceOption[] => {
  return (
    Array.from(tabStore.currentTab?.databaseQueryResultMap?.values() || [])
      .map((result) => result?.advices || [])
      .flat() ?? []
  ).map((advice) => ({
    severity: "ERROR",
    message: advice.content,
    startLineNumber: advice.line,
    endLineNumber: advice.line,
    startColumn: advice.column,
    endColumn: advice.column,
    source: advice.detail,
  }));
});
const { instance } = useInstanceV1ByUID(
  computed(() => tabStore.currentTab.connection.instanceId)
);
const { database } = useDatabaseV1ByUID(
  computed(() => tabStore.currentTab.connection.databaseId)
);
const instanceEngine = computed(() => {
  return formatEngineV1(instance.value);
});
const language = useInstanceV1EditorLanguage(instance);
const dialect = computed((): SQLDialect => {
  const engine = instance.value.engine;
  return dialectOfEngineV1(engine);
});
const readonly = computed(() => sheetAndTabStore.isReadOnly);
const currentTabId = computed(() => tabStore.currentTabId);
const isSwitchingTab = ref(false);

const filename = computed(() => {
  const name = tabStore.currentTab.id || uuidv1();
  const ext = extensionNameOfLanguage(language.value);
  return `${name}.${ext}`;
});

watch(currentTabId, () => {
  isSwitchingTab.value = true;
  nextTick(() => {
    isSwitchingTab.value = false;
  });
});

const handleChange = (value: string) => {
  // When we are switching between tabs, the MonacoEditor emits a 'change'
  // event, but we shouldn't update the current tab;
  if (isSwitchingTab.value) {
    return;
  }
  if (value === tabStore.currentTab.statement) {
    return;
  }
  // Clear old advices when the statement is changed.
  tabStore.currentTab.databaseQueryResultMap?.forEach((result) => {
    result.advices = [];
  });
  tabStore.updateCurrentTab({
    statement: value,
    isSaved: false,
  });
};

const handleChangeSelection = (value: string) => {
  tabStore.updateCurrentTab({
    selectedStatement: value,
  });
};

const handleSaveSheet = () => {
  editorEvents.emit("save-sheet", { title: tabStore.currentTab.name });
};

const runQueryAction = (explain = false) => {
  const tab = tabStore.currentTab;
  const query = tab.selectedStatement || tab.statement || "";
  emit("execute", query, { databaseType: instanceEngine.value }, { explain });
  uiStateStore.saveIntroStateByKey({
    key: "data.query",
    newState: true,
  });
};

const handleEditorReady = (
  monaco: MonacoModule,
  editor: IStandaloneCodeEditor
) => {
  editor.addAction({
    id: "RunQuery",
    label: "Run Query",
    keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter],
    contextMenuGroupId: "operation",
    contextMenuOrder: 0,
    run: () => runQueryAction(false),
  });
  editor.addAction({
    id: "ExplainQuery",
    label: "Explain Query",
    keybindings: [monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyE],
    contextMenuGroupId: "operation",
    contextMenuOrder: 0,
    run: () => runQueryAction(true),
  });
  editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, () => {
    handleSaveSheet();
  });

  watch(
    () => sqlEditorStore.shouldFormatContent,
    async (shouldFormat) => {
      if (shouldFormat) {
        await formatEditorContent(editor, dialect.value);
        sqlEditorStore.setShouldFormatContent(false);
      }
    },
    {
      immediate: true,
    }
  );
};
</script>
