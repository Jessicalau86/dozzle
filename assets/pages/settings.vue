<template>
  <PageWithLinks>
    <section>
      <div class="has-underline">
        <h2>{{ $t("settings.about") }}</h2>
      </div>

      <div>
        <span v-html="$t('settings.using-version', { version: config.version })"></span>
        <div
          v-if="hasUpdate"
          v-html="$t('settings.update-available', { nextVersion: latest?.name, href: latest?.htmlUrl })"
        ></div>
      </div>
    </section>

    <section class="flex flex-col @container">
      <div class="has-underline">
        <h2>{{ $t("settings.display") }}</h2>
      </div>

      <section class="grid-cols-2 gap-4 @3xl:grid">
        <div class="flex flex-col gap-2 text-balance @3xl:pr-8">
          <Toggle v-model="compact"> {{ $t("settings.compact") }} </Toggle>

          <Toggle v-model="smallerScrollbars"> {{ $t("settings.small-scrollbars") }} </Toggle>

          <Toggle v-model="showTimestamp">{{ $t("settings.show-timesamps") }}</Toggle>

          <Toggle v-model="showStd">{{ $t("settings.show-std") }}</Toggle>

          <Toggle v-model="softWrap">{{ $t("settings.soft-wrap") }}</Toggle>

          <LabeledInput>
            <template #label>
              {{ $t("settings.locale") }}
            </template>
            <template #input>
              <DropdownMenu
                v-model="locale"
                :options="[
                  { label: 'Auto', value: '' },
                  ...availableLocales.map((l) => ({ label: l.toLocaleUpperCase(), value: l })),
                ]"
              />
            </template>
          </LabeledInput>

          <LabeledInput>
            <template #label>
              {{ $t("settings.datetime-format") }}
            </template>
            <template #input>
              <div class="flex gap-2">
                <DropdownMenu
                  v-model="dateLocale"
                  :options="[
                    { label: 'Auto', value: 'auto' },
                    { label: 'MM/DD/YYYY', value: 'en-US' },
                    { label: 'DD/MM/YYYY', value: 'en-GB' },
                    { label: 'DD.MM.YYYY', value: 'de-DE' },
                    { label: 'YYYY-MM-DD', value: 'en-CA' },
                  ]"
                />
                <DropdownMenu
                  v-model="hourStyle"
                  :options="[
                    { label: 'Auto', value: 'auto' },
                    { label: '12', value: '12' },
                    { label: '24', value: '24' },
                  ]"
                />
              </div>
            </template>
          </LabeledInput>

          <LabeledInput>
            <template #label>
              {{ $t("settings.font-size") }}
            </template>
            <template #input>
              <DropdownMenu
                v-model="size"
                :options="[
                  { label: 'Small', value: 'small' },
                  { label: 'Medium', value: 'medium' },
                  { label: 'Large', value: 'large' },
                ]"
              />
            </template>
          </LabeledInput>

          <LabeledInput>
            <template #label>
              {{ $t("settings.color-scheme") }}
            </template>
            <template #input>
              <DropdownMenu
                v-model="lightTheme"
                :options="[
                  { label: 'Auto', value: 'auto' },
                  { label: 'Dark', value: 'dark' },
                  { label: 'Light', value: 'light' },
                ]"
              />
            </template>
          </LabeledInput>
        </div>
        <LogList
          :messages="fakeMessages"
          :last-selected-item="undefined"
          :show-container-name="false"
          class="hidden overflow-hidden rounded-lg border border-base-content/50 shadow @3xl:block"
        />
      </section>
    </section>

    <section class="flex flex-col gap-2">
      <div class="has-underline">
        <h2>{{ $t("settings.options") }}</h2>
      </div>
      <div>
        <toggle v-model="search">
          {{ $t("settings.search") }} <key-shortcut char="f" class="align-top"></key-shortcut>
        </toggle>
      </div>

      <div>
        <toggle v-model="showAllContainers">{{ $t("settings.show-stopped-containers") }}</toggle>
      </div>

      <div>
        <toggle v-model="automaticRedirect">{{ $t("settings.automatic-redirect") }}</toggle>
      </div>
    </section>
  </PageWithLinks>
</template>

<script lang="ts" setup>
import { ComplexLogEntry, SimpleLogEntry } from "@/models/LogEntry";

import {
  automaticRedirect,
  compact,
  hourStyle,
  dateLocale,
  lightTheme,
  search,
  showAllContainers,
  showStd,
  showTimestamp,
  size,
  smallerScrollbars,
  softWrap,
  locale,
} from "@/stores/settings";

import { availableLocales, i18n } from "@/modules/i18n";

const { t } = useI18n();

setTitle(t("title.settings"));
const { latest, hasUpdate } = useReleases();

const now = new Date();
const hoursAgo = (hours: number) => {
  const date = new Date(now);
  date.setHours(date.getHours() - hours);
  return date;
};

const fakeMessages = computedWithControl(
  () => i18n.global.locale.value,
  () => [
    new SimpleLogEntry(t("settings.log.preview"), "123", 1, hoursAgo(16), "info", undefined, "stdout"),
    new SimpleLogEntry(t("settings.log.warning"), "123", 2, hoursAgo(12), "warn", undefined, "stdout"),
    new SimpleLogEntry(
      t("settings.log.multi-line-error.start-line"),
      "123",
      3,
      hoursAgo(7),
      "error",
      "start",
      "stderr",
    ),
    new SimpleLogEntry(
      t("settings.log.multi-line-error.middle-line"),
      "123",
      4,
      hoursAgo(2),
      "error",
      "middle",
      "stderr",
    ),
    new SimpleLogEntry(t("settings.log.multi-line-error.end-line"), "123", 5, new Date(), "error", "end", "stderr"),
    new ComplexLogEntry(
      {
        message: t("settings.log.complex"),
        context: {
          key: "value",
          key2: "value2",
        },
      },
      "123",
      6,
      new Date(),
      "info",
      "stdout",
    ),
    new SimpleLogEntry(t("settings.log.simple"), "123", 7, new Date(), "debug", undefined, "stderr"),
  ],
);
</script>
<style lang="postcss" scoped>
.has-underline {
  @apply mb-4 border-b border-base-content/50 py-4;
}

:deep(a:not(.menu a)) {
  @apply text-primary underline-offset-4 hover:underline;
}
</style>
