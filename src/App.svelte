<script>
  import Network from "./lib/Network.svelte";
  import GlazeWmWorkspaces from "./lib/GlazeWmWorkspaces.svelte";
  import GlazeWmModes from "./lib/GlazeWmModes.svelte";
  import { createProviderGroup } from "zebar";
    import Battery from "./lib/Battery.svelte";
    import Cpu from "./lib/CPU.svelte";
    import Memory from "./lib/Memory.svelte";

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
      {#if output.glazewm}
        <GlazeWmWorkspaces glazewm={output.glazewm} />
      {/if}
    </div>
    <div class="center">

    </div>
    <div class="right">
      {#if output.glazewm}
        <GlazeWmModes glazewm={output.glazewm}/>
      {/if}
      {#if output.network}
        <Network network={output.network} />
        {/if}
      {#if output.memory}
         <Memory memory={output.memory} />
      {/if}
      {#if output.cpu}
        <Cpu cpu={output.cpu} />
      {/if}
      {#if output.battery}
        <Battery battery={output.battery} />
      {/if}
    </div>
  </div>
{/if}

<style>
  .bar {
    display: grid;
    grid-template-columns: 1fr auto 1fr;
    align-items: center;
    height: 100%;
    padding: 4px;
  }

  .left,
  .center,
  .right {
    display: flex;
    align-items: center;
    gap: 16px;
  }

  .center {
    justify-self: center;
    border: 0px 1px solid var(--lavender);
  }

  .right {
    justify-self: end;
    padding-right: 8px;
  }
</style>
