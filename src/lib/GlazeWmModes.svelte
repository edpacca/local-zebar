<script lang="ts">
  import type { GlazeWmOutput } from "zebar";

  export let glazewm: GlazeWmOutput | null;

  $: iconClassTiling = `nf nf-md-swap_${glazewm?.tilingDirection}`;

  const iconClassPauased = `nf nf-md-pause`;

  const setBindingMode = (name: string) => {
    const command = `wm-disable-binding-mode --name ${name}`;
    runCommand(command);
  };

  const toggleTilingDirection = () => {
    runCommand("toggle-tiling-direction");
  };

  const togglePaused = () => {
    runCommand("wm-toggle-pause");
  };

  const runCommand = (command: string) => {
    glazewm?.runCommand(command);
  };
</script>

{#if glazewm}
  <div class="flex-gap">
    <!-- glazewm type not properly defined but does gain isPaused when paused -->
    {#if (glazewm as any).isPaused}
      <!-- svelte-ignore a11y_consider_explicit_label -->
      <button class={iconClassPauased} onclick={togglePaused}></button>
    {/if}
    {#each glazewm?.bindingModes as bindingMode}
      <button onclick={() => setBindingMode(bindingMode.name)}>
        {bindingMode.displayName ?? bindingMode.name}
      </button>
    {/each}
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button class={iconClassTiling} onclick={toggleTilingDirection}></button>
  </div>
{/if}
