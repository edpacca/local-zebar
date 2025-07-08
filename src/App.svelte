<script>
    import GlazeWmWorkspaces from "./lib/GlazeWmWorkspaces.svelte";
    import GlazeWmModes from "./lib/GlazeWmModes.svelte";
    import { createProviderGroup } from "zebar";
    import Cpu from "./lib/CPU.svelte";
    import Memory from "./lib/Memory.svelte";

    const providers = createProviderGroup({
        glazewm: { type: "glazewm" },
        cpu: { type: "cpu" },
        date: { type: "date", formatting: "dd/MM/yy   t" },
        memory: { type: "memory" },
    });
    /**
   * @type {{ glazewm: import("zebar").GlazeWmOutput | null; cpu: import("zebar").CpuOutput | null; date: import("zebar").DateOutput | null; memory: import("zebar").MemoryOutput | null; }}
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
            <div class="flex-gap">
                {#if output.glazewm}
                    <div class="state-indicators">
                        <GlazeWmModes glazewm={output.glazewm} />
                    </div>
                    {/if}
            </div>
            <div class="lr-border diagnostics">
                {#if output.memory}
                    <Memory memory={output.memory} />
                {/if}
                {#if output.cpu}
                    <Cpu cpu={output.cpu} />
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
        grid-template-columns: 1fr auto auto auto;
        justify-self: end;
        padding-right: 8px;
    }

    .media {
        min-width: 0;
    }

    .state-indicators {
        padding-right: 8px;
    }

    .diagnostics {
        display: flex;
        flex-direction: row;
        gap: 12px;
    }
</style>
