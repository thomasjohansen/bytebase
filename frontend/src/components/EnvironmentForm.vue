<template>
  <component :is="create ? DrawerContent : 'div'" v-bind="bindings">
    <div
      class="divide-y divide-block-border"
      :class="create ? 'w-[36rem]' : 'w-full px-4 pb-4'"
    >
      <div class="flex flex-col gap-y-6">
        <div class="flex flex-col gap-y-2">
          <label for="name" class="textlabel">
            {{ $t("common.environment-name") }}
            <span class="text-red-600">*</span>
          </label>
          <NInput
            v-model:value="state.environment.title"
            :disabled="!allowEdit"
            size="large"
          />

          <ResourceIdField
            ref="resourceIdField"
            resource-type="environment"
            :readonly="!create"
            :value="extractEnvironmentResourceName(state.environment.name)"
            :resource-title="state.environment.title"
            :validate="validateResourceId"
          />
        </div>

        <div class="flex flex-col gap-y-2">
          <label class="textlabel flex items-center">
            {{ $t("policy.environment-tier.name") }}
            <FeatureBadge feature="bb.feature.environment-tier-policy" />
          </label>
          <p class="text-sm text-gray-600">
            <i18n-t tag="span" keypath="policy.environment-tier.description">
              <template #newline><br /></template>
            </i18n-t>
            <a
              class="inline-flex items-center text-blue-600 ml-1 hover:underline"
              href="https://www.bytebase.com/docs/administration/environment-policy/tier"
              target="_blank"
              >{{ $t("common.learn-more")
              }}<heroicons-outline:external-link class="w-4 h-4"
            /></a>
          </p>
          <NCheckbox
            :checked="state.environmentTier === EnvironmentTier.PROTECTED"
            :disabled="!allowEdit"
            style="--n-label-padding: 0 0 0 1rem"
            @update:checked="(on: boolean) => {
                state.environmentTier = on ? EnvironmentTier.PROTECTED : EnvironmentTier.UNPROTECTED
              }"
          >
            {{ $t("policy.environment-tier.mark-env-as-production") }}
          </NCheckbox>
        </div>

        <div class="flex flex-col gap-y-2">
          <label class="textlabel">
            {{ $t("policy.rollout.name") }}
          </label>
          <span
            v-show="!create && valueChanged('rolloutPolicy')"
            class="textlabeltip !ml-0"
            >{{ $t("policy.rollout.tip") }}</span
          >
          <div class="textinfolabel">
            {{ $t("policy.rollout.info") }}
            <a
              class="inline-flex items-center text-blue-600 ml-1 hover:underline"
              href="https://www.bytebase.com/docs/administration/environment-policy/rollout-policy"
              target="_blank"
              >{{ $t("common.learn-more")
              }}<heroicons-outline:external-link class="w-4 h-4"
            /></a>
          </div>
          <RolloutPolicyConfig
            v-model:policy="state.rolloutPolicy"
            :disabled="!allowEdit"
          />
        </div>

        <div class="flex flex-col gap-y-2">
          <label class="textlabel"> {{ $t("policy.backup.name") }} </label>
          <span
            v-show="valueChanged('backupPolicy')"
            class="textlabeltip !ml-0"
            >{{ $t("policy.backup.tip") }}</span
          >
          <NRadioGroup
            v-model:value="state.backupPolicy.backupPlanPolicy!.schedule"
            :disabled="!allowEdit"
            class="!flex flex-col gap-y-2"
          >
            <NRadio
              :value="BackupPlanSchedule.UNSET"
              style="--n-label-padding: 0 0 0 1rem"
            >
              <div class="flex flex-col">
                <div class="textlabel">
                  {{ $t("policy.backup.not-enforced") }}
                </div>
                <div class="textinfolabel">
                  {{ $t("policy.backup.not-enforced-info") }}
                </div>
              </div>
            </NRadio>
            <NRadio
              :value="BackupPlanSchedule.DAILY"
              style="--n-label-padding: 0 0 0 1rem"
            >
              <div class="flex flex-col">
                <div class="textlabel flex">
                  {{ $t("policy.backup.daily") }}
                  <FeatureBadge feature="bb.feature.backup-policy" />
                </div>
                <div class="textinfolabel">
                  {{ $t("policy.backup.daily-info") }}
                </div>
              </div>
            </NRadio>
            <NRadio
              :value="BackupPlanSchedule.WEEKLY"
              style="--n-label-padding: 0 0 0 1rem"
            >
              <div class="flex flex-col">
                <div class="textlabel flex">
                  {{ $t("policy.backup.weekly") }}
                  <FeatureBadge feature="bb.feature.backup-policy" />
                </div>
                <div class="mt-1 textinfolabel">
                  {{ $t("policy.backup.weekly-info") }}
                </div>
              </div>
            </NRadio>
          </NRadioGroup>
        </div>

        <div v-if="!create" class="flex flex-col gap-y-2">
          <label class="textlabel">
            {{ $t("sql-review.title") }}
          </label>
          <div>
            <div
              v-if="sqlReviewPolicy"
              class="inline-flex items-center gap-x-2"
            >
              <Switch
                v-if="allowEditSQLReviewPolicy"
                :value="sqlReviewPolicy.enforce"
                :text="true"
                @update:value="toggleSQLReviewPolicy"
              />
              <NButton
                quaternary
                type="primary"
                @click="onSQLReviewPolicyClick"
              >
                {{ sqlReviewPolicy.name }}
              </NButton>
            </div>
            <NButton
              v-else-if="hasPermission('bb.policies.update')"
              @click.prevent="onSQLReviewPolicyClick"
            >
              {{ $t("sql-review.configure-policy") }}
            </NButton>
            <span v-else class="textinfolabel">
              {{ $t("sql-review.no-policy-set") }}
            </span>
          </div>
        </div>
      </div>

      <div v-if="!create" class="mt-6 flex justify-between items-center pt-5">
        <template
          v-if="(state.environment as Environment).state === State.ACTIVE"
        >
          <BBButtonConfirm
            v-if="allowArchive"
            :style="'ARCHIVE'"
            :button-text="$t('environment.archive')"
            :ok-text="$t('common.archive')"
            :confirm-title="
            $t('environment.archive') + ` '${(state.environment as Environment).title}'?`
          "
            :confirm-description="$t('environment.archive-info')"
            :require-confirm="true"
            @confirm="archiveEnvironment"
          />
        </template>
        <template
          v-else-if="(state.environment as Environment).state === State.DELETED"
        >
          <BBButtonConfirm
            v-if="allowRestore"
            :style="'RESTORE'"
            :button-text="$t('environment.restore')"
            :ok-text="$t('common.restore')"
            :confirm-title="
            $t('environment.restore') + ` '${(state.environment as Environment).title}'?`
          "
            :confirm-description="''"
            :require-confirm="true"
            @confirm="restoreEnvironment"
          />
        </template>
        <div v-else></div>

        <div v-if="allowEdit" class="flex items-center justify-end gap-x-3">
          <NButton
            :disabled="!valueChanged()"
            @click.prevent="revertEnvironment"
          >
            {{ $t("common.revert") }}
          </NButton>
          <NButton
            type="primary"
            :disabled="!valueChanged()"
            @click.prevent="updateEnvironment"
          >
            {{ $t("common.update") }}
          </NButton>
        </div>
      </div>
    </div>

    <template v-if="create" #footer>
      <div class="flex justify-end items-center gap-x-3">
        <NButton @click.prevent="$emit('cancel')">
          {{ $t("common.cancel") }}
        </NButton>
        <NButton
          type="primary"
          :disabled="!allowCreate"
          @click.prevent="createEnvironment"
        >
          {{ $t("common.create") }}
        </NButton>
      </div>
      <!-- Update button group -->
    </template>
  </component>
</template>

<script lang="ts" setup>
import { useEventListener } from "@vueuse/core";
import { cloneDeep, isEqual, isEmpty } from "lodash-es";
import { NButton, NCheckbox, NInput, NRadioGroup } from "naive-ui";
import { Status } from "nice-grpc-common";
import { computed, reactive, PropType, watch, watchEffect, ref } from "vue";
import { useI18n } from "vue-i18n";
import { useRouter, onBeforeRouteLeave } from "vue-router";
import { DrawerContent, Switch } from "@/components/v2";
import ResourceIdField from "@/components/v2/Form/ResourceIdField.vue";
import {
  SETTING_ROUTE_WORKSPACE_SQL_REVIEW_DETAIL,
  SETTING_ROUTE_WORKSPACE_SQL_REVIEW_CREATE,
} from "@/router/dashboard/workspaceSetting";
import {
  pushNotification,
  useCurrentUserV1,
  useEnvironmentV1List,
  useSQLReviewStore,
} from "@/store";
import { environmentNamePrefix } from "@/store/modules/v1/common";
import { useEnvironmentV1Store } from "@/store/modules/v1/environment";
import type {
  ResourceId,
  SQLReviewPolicy,
  ValidatedMessage,
  WorkspacePermission,
} from "@/types";
import { State } from "@/types/proto/v1/common";
import {
  Environment,
  EnvironmentTier,
} from "@/types/proto/v1/environment_service";
import {
  Policy,
  PolicyType,
  BackupPlanSchedule,
} from "@/types/proto/v1/org_policy_service";
import {
  extractEnvironmentResourceName,
  hasWorkspacePermissionV2,
  sqlReviewPolicySlug,
} from "@/utils";
import { getErrorCode } from "@/utils/grpcweb";
import RolloutPolicyConfig from "./EnvironmentForm/RolloutPolicyConfig.vue";

interface LocalState {
  environment: Environment;
  rolloutPolicy: Policy;
  backupPolicy: Policy;
  environmentTier: EnvironmentTier;
}

const props = defineProps({
  create: {
    type: Boolean,
    default: false,
  },
  environment: {
    required: true,
    type: Object as PropType<Environment>,
  },
  rolloutPolicy: {
    required: true,
    type: Object as PropType<Policy>,
  },
  backupPolicy: {
    required: true,
    type: Object as PropType<Policy>,
  },
  environmentTier: {
    required: true,
    type: Number as PropType<EnvironmentTier>,
  },
});

const emit = defineEmits([
  "create",
  "update",
  "cancel",
  "archive",
  "restore",
  "update-policy",
]);

const { t } = useI18n();
const router = useRouter();
const currentUserV1 = useCurrentUserV1();
const environmentV1Store = useEnvironmentV1Store();
const environmentList = useEnvironmentV1List();
const sqlReviewStore = useSQLReviewStore();
const state = reactive<LocalState>({
  environment: cloneDeep(props.environment),
  rolloutPolicy: cloneDeep(props.rolloutPolicy),
  backupPolicy: cloneDeep(props.backupPolicy),
  environmentTier: props.environmentTier,
});
const resourceIdField = ref<InstanceType<typeof ResourceIdField>>();

const bindings = computed(() => {
  if (props.create) {
    return {
      title: t("environment.create"),
    };
  }
  return {};
});

const environmentId = computed(() => {
  if (props.create) {
    return;
  }
  return extractEnvironmentResourceName(
    (props.environment as Environment).name
  );
});

const prepareSQLReviewPolicy = () => {
  if (!environmentId.value) {
    return;
  }
  return sqlReviewStore.getOrFetchReviewPolicyByEnvironmentId(
    environmentId.value
  );
};
watchEffect(prepareSQLReviewPolicy);

const sqlReviewPolicy = computed((): SQLReviewPolicy | undefined => {
  if (!environmentId.value) {
    return;
  }
  return sqlReviewStore.getReviewPolicyByEnvironmentId(environmentId.value);
});

const onSQLReviewPolicyClick = () => {
  if (sqlReviewPolicy.value) {
    router.push({
      name: SETTING_ROUTE_WORKSPACE_SQL_REVIEW_DETAIL,
      params: {
        sqlReviewPolicySlug: sqlReviewPolicySlug(sqlReviewPolicy.value),
      },
    });
  } else {
    router.push({
      name: SETTING_ROUTE_WORKSPACE_SQL_REVIEW_CREATE,
      query: {
        environmentId: environmentId.value,
      },
    });
  }
};

watch(
  () => props.environment,
  (cur) => {
    state.environment = cloneDeep(cur);
  }
);

watch(
  () => props.rolloutPolicy,
  (cur: Policy) => {
    state.rolloutPolicy = cloneDeep(cur);
  }
);

watch(
  () => props.backupPolicy,
  (cur: Policy) => {
    state.backupPolicy = cloneDeep(cur);
  }
);

watch(
  () => props.environmentTier,
  (cur: EnvironmentTier) => {
    state.environmentTier = cur;
  }
);

const hasPermission = (permission: WorkspacePermission) => {
  return hasWorkspacePermissionV2(currentUserV1.value, permission);
};

const validateResourceId = async (
  resourceId: ResourceId
): Promise<ValidatedMessage[]> => {
  if (!resourceId) {
    return [];
  }

  try {
    const env = await environmentV1Store.getOrFetchEnvironmentByName(
      environmentNamePrefix + resourceId,
      true /* silent */
    );
    if (env) {
      return [
        {
          type: "error",
          message: t("resource-id.validation.duplicated", {
            resource: t("resource.environment"),
          }),
        },
      ];
    }
  } catch (error) {
    if (getErrorCode(error) !== Status.NOT_FOUND) {
      throw error;
    }
  }
  return [];
};

const allowArchive = computed(() => {
  return (
    hasPermission("bb.environments.delete") && environmentList.value.length > 1
  );
});

const allowRestore = computed(() => {
  return hasPermission("bb.environments.undelete");
});

const allowEdit = computed(() => {
  return (
    props.create ||
    ((state.environment as Environment).state === State.ACTIVE &&
      hasPermission("bb.environments.update"))
  );
});

const allowCreate = computed(() => {
  return (
    !isEmpty(state.environment?.title) &&
    resourceIdField.value?.resourceId &&
    resourceIdField.value?.isValidated
  );
});

const allowEditSQLReviewPolicy = computed(() => {
  return hasWorkspacePermissionV2(currentUserV1.value, "bb.policies.update");
});

const valueChanged = (
  field?:
    | "environment"
    | "approvalPolicy"
    | "rolloutPolicy"
    | "backupPolicy"
    | "environmentTier"
): boolean => {
  switch (field) {
    case "environment":
      return !isEqual(props.environment, state.environment);
    case "rolloutPolicy":
      return !isEqual(props.rolloutPolicy, state.rolloutPolicy);
    case "backupPolicy":
      return !isEqual(props.backupPolicy, state.backupPolicy);
    case "environmentTier":
      return !isEqual(props.environmentTier, state.environmentTier);

    default:
      return (
        !isEqual(props.environment, state.environment) ||
        !isEqual(props.rolloutPolicy, state.rolloutPolicy) ||
        !isEqual(props.backupPolicy, state.backupPolicy) ||
        !isEqual(props.environmentTier, state.environmentTier)
      );
  }
};

useEventListener("beforeunload", (e) => {
  if (props.create || !valueChanged()) {
    return;
  }
  e.returnValue = t("common.leave-without-saving");
  return e.returnValue;
});

onBeforeRouteLeave((to, from, next) => {
  if (!props.create && valueChanged()) {
    if (!window.confirm(t("common.leave-without-saving"))) {
      return;
    }
  }
  next();
});

const revertEnvironment = () => {
  state.environment = cloneDeep(props.environment!);
  state.rolloutPolicy = cloneDeep(props.rolloutPolicy!);
  state.backupPolicy = cloneDeep(props.backupPolicy!);
  state.environmentTier = cloneDeep(props.environmentTier!);
};

const createEnvironment = () => {
  emit(
    "create",
    {
      name: resourceIdField.value?.resourceId,
      title: state.environment.title,
    },
    state.rolloutPolicy,
    state.backupPolicy,
    state.environmentTier
  );
};

const updateEnvironment = () => {
  const env = props.environment;
  if (
    state.environment.title !== env.title ||
    state.environmentTier !== env.tier
  ) {
    const patchedEnvironment = {
      title: state.environment.title,
      tier: state.environmentTier,
    };
    emit("update", patchedEnvironment);
  }

  if (!isEqual(props.rolloutPolicy, state.rolloutPolicy)) {
    emit(
      "update-policy",
      state.environment,
      PolicyType.ROLLOUT_POLICY,
      state.rolloutPolicy
    );
  }

  if (!isEqual(props.backupPolicy, state.backupPolicy)) {
    emit(
      "update-policy",
      state.environment,
      PolicyType.BACKUP_PLAN,
      state.backupPolicy
    );
  }
};

const archiveEnvironment = () => {
  emit("archive", state.environment);
};

const restoreEnvironment = () => {
  emit("restore", state.environment);
};

const toggleSQLReviewPolicy = async (on: boolean) => {
  const policy = sqlReviewPolicy.value;
  if (!policy) return;
  const originalOn = policy.enforce;
  if (on === originalOn) return;
  await useSQLReviewStore().updateReviewPolicy({
    id: policy.id,
    enforce: on,
  });
  pushNotification({
    module: "bytebase",
    style: "SUCCESS",
    title: t("sql-review.policy-updated"),
  });
};
</script>
