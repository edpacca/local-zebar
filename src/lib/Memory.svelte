<script lang="ts">
    import type { MemoryOutput } from "zebar";
    import IconValue from "./IconValue.svelte";
    export let memory: MemoryOutput
    const iconClassMemory = "nf nf-md-chip"
</script>

<IconValue
    iconClass={iconClassMemory}
    value={Math.round(memory.usage)}
    isPercentage={true}
    severeWarningValue={95}
/>
