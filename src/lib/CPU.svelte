<script lang="ts">
  import type { CpuOutput } from "zebar";
  import IconValue from "./IconValue.svelte";
  export let cpu: CpuOutput;
  const iconClassCpu = "nf nf-oct-cpu";
</script>

<IconValue
  iconClass={iconClassCpu}
  value={Math.round(cpu.usage)}
  isPercentage={true}
  warningValue={65}
  severeWarningValue={85}
/>

