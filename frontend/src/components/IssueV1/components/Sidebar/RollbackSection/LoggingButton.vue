<template>
  <div
    class="select-none inline-flex items-center gap-x-1 h-[30px] text-control bg-white focus:ring-control focus:outline-none focus-visible:ring-2 focus:ring-offset-2"
  >
    <BBSpin class="!w-4 !h-4" />
    <span class="text-sm leading-5 font-medium">{{
      $t("task.rollback.generating")
    }}</span>
    <NTooltip v-if="!disabled" :animated="false" :delay="150">
      <template #trigger>
        <span
          class="ml-1 py-0.5 px-1.5 text-xs flex items-center rounded border border-control-border hover:bg-control-bg-hover cursor-pointer"
          @click="toggleRollback(false)"
        >
          {{ $t("common.cancel") }}
        </span>
      </template>

      <span>{{ $t("task.rollback.cancel-generating") }}</span>
    </NTooltip>
  </div>
</template>

<script lang="ts" setup>
import { NTooltip } from "naive-ui";
import { computed } from "vue";
import { useRollbackContext } from "./common";

const { allowRollback, toggleRollback } = useRollbackContext();

const disabled = computed(() => !allowRollback.value);
</script>
