<template>
  <Tag size="small" :std="std" class="!items-start">
    {{ std }}
  </Tag>
</template>

<script lang="ts" setup>
import { Std } from "@/models/LogEntry";

defineProps<{
  std: Std;
}>();
</script>

<style lang="postcss" scoped>
[std="stdout"] {
  @apply text-blue;
}

[std="stderr"] {
  @apply text-red;
}
</style>
