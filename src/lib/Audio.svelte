<script lang="ts">
    import type { AudioOutput } from "zebar";

    export let audio: AudioOutput;

    const deviceRegex =  RegExp(/([^\s]+)/m);

    $: device =audio.defaultPlaybackDevice?.name.split(" ")[0];
    $: volume = audio.defaultPlaybackDevice?.volume ?? 0 / 100;

    const mediumVolume = 50;
    const highVolume = 75;
</script>

{#if audio}
    <div class="flex-gap">
        <div class="subtext">{device}</div>
        <div class="volume-container">
            <div
                class="volume-bar"
                style={`width: ${volume}px;`}
                class:medium={volume > mediumVolume}
                class:high={volume > highVolume}
            ></div>
        </div>
    </div>
{/if}

<style>
    .volume-container {
        width: 100px;
        height: 10px;
        background-color: var(--surface1);
        border-radius: 5px;
        margin-left: 4px;
    }

    .volume-bar {
        border-radius: 4px;
        height: 8px;
        background-color: var(--teal);
    }

    .volume-bar.medium {
        background-color: var(--peach);
    }

    .volume-bar.high {
        background-color: var(--red);
    }
</style>
