<script>
  import { onMount } from "svelte";
  import Network from "./lib/Network.svelte";
  import GlazeWmWorkspaces from "./lib/GlazeWmWorkspaces.svelte";
  import GlazeWmModes from "./lib/GlazeWmModes.svelte";
  import { createProviderGroup } from "zebar";

  const providers = createProviderGroup({
    network: { type: "network" },
    glazewm: { type: "glazewm" },
    cpu: { type: "cpu" },
    date: { type: "date", formatting: "EEE d MMM t" },
    battery: { type: "battery" },
    memory: { type: "memory" },
  });
  /**
   * @type {{ network: import("zebar").NetworkOutput | null; glazewm: import("zebar").GlazeWmOutput | null; cpu: import("zebar").CpuOutput | null; date: import("zebar").DateOutput | null; battery: import("zebar").BatteryOutput | null; memory: import("zebar").MemoryOutput | null; }}
   */
  let output;
  $: providers.onOutput(() => (output = providers.outputMap));
</script>

{#if output}
  <div class="bar">
    <div class="left">
      <GlazeWmWorkspaces glazewm={output.glazewm} />
    </div>
    <div class="center">

    </div>
    <div class="right">
      <GlazeWmModes glazewm={output.glazewm} />
      <Network network={output.network} />
    </div>
  </div>
{/if}

<style>
  .bar {
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    align-items: center;
    height: 100%;
    padding: 4px 1.5vw;
  }

  .left,
  .center,
  .right {
    display: flex;
    align-items: center;
    gap: 4px;
  }

  .center {
    justify-self: center;
    border: 1px 0px solid var(--lavender);
  }

  .right {
    justify-self: end;
  }
</style>
