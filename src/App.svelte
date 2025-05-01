<script>
  import { onMount } from "svelte";
  import * as zebar from "zebar";
  import Network from "./lib/Network.svelte";

  const providers = zebar.createProviderGroup({
    network: { type: "network" },
    glazewm: { type: "glazewm" },
    cpu: { type: "cpu" },
    date: { type: "date", formatting: "EEE d MMM t" },
    battery: { type: "battery" },
    memory: { type: "memory" },
  });

  $: output = providers.outputMap;
  onMount(() => providers.onOutput(() => output = providers.outputMap));

</script>

<div>
  <div class="left">

  </div>
  <div class="center">

  </div>
  <div class="right">
    <Network networkOutput={output.network}/>
  </div>
</div>

<style>
</style>
