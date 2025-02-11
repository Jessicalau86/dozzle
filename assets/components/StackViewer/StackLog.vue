<template>
  <ScrollableView :scrollable="scrollable" v-if="stack.name">
    <template #header>
      <div class="mx-2 flex items-center gap-2 md:ml-4">
        <div class="flex flex-1 gap-1.5 truncate @container md:gap-2">
          <div class="inline-flex font-mono text-sm">
            <div class="font-semibold">{{ stack.name }}</div>
          </div>
          <Tag class="mobile-hidden hidden font-mono @3xl:block" size="small">
            {{ $t("label.container", stack.containers.length) }}
          </Tag>
          <Tag class="mobile-hidden hidden font-mono @3xl:block" size="small">
            {{ $t("label.serivce", stack.services.length) }}
          </Tag>
        </div>
        <MultiContainerStat class="ml-auto" :containers="stack.containers" />
        <MultiContainerActionToolbar class="mobile-hidden" @clear="viewer?.clear()" />
      </div>
    </template>
    <template #default>
      <ViewerWithSource
        ref="viewer"
        :stream-source="useStackStream"
        :entity="stack"
        :visible-keys="new Map<string[], boolean>()"
      />
    </template>
  </ScrollableView>
</template>

<script lang="ts" setup>
import { Stack } from "@/models/Stack";
import ViewerWithSource from "@/components/LogViewer/ViewerWithSource.vue";
import { ComponentExposed } from "vue-component-type-helpers";
const { name, scrollable = false } = defineProps<{
  scrollable?: boolean;
  name: string;
}>();

const viewer = ref<ComponentExposed<typeof ViewerWithSource>>();
const store = useSwarmStore();
const { stacks } = storeToRefs(store) as unknown as { stacks: Ref<Stack[]> };
const stack = computed(() => stacks.value.find((s) => s.name === name) ?? new Stack("", [], []));
provideLoggingContext(
  toRef(() => stack.value.containers),
  { showContainerName: true, showHostname: false },
);
</script>
