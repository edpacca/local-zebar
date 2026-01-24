<script lang="ts">
    import type { Workspace } from "glazewm";
    export let workspace: Workspace;
    const baseName = (workspace.displayName ?? workspace.name).split(" ");
    const name = baseName[0]
    const iconClass = baseName.length > 1 ? baseName[1] : undefined;
</script>

{name}
{#if iconClass}
   <i class={`nf ${iconClass}`}></i>
{/if}
