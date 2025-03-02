<template>
  <template v-for="[role, bindings] in roleGroup" :key="role">
    <p class="mb-2 text-base pl-2">{{ displayRoleTitle(role) }}</p>
    <BBGrid
      :column-list="columnList"
      :row-clickable="false"
      :data-source="bindings"
      :custom-header="true"
      :ready="ready"
      class="border mb-4"
      row-key="role"
    >
      <template #header>
        <div role="table-row" class="bb-grid-row bb-grid-header-row group">
          <div
            v-for="(column, index) in columnList"
            :key="index"
            role="table-cell"
            class="bb-grid-header-cell capitalize"
            :class="[column.class]"
          >
            {{ column.title }}
          </div>
        </div>
      </template>
      <template #item="{ item: binding }: ProjectRoleRow">
        <div class="bb-grid-cell gap-x-2">
          {{ binding.condition?.title || displayRoleTitle(binding.role) }}
        </div>
        <div
          class="bb-grid-cell flex-wrap gap-x-2 gap-y-1"
          :class="isExpired(binding) ? 'line-through' : ''"
        >
          {{ getExpiredTimeString(binding) || "*" }}
        </div>
        <div class="bb-grid-cell flex-wrap gap-x-2 gap-y-1">
          <div class="flex flex-row justify-start items-start flex-wrap gap-1">
            <div
              v-for="user in getUserList(binding)"
              :key="user.name"
              class="flex flex-row gap-x-1 justify-start items-center border border-gray-200 rounded-md p-1 px-2"
            >
              <UserAvatar size="TINY" :user="user" />
              <span>{{ user.title }}</span>
              <YouTag v-if="currentUserV1.name === user.name" />
            </div>
          </div>
        </div>
        <div class="bb-grid-cell gap-x-2 justify-end">
          <NTooltip v-if="allowEdit" trigger="hover">
            <template #trigger>
              <button
                class="cursor-pointer opacity-60 hover:opacity-100"
                @click="editingBinding = binding"
              >
                <heroicons-outline:pencil class="w-4 h-4" />
              </button>
            </template>
            {{ $t("common.edit") }}
          </NTooltip>
        </div>
      </template>

      <template #placeholder-content>
        <div class="p-2">-</div>
      </template>
    </BBGrid>
  </template>

  <EditProjectRolePanel
    v-if="editingBinding !== null"
    :project="project"
    :binding="editingBinding"
    @close="editingBinding = null"
  />
</template>

<script setup lang="ts">
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import { type BBGridColumn, type BBGridRow, BBGrid } from "@/bbkit";
import YouTag from "@/components/misc/YouTag.vue";
import {
  extractUserEmail,
  useCurrentUserV1,
  useProjectIamPolicy,
  useUserStore,
} from "@/store";
import { ComposedProject, PRESET_ROLES, PresetRoleType } from "@/types";
import { State } from "@/types/proto/v1/common";
import { Binding } from "@/types/proto/v1/iam_policy";
import { displayRoleTitle } from "@/utils";
import { convertFromExpr } from "@/utils/issue/cel";
import EditProjectRolePanel from "./EditProjectRolePanel.vue";
import { getExpiredTimeString, isExpired, getExpiredDateTime } from "./utils";

export type ProjectRoleRow = BBGridRow<Binding>;

const props = defineProps<{
  project: ComposedProject;
  searchText: string;
  ready?: boolean;
  allowEdit: boolean;
}>();

const { t } = useI18n();
const userStore = useUserStore();
const currentUserV1 = useCurrentUserV1();
const editingBinding = ref<Binding | null>(null);

const projectResourceName = computed(() => props.project.name);
const { policy: iamPolicy } = useProjectIamPolicy(projectResourceName);

const columnList = computed(() => {
  const ROLE_NAME: BBGridColumn = {
    title: t("project.members.condition-name"),
    width: "1fr",
  };
  const EXPIRATION: BBGridColumn = {
    title: t("common.expiration"),
    width: "1fr",
  };
  const USERS: BBGridColumn = {
    title: t("common.user"),
    width: "1fr",
  };
  const OPERATIONS: BBGridColumn = {
    title: "",
    width: "4rem",
  };
  return [ROLE_NAME, EXPIRATION, USERS, OPERATIONS];
});

const roleGroup = computed(() => {
  let roleMap = new Map<string, Binding[]>();
  for (const binding of iamPolicy.value.bindings) {
    // Don't show EXPORTER role if it has a non-empty statement condition.
    if (binding.role === PresetRoleType.PROJECT_EXPORTER) {
      const parsedExpr = binding.parsedExpr;
      if (parsedExpr?.expr) {
        const expression = convertFromExpr(parsedExpr.expr);
        if (expression.statement && expression.statement !== "") {
          continue;
        }
      }
    }

    // Filter by search text.
    if (props.searchText !== "") {
      let isMatch = false;
      for (const member of binding.members) {
        const userEmail = extractUserEmail(member);
        const user = userStore.getUserByEmail(userEmail);
        if (!user) {
          continue;
        }
        if (
          user.title.toLowerCase().includes(props.searchText.toLowerCase()) ||
          user.email.toLowerCase().includes(props.searchText.toLowerCase())
        ) {
          isMatch = true;
          break;
        }
      }
      if (!isMatch) {
        continue;
      }
    }

    const role = binding.role;
    if (!roleMap.has(role)) {
      roleMap.set(role, []);
    }
    roleMap.get(role)?.push(binding);
  }
  // Sort by role type.
  roleMap = new Map(
    [...roleMap].sort((a, b) => {
      if (!PRESET_ROLES.includes(a[0])) return -1;
      if (!PRESET_ROLES.includes(b[0])) return 1;
      return PRESET_ROLES.indexOf(a[0]) - PRESET_ROLES.indexOf(b[0]);
    })
  );
  // Sort by expiration time.
  for (const role of roleMap.keys()) {
    roleMap.set(
      role,
      roleMap.get(role)?.sort((a, b) => {
        return (
          (getExpiredDateTime(b)?.getTime() ?? -1) -
          (getExpiredDateTime(a)?.getTime() ?? -1)
        );
      }) || []
    );
  }
  return roleMap;
});

const getUserList = (binding: Binding) => {
  const userList = [];
  for (const member of binding.members) {
    const userEmail = extractUserEmail(member);
    const user = userStore.getUserByEmail(userEmail);
    if (user && user.state === State.ACTIVE) {
      userList.push(user);
    }
  }
  return userList;
};
</script>
