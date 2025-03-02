<template>
  <div
    v-show="show"
    ref="containerElRef"
    class="w-full h-full overflow-x-auto"
    :data-height="containerHeight"
    :data-table-header-height="tableHeaderHeight"
    :data-table-body-height="tableBodyHeight"
  >
    <NDataTable
      v-bind="$attrs"
      ref="dataTableRef"
      size="small"
      :row-key="getColumnKey"
      :columns="columns"
      :data="layoutReady ? shownColumnList : []"
      :row-class-name="classesForRow"
      :max-height="tableBodyHeight"
      :virtual-scroll="true"
      :striped="true"
      :bordered="true"
      :bottom-bordered="true"
      class="schema-editor-table-column-editor"
      :class="[disableDiffColoring && 'disable-diff-coloring']"
    />
  </div>

  <ColumnDefaultValueExpressionModal
    v-if="editColumnDefaultValueExpressionContext"
    :expression="editColumnDefaultValueExpressionContext.defaultExpression"
    @close="editColumnDefaultValueExpressionContext = undefined"
    @update:expression="handleSelectColumnDefaultValueExpression"
  />

  <SelectClassificationDrawer
    v-if="classificationConfig && state.pendingUpdateColumn"
    :show="state.showClassificationDrawer"
    :classification-config="classificationConfig"
    @dismiss="state.showClassificationDrawer = false"
    @apply="onClassificationSelect"
  />

  <SemanticTypesDrawer
    v-if="state.pendingUpdateColumn"
    :show="state.showSemanticTypesDrawer"
    :semantic-type-list="semanticTypeList"
    @dismiss="state.showSemanticTypesDrawer = false"
    @apply="onSemanticTypeApply($event)"
  />

  <LabelEditorDrawer
    v-if="state.pendingUpdateColumn"
    :show="state.showLabelsDrawer"
    :readonly="false"
    :title="
      $t('db.labels-for-resource', {
        resource: `'${state.pendingUpdateColumn.name}'`,
      })
    "
    :labels="[configForColumn(state.pendingUpdateColumn).labels]"
    @dismiss="state.showLabelsDrawer = false"
    @apply="onLabelsApply"
  />
</template>

<script lang="ts" setup>
import { useElementSize } from "@vueuse/core";
import { pick } from "lodash-es";
import {
  DataTableColumn,
  DataTableInst,
  NCheckbox,
  NDataTable,
} from "naive-ui";
import { computed, h, reactive, ref } from "vue";
import { useI18n } from "vue-i18n";
import SelectClassificationDrawer from "@/components/SchemaTemplate/SelectClassificationDrawer.vue";
import SemanticTypesDrawer from "@/components/SensitiveData/components/SemanticTypesDrawer.vue";
import { InlineInput } from "@/components/v2";
import { useSettingV1Store, useSubscriptionV1Store } from "@/store/modules";
import { ComposedDatabase } from "@/types";
import { Engine } from "@/types/proto/v1/common";
import {
  ColumnConfig,
  ColumnMetadata,
  DatabaseMetadata,
  ForeignKeyMetadata,
  SchemaMetadata,
  TableMetadata,
} from "@/types/proto/v1/database_service";
import { DataClassificationSetting_DataClassificationConfig as DataClassificationConfig } from "@/types/proto/v1/setting_service";
import ColumnDefaultValueExpressionModal from "../../Modals/ColumnDefaultValueExpressionModal.vue";
import { useSchemaEditorContext } from "../../context";
import { EditStatus } from "../../types";
import { getDefaultValueByKey, isTextOfColumnType } from "../../utils";
import { markUUID } from "../common";
import {
  ClassificationCell,
  DataTypeCell,
  ForeignKeyCell,
  OperationCell,
  ReorderCell,
  SelectionCell,
  SemanticTypeCell,
} from "./components";
import DefaultValueCell from "./components/DefaultValueCell.vue";
import LabelsCell from "./components/LabelsCell.vue";

interface LocalState {
  pendingUpdateColumn?: ColumnMetadata;
  showClassificationDrawer: boolean;
  showSemanticTypesDrawer: boolean;
  showLabelsDrawer: boolean;
}

const props = withDefaults(
  defineProps<{
    show?: boolean;
    readonly: boolean;
    showForeignKey?: boolean;
    db: ComposedDatabase;
    database: DatabaseMetadata;
    schema: SchemaMetadata;
    table: TableMetadata;
    engine: Engine;
    classificationConfigId?: string;
    disableChangeTable?: boolean;
    allowReorderColumns?: boolean;
    maxBodyHeight?: number;
    filterColumn?: (column: ColumnMetadata) => boolean;
    disableAlterColumn?: (column: ColumnMetadata) => boolean;
    getColumnItemComputedClassList?: (column: ColumnMetadata) => string;
  }>(),
  {
    show: true,
    showForeignKey: true,
    disableChangeTable: false,
    allowReorderColumns: false,
    maxBodyHeight: undefined,
    classificationConfigId: "",
    filterColumn: (_: ColumnMetadata) => true,
    disableAlterColumn: (_: ColumnMetadata) => false,
    getColumnItemComputedClassList: (_: ColumnMetadata) => "",
  }
);

const emit = defineEmits<{
  (event: "drop", column: ColumnMetadata): void;
  (event: "restore", column: ColumnMetadata): void;
  (
    event: "reorder",
    column: ColumnMetadata,
    index: number,
    delta: -1 | 1
  ): void;
  (
    event: "foreign-key-edit",
    column: ColumnMetadata,
    fk: ForeignKeyMetadata | undefined
  ): void;
  (
    event: "foreign-key-click",
    column: ColumnMetadata,
    fk: ForeignKeyMetadata
  ): void;
  (
    event: "primary-key-set",
    column: ColumnMetadata,
    isPrimaryKey: boolean
  ): void;
}>();

const state = reactive<LocalState>({
  showClassificationDrawer: false,
  showSemanticTypesDrawer: false,
  showLabelsDrawer: false,
});

const {
  resourceType,
  disableDiffColoring,
  selectionEnabled,
  markEditStatus,
  getColumnStatus,
  getColumnConfig,
  upsertColumnConfig,
  useConsumePendingScrollToColumn,
  getAllColumnsSelectionState,
  updateAllColumnsSelection,
} = useSchemaEditorContext();
const dataTableRef = ref<DataTableInst>();
const containerElRef = ref<HTMLElement>();
const tableHeaderElRef = computed(
  () =>
    containerElRef.value?.querySelector("thead.n-data-table-thead") as
      | HTMLElement
      | undefined
);
const { height: containerHeight } = useElementSize(containerElRef);
const { height: tableHeaderHeight } = useElementSize(tableHeaderElRef);
const tableBodyHeight = computed(() => {
  const bodyHeight = containerHeight.value - tableHeaderHeight.value - 2;
  const { maxBodyHeight = 0 } = props;
  if (maxBodyHeight > 0) {
    return Math.min(maxBodyHeight, bodyHeight);
  }
  return bodyHeight;
});
// Use this to avoid unnecessary initial rendering
const layoutReady = computed(() => tableHeaderHeight.value > 0);
const { t } = useI18n();
const subscriptionV1Store = useSubscriptionV1Store();
const settingStore = useSettingV1Store();
const editColumnDefaultValueExpressionContext = ref<ColumnMetadata>();

const metadataForColumn = (column: ColumnMetadata) => {
  return {
    database: props.database,
    schema: props.schema,
    table: props.table,
    column,
  };
};
const statusForColumn = (column: ColumnMetadata) => {
  return getColumnStatus(props.db, metadataForColumn(column));
};
const markColumnStatus = (
  column: ColumnMetadata,
  status: EditStatus,
  oldStatus: EditStatus | undefined = undefined
) => {
  if (!oldStatus) {
    oldStatus = statusForColumn(column);
  }
  if (
    (oldStatus === "created" || oldStatus === "dropped") &&
    status === "updated"
  ) {
    markEditStatus(props.db, metadataForColumn(column), oldStatus);
    return;
  }
  markEditStatus(props.db, metadataForColumn(column), status);
};
const configForColumn = (column: ColumnMetadata) => {
  return (
    getColumnConfig(props.db, metadataForColumn(column)) ??
    ColumnConfig.fromPartial({
      name: column.name,
    })
  );
};

const primaryKey = computed(() => {
  return props.table.indexes.find((idx) => idx.primary);
});
const classificationConfig = computed(() => {
  if (!props.classificationConfigId) {
    return;
  }
  return settingStore.getProjectClassification(props.classificationConfigId);
});
const semanticTypeList = computed(() => {
  return (
    settingStore.getSettingByName("bb.workspace.semantic-types")?.value
      ?.semanticTypeSettingValue?.types ?? []
  );
});
const showDatabaseConfigColumn = computed(
  () => resourceType.value === "branch"
);
const showSemanticTypeColumn = computed(() => {
  return (
    subscriptionV1Store.hasFeature("bb.feature.sensitive-data") &&
    showDatabaseConfigColumn.value
  );
});

const columns = computed(() => {
  const columns: (DataTableColumn<ColumnMetadata> & { hide?: boolean })[] = [
    {
      key: "__selected__",
      width: 32,
      hide: !selectionEnabled.value,
      title: () => {
        const state = getAllColumnsSelectionState(
          props.db,
          pick(props, "database", "schema", "table"),
          shownColumnList.value
        );
        return h(NCheckbox, {
          checked: state.checked,
          indeterminate: state.indeterminate,
          onUpdateChecked: (on: boolean) => {
            updateAllColumnsSelection(
              props.db,
              pick(props, "database", "schema", "table"),
              shownColumnList.value,
              on
            );
          },
        });
      },
      render: (column) => {
        return h(SelectionCell, {
          db: props.db,
          metadata: {
            ...pick(props, "database", "schema", "table"),
            column,
          },
        });
      },
    },
    {
      key: "reorder",
      title: "",
      resizable: false,
      width: 44,
      hide: props.readonly || !props.allowReorderColumns,
      className: "!px-0",
      render: (column, index) => {
        return h(ReorderCell, {
          column,
          allowMoveUp: index > 0,
          allowMoveDown: index < shownColumnList.value.length - 1,
          disabled: props.disableChangeTable,
          onReorder: (delta) => emit("reorder", column, index, delta),
        });
      },
    },
    {
      key: "name",
      title: t("schema-editor.column.name"),
      resizable: true,
      minWidth: 140,
      className: "input-cell",
      render: (column) => {
        return h(InlineInput, {
          value: column.name,
          disabled: props.readonly || props.disableAlterColumn(column),
          placeholder: "column name",
          style: {
            "--n-padding-left": "6px",
            "--n-padding-right": "4px",
            "--n-text-color-disabled": "rgb(var(--color-main))",
          },
          "onUpdate:value": (value) => {
            const oldStatus = statusForColumn(column);
            column.name = value;
            markColumnStatus(column, "updated", oldStatus);
          },
        });
      },
    },
    {
      key: "semantic-types",
      title: t("settings.sensitive-data.semantic-types.self"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      hide: !showSemanticTypeColumn.value,
      render: (column) => {
        return h(SemanticTypeCell, {
          db: props.db,
          database: props.database,
          schema: props.schema,
          table: props.table,
          column,
          readonly: props.readonly,
          disabled: props.disableChangeTable,
          semanticTypeList: semanticTypeList.value,
          disableAlterColumn: props.disableAlterColumn,
          onRemove: () => onSemanticTypeRemove(column),
          onEdit: () => openSemanticTypeDrawer(column),
        });
      },
    },
    {
      key: "classification",
      title: t("schema-editor.column.classification"),
      hide: !classificationConfig.value,
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      render: (column) => {
        return h(ClassificationCell, {
          classification: column.classification,
          readonly: props.readonly,
          disabled: props.disableChangeTable,
          classificationConfig:
            classificationConfig.value ??
            DataClassificationConfig.fromPartial({}),
          onEdit: () => openClassificationDrawer(column),
          onRemove: () => {
            column.classification = "";
            markColumnStatus(column, "updated");
          },
        });
      },
    },
    {
      key: "type",
      title: t("schema-editor.column.type"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "input-cell",
      render: (column) => {
        return h(DataTypeCell, {
          column,
          disabled: props.readonly || props.disableAlterColumn(column),
          schemaTemplateColumnTypes: schemaTemplateColumnTypes.value,
          engine: props.engine,
          "onUpdate:value": (value) => {
            column.type = value;
            markColumnStatus(column, "updated");
          },
        });
      },
    },
    {
      key: "default-value",
      title: t("schema-editor.column.default"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "input-cell",
      render: (column) => {
        return h(DefaultValueCell, {
          column,
          disabled: props.readonly || props.disableAlterColumn(column),
          schemaTemplateColumnTypes: schemaTemplateColumnTypes.value,
          engine: props.engine,
          onInput: (value) => handleColumnDefaultInput(column, value),
          onSelect: (option) => handleColumnDefaultSelect(column, option.key),
        });
      },
    },
    {
      key: "on-update",
      title: t("schema-editor.column.on-update"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      hide: props.engine !== Engine.MYSQL && props.engine !== Engine.TIDB,
      className: "input-cell",
      render: (column) => {
        return h(InlineInput, {
          value: column.onUpdate,
          disabled: props.readonly || props.disableAlterColumn(column),
          placeholder: "",
          style: {
            "--n-padding-left": "6px",
            "--n-padding-right": "4px",
            "--n-text-color-disabled": "rgb(var(--color-main))",
          },
          "onUpdate:value": (value) => {
            column.onUpdate = value;
            markColumnStatus(column, "updated");
          },
        });
      },
    },
    {
      key: "comment",
      title: t("schema-editor.column.comment"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "input-cell",
      render: (column) => {
        return h(InlineInput, {
          value: column.userComment,
          disabled: props.readonly || props.disableAlterColumn(column),
          placeholder: "comment",
          style: {
            "--n-padding-left": "6px",
            "--n-padding-right": "4px",
            "--n-text-color-disabled": "rgb(var(--color-main))",
          },
          "onUpdate:value": (value) => {
            column.userComment = value;
            markColumnStatus(column, "updated");
          },
        });
      },
    },
    {
      key: "not-null",
      title: t("schema-editor.column.not-null"),
      resizable: true,
      minWidth: 80,
      maxWidth: 160,
      className: "checkbox-cell",
      render: (column) => {
        return h(NCheckbox, {
          checked: !column.nullable,
          disabled:
            props.readonly ||
            props.disableAlterColumn(column) ||
            isColumnPrimaryKey(column),
          "onUpdate:checked": (checked: boolean) => {
            column.nullable = !checked;
            markColumnStatus(column, "updated");
          },
        });
      },
    },
    {
      key: "primary",
      title: t("schema-editor.column.primary"),
      resizable: true,
      minWidth: 80,
      maxWidth: 160,
      className: "checkbox-cell",
      render: (column) => {
        return h(NCheckbox, {
          checked: isColumnPrimaryKey(column),
          disabled: props.readonly || props.disableAlterColumn(column),
          "onUpdate:checked": (checked: boolean) =>
            emit("primary-key-set", column, checked),
        });
      },
    },
    {
      key: "foreign-key",
      title: t("schema-editor.column.foreign-key"),
      hide: !props.showForeignKey,
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      className: "text-cell",
      render: (column) => {
        return h(ForeignKeyCell, {
          db: props.db,
          database: props.database,
          schema: props.schema,
          table: props.table,
          column: column,
          readonly: props.readonly,
          disabled: props.readonly || props.disableAlterColumn(column),
          onClick: (fk) => emit("foreign-key-click", column, fk),
          onEdit: (fk) => emit("foreign-key-edit", column, fk),
        });
      },
    },
    {
      key: "labels",
      title: t("common.labels"),
      resizable: true,
      minWidth: 140,
      maxWidth: 320,
      hide: !showDatabaseConfigColumn.value,
      render: (column) => {
        return h(LabelsCell, {
          db: props.db,
          database: props.database,
          schema: props.schema,
          table: props.table,
          column,
          readonly: props.readonly,
          disabled: props.disableChangeTable,
          onEdit: () => openLabelsDrawer(column),
        });
      },
    },
    {
      key: "operations",
      title: "",
      resizable: false,
      width: 30,
      hide: props.readonly,
      className: "!px-0",
      render: (column) => {
        return h(OperationCell, {
          column,
          dropped: isDroppedColumn(column),
          disabled: props.disableChangeTable,
          onDrop: () => emit("drop", column),
          onRestore: () => emit("restore", column),
        });
      },
    },
  ];
  return columns.filter((header) => !header.hide);
});

const shownColumnList = computed(() => {
  return props.table.columns.filter(props.filterColumn);
});

const isColumnPrimaryKey = (column: ColumnMetadata): boolean => {
  const pk = primaryKey.value;
  if (!pk) return false;
  return pk.expressions.includes(column.name);
};

const schemaTemplateColumnTypes = computed(() => {
  const setting = settingStore.getSettingByName("bb.workspace.schema-template");
  const columnTypes = setting?.value?.schemaTemplateSettingValue?.columnTypes;
  if (columnTypes && columnTypes.length > 0) {
    const columnType = columnTypes.find(
      (columnType) => columnType.engine === props.engine
    );
    if (columnType && columnType.enabled) {
      return columnType.types;
    }
  }
  return [];
});

const handleColumnDefaultInput = (column: ColumnMetadata, value: string) => {
  column.hasDefault = true;
  column.defaultNull = undefined;
  // If column is text type or has default string, we will treat user's input as string.
  if (
    isTextOfColumnType(props.engine, column.type) ||
    column.defaultString !== undefined
  ) {
    column.defaultString = value;
    column.defaultExpression = undefined;
    markColumnStatus(column, "updated");
    return;
  }
  // Otherwise we will treat user's input as expression.
  column.defaultExpression = value;
  markColumnStatus(column, "updated");
};

const handleColumnDefaultSelect = (column: ColumnMetadata, key: string) => {
  if (key === "expression") {
    editColumnDefaultValueExpressionContext.value = column;
    markColumnStatus(column, "updated");
    return;
  }

  const defaultValue = getDefaultValueByKey(key);
  if (!defaultValue) {
    return;
  }

  column.hasDefault = defaultValue.hasDefault;
  column.defaultNull = defaultValue.defaultNull;
  column.defaultString = defaultValue.defaultString;
  column.defaultExpression = defaultValue.defaultExpression;
  if (column.hasDefault && column.defaultNull) {
    column.nullable = true;
  }
  markColumnStatus(column, "updated");
};

const handleSelectColumnDefaultValueExpression = (expression: string) => {
  const column = editColumnDefaultValueExpressionContext.value;
  if (!column) {
    return;
  }
  column.hasDefault = true;
  column.defaultNull = undefined;
  column.defaultString = undefined;
  column.defaultExpression = expression;

  markColumnStatus(column, "updated");
};

const openClassificationDrawer = (column: ColumnMetadata) => {
  state.pendingUpdateColumn = column;
  state.showClassificationDrawer = true;
};

const openSemanticTypeDrawer = (column: ColumnMetadata) => {
  state.pendingUpdateColumn = column;
  state.showSemanticTypesDrawer = true;
};

const openLabelsDrawer = (column: ColumnMetadata) => {
  state.pendingUpdateColumn = column;
  state.showLabelsDrawer = true;
};

const onClassificationSelect = (classificationId: string) => {
  if (!state.pendingUpdateColumn) {
    return;
  }
  state.pendingUpdateColumn.classification = classificationId;
  markColumnStatus(state.pendingUpdateColumn, "updated");
};

const updateColumnConfig = (
  column: ColumnMetadata,
  update: (config: ColumnConfig) => void
) => {
  upsertColumnConfig(props.db, metadataForColumn(column), update);
  markColumnStatus(column, "updated");
};

const onSemanticTypeApply = async (semanticTypeId: string) => {
  if (!state.pendingUpdateColumn) {
    return;
  }

  updateColumnConfig(state.pendingUpdateColumn, (config) => {
    config.semanticTypeId = semanticTypeId;
  });
};

const onLabelsApply = (labelsList: { [key: string]: string }[]) => {
  if (!state.pendingUpdateColumn) {
    return;
  }
  updateColumnConfig(state.pendingUpdateColumn, (config) => {
    config.labels = labelsList[0];
  });
  markColumnStatus(state.pendingUpdateColumn, "updated");
};

const onSemanticTypeRemove = async (column: ColumnMetadata) => {
  markColumnStatus(column, "updated");
  updateColumnConfig(column, (config) => {
    config.semanticTypeId = "";
  });
};

const classesForRow = (column: ColumnMetadata, index: number) => {
  return props.getColumnItemComputedClassList(column);
};

const isDroppedColumn = (column: ColumnMetadata): boolean => {
  return statusForColumn(column) === "dropped";
};

const getColumnKey = (column: ColumnMetadata) => {
  return markUUID(column);
};

const vlRef = computed(() => {
  return (dataTableRef.value as any)?.$refs?.mainTableInstRef?.bodyInstRef
    ?.virtualListRef;
});
useConsumePendingScrollToColumn(
  computed(() => ({
    db: props.db,
    metadata: {
      database: props.database,
      schema: props.schema,
      table: props.table,
    },
  })),
  vlRef,
  (params, vl) => {
    const key = getColumnKey(params.metadata.column);
    if (!key) return;
    requestAnimationFrame(() => {
      try {
        console.debug("scroll-to-column", vl, params, key);
        vl.scrollTo({ key });
        // TODO: focus name or type input element
      } catch {
        // Do nothing
      }
    });
  }
);
</script>

<style lang="postcss" scoped>
.schema-editor-table-column-editor
  :deep(.n-data-table-th .n-data-table-resize-button::after) {
  @apply bg-control-bg h-2/3;
}
.schema-editor-table-column-editor :deep(.n-data-table-td.input-cell) {
  @apply pl-0.5 pr-1 py-0;
}
.schema-editor-table-column-editor
  :deep(.n-data-table-td.input-cell .n-input__placeholder),
.schema-editor-table-column-editor
  :deep(.n-data-table-td.input-cell .n-base-selection-placeholder) {
  @apply italic;
}
.schema-editor-table-column-editor :deep(.n-data-table-td.checkbox-cell) {
  @apply pr-1 py-0;
}
.schema-editor-table-column-editor :deep(.n-data-table-td.text-cell) {
  @apply pr-1 py-0;
}
.schema-editor-table-column-editor:not(.disable-diff-coloring)
  :deep(.n-data-table-tr.created .n-data-table-td) {
  @apply text-green-700 !bg-green-50;
}
.schema-editor-table-column-editor:not(.disable-diff-coloring)
  :deep(.n-data-table-tr.dropped .n-data-table-td) {
  @apply text-red-700 cursor-not-allowed !bg-red-50 opacity-70;
}
.schema-editor-table-column-editor:not(.disable-diff-coloring)
  :deep(.n-data-table-tr.updated .n-data-table-td) {
  @apply text-yellow-700 !bg-yellow-50;
}
</style>
