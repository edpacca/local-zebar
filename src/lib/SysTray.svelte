<script lang="ts">
  import type { SystrayOutput } from "zebar";
  export let systray: SystrayOutput;
  export let allowedIds: string[];
</script>

<div class="flex-gap">
  <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
  {#each systray.icons.filter((i) => allowedIds.includes(i.id)) as icon}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <img
      class="systray-icon"
      src={icon.iconUrl}
      title={`${icon.tooltip} \n${icon.id} \n ${icon.iconUrl}`}
      alt={icon.tooltip}
      onclick={(e) => {
        e.preventDefault();
        systray.onLeftClick(icon.id);
      }}
      oncontextmenu={(e) => {
        e.preventDefault();
        systray.onRightClick(icon.id);
      }}
    />
  {/each}
</div>

<style>
  .systray-icon {
    width: 12px;
    height: 12px;
  }
</style>
