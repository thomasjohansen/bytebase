<template>
  <NTooltip :animated="false" :delay="250">
    <template #trigger>
      <heroicons-outline:eye-slash
        class="w-[12px] h-[12px] -mb-[2px]"
        :class="[clickable && 'cursor-pointer']"
        v-bind="$attrs"
        @click="handleClick"
      />
    </template>

    <span class="whitespace-nowrap">
      {{ $t("sensitive-data.self") }}
    </span>
  </NTooltip>
</template>

<script lang="ts" setup>
import { NTooltip } from "naive-ui";
import { computed } from "vue";
import { useRouter } from "vue-router";
import { SETTING_ROUTE_WORKSPACE_SENSITIVE_DATA } from "@/router/dashboard/workspaceSetting";
import { useCurrentUserV1 } from "@/store";
import { hasWorkspacePermissionV2 } from "@/utils";

const user = useCurrentUserV1();
const router = useRouter();

const clickable = computed(() => {
  return hasWorkspacePermissionV2(user.value, "bb.policies.update");
});

const handleClick = () => {
  if (!clickable.value) {
    return;
  }
  const url = router.resolve({ name: SETTING_ROUTE_WORKSPACE_SENSITIVE_DATA });
  window.open(url.href, "_BLANK");
};
</script>
