<script lang="ts">
    import type { MediaSession } from "zebar";
    export let session: MediaSession;
</script>

{#if session}
    <div class="media flex-gap r-border">
        <!-- <div class="subtext">{`{${session.sessionId}}`}</div> -->
        <div class="italic flex-gap">
            <div class="title">{session.title}</div>
        </div>
    </div>
{/if}

<style>
    .media {
        font-size: 8px;
        flex-shrink: 2;
    }

    .title {
        color: var(--teal)
    }
</style>
