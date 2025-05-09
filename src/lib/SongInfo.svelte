<script lang="ts">
    export let artist: string | null | undefined;
    export let title: string | null | undefined;
</script>

<div class="media flex-gap r-border">
    <div class="text-container italic flex-gap">
        <div class="subtext">{artist ?? ""}</div>
        <div class="title">{title ?? ""}</div>
    </div>
</div>

<style>
    .media {
        font-size: 8px;
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        min-width: 0;
        display: inline-block;
        max-width: 50%;
    }

    .title {
        color: var(--teal);
    }

    .text-container {
        padding-left: 100%;
        animation: marquee-bidirectional 10s linear infinite alternate;
    }

    @keyframes marquee-bidirectional {
        0% {
            transform: translateX(0);
        }
        100% {
            transform: translateX(-100%);
        }
    }
</style>
