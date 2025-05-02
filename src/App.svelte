<script>
    import Network from "./lib/Network.svelte";
    import GlazeWmWorkspaces from "./lib/GlazeWmWorkspaces.svelte";
    import GlazeWmModes from "./lib/GlazeWmModes.svelte";
    import { createProviderGroup } from "zebar";
    import Battery from "./lib/Battery.svelte";
    import Cpu from "./lib/CPU.svelte";
    import Memory from "./lib/Memory.svelte";
    import Media from "./lib/Media.svelte";
    import Audio from "./lib/Audio.svelte";

    const providers = createProviderGroup({
        network: { type: "network" },
        glazewm: { type: "glazewm" },
        cpu: { type: "cpu" },
        date: { type: "date", formatting: "dd/MM/yy   t" },
        battery: { type: "battery" },
        memory: { type: "memory" },
        media: { type: "media" },
        audio: { type: "audio"}
    });
    /**
     * @type {{ network: import("zebar").NetworkOutput | null; glazewm: import("zebar").GlazeWmOutput | null; cpu: import("zebar").CpuOutput | null; date: import("zebar").DateOutput | null; battery: import("zebar").BatteryOutput | null; memory: import("zebar").MemoryOutput | null; media: import("zebar").MediaOutput | null; audio: import("zebar").AudioOutput | null; }}
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
        <div class="right">
            <div class="lr-border media">
                {#each output.media?.allSessions ?? [] as session }
                    <Media session={session}/>
                {/each}
                {#if output.audio}
                    <Audio audio={output.audio}/>
                {:else}
                    <div>no audio</div>
                {/if}
            </div>
            {#if output.glazewm}
                <GlazeWmModes glazewm={output.glazewm} />
            {/if}
            {#if output.network}
                <Network network={output.network} />
            {/if}
            <div class="diagnostics lr-border">
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
            <div>
                {output.date?.formatted}
            </div>
        </div>
    </div>
{/if}

<style>
    .bar {
        display: grid;
        grid-template-columns: 1fr auto;
        align-items: center;
        height: 100%;
        padding: 4px;
    }

    .left,
    .right, .media {
        display: flex;
        align-items: center;
        flex-direction: row;
        gap: 16px;
    }

    .left {
        flex-grow: 1.5;
    }

    .diagnostics {
        display: flex;
        flex-direction: row;
        gap: 12px;
    }

    .right {
        justify-self: end;
        padding-right: 8px;
    }
</style>
