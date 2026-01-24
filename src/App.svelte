<script>
  import Network from "./lib/Network.svelte";
  import GlazeWmWorkspaces from "./lib/GlazeWmWorkspaces.svelte";
  import GlazeWmModes from "./lib/GlazeWmModes.svelte";
  import { createProviderGroup } from "zebar";
  import Battery from "./lib/Battery.svelte";
  import Cpu from "./lib/CPU.svelte";
  import Memory from "./lib/Memory.svelte";
  import SongInfo from "./lib/SongInfo.svelte";
  import Audio from "./lib/Audio.svelte";
  import SysTray from "./lib/SysTray.svelte";
  import Vlc from "./lib/VLC.svelte";
  // import Keyboard from "./lib/Keyboard.svelte";

  const rawSystemIds = import.meta.env.VITE_ALLOWED_SYSTRAY_IDS || "";
  const allowedSystrayIds = rawSystemIds
    .split(",")
    .map((/** @type {string} */ s) => s.trim());

  const vlcId = import.meta.env.VITE_VLC_ID;

  const providers = createProviderGroup({
    network: { type: "network" },
    glazewm: { type: "glazewm" },
    cpu: { type: "cpu" },
    date: { type: "date", formatting: "dd/MM/yy   t" },
    battery: { type: "battery" },
    memory: { type: "memory" },
    media: { type: "media" },
    audio: { type: "audio" },
    systray: { type: "systray" },
  });
  /**
   * @type {{ network: import("zebar").NetworkOutput | null; glazewm: import("zebar").GlazeWmOutput | null; cpu: import("zebar").CpuOutput | null; date: import("zebar").DateOutput | null; battery: import("zebar").BatteryOutput | null; memory: import("zebar").MemoryOutput | null; media: import("zebar").MediaOutput | null; audio: import("zebar").AudioOutput | null;  systray: import("zebar").SystrayOutput | null; }}
   */
  let output;
  $: providers.onOutput(() => (output = providers.outputMap));

//   let vlcSongInfo = () => {
//     if (output.systray) {
//         return output.systray.icons.find(i => i.id = vlcId)?.tooltip;
//     }
//     return undefined;
//   };
</script>

{#if output}
  <div class="bar">
    <div class="left">
      {#if output.glazewm}
        <GlazeWmWorkspaces glazewm={output.glazewm} />
        <GlazeWmModes glazewm={output.glazewm} />
      {/if}
    </div>
    <div class="right">
      <div class="lr-border media">
        {#each output.media?.allSessions ?? [] as session}
          <SongInfo artist={session.artist} title={session.title} />
        {/each}
        <!-- {#if vlcSongInfo()}
          <SongInfo artist={""} title={vlcSongInfo()} />
        {/if} -->
        <!-- <Vlc/> -->
        {#if output.audio}
          <Audio audio={output.audio} />
        {/if}
      </div>
      <div class="r-border diagnostics">
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
        {#if output.systray}
          <SysTray systray={output.systray} allowedIds={allowedSystrayIds} />
        {/if}
      <div>
        {output.date?.formatted}
      </div>
    </div>
  </div>
{/if}

<style>
  .bar {
    display: grid;
    grid-template-columns: auto 1fr;
    align-items: center;
    height: 100%;
    padding: 4px;
  }

  .left,
  .right,
  .media {
    display: flex;
    align-items: center;
    flex-direction: row;
    gap: 16px;
  }
  .left {
    width: max-content;
    flex-shrink: 0;
  }

  .right {
    display: grid;
    grid-template-columns: 1fr auto auto auto auto;
    justify-self: end;
    padding-right: 8px;
  }

  .media {
    min-width: 0;
  }

  .diagnostics {
    display: flex;
    flex-direction: row;
    gap: 12px;
  }
</style>
