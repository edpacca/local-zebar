<script lang="ts">
  import type { NetworkGateway, NetworkOutput } from "zebar";

  export let network: NetworkOutput;

  const iconNameEthernet = "nf nf-fa-ethernet";
  const iconNameWifi = "wifi_strength";
  const iconNameNoConnection = "wifi_strength_off_outline";
  const iconNameError = "wifi_strength_alert_outline";
  const iconNameHidden = "nf nf-fa-dungeon";
  // const iconNameTrafficUpDown = "nf nf-fa-arrows_up_down";

  let isHidingIp = false;
  $: iconClass = getIconClass(network);
  $: ssid = network?.defaultGateway?.ssid;
  $: ipaddr = network?.defaultGateway?.ipv4Addresses
    ? network?.defaultGateway?.ipv4Addresses[0]
    : "";
  $: trafficDown = ((network?.traffic?.received.bytes ?? 0) / 1000).toFixed(0);
  $: trafficUp = ((network?.traffic?.transmitted.bytes ?? 0) / 1000).toFixed(0);

  const getIconClass = (networkOutput: NetworkOutput) => {
    switch (networkOutput.defaultInterface?.type) {
      case "ethernet":
        return iconNameEthernet;
      case "wifi":
        return wifiIcon(networkOutput.defaultGateway);
      default:
        return iconNameNoConnection;
    }
  };

  const wifiIconValue = (signalStrength: number | null): string => {
    if (!signalStrength || signalStrength < 20) {
      return "outline";
    } else if (signalStrength > 80) {
      return "4";
    } else if (signalStrength > 60) {
      return "3";
    } else if (signalStrength > 40) {
      return "2";
    } else {
      return "1";
    }
  };

  const wifiIcon = (defaultGateway: NetworkGateway | null) => {
    if (!defaultGateway) {
      return iconNameError;
    }
    return `${iconNameWifi}_${wifiIconValue(defaultGateway.signalStrength)}`;
  };

  const toggleHidingIp = () => {
    isHidingIp = !isHidingIp;
  };
</script>

<div class="flex-gap">
  <i class={`nf nf-md-${iconClass}`}></i>
  {#if ssid}
    <button onclick={toggleHidingIp}>
      {#if isHidingIp}
        <span><i class={iconNameHidden}></i></span>
      {:else}
        {ssid} | <span>{ipaddr}</span> |
      {/if}
    </button>
  {/if}
  <div class="traffic">
    {trafficUp} / {trafficDown}
  </div>
</div>

<style>
  span {
    color: var(--subtext0);
  }

  .flex-gap {
    gap: 4px;
  }

  .traffic {
    color: var(--teal);
    font-size: 8px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 0px;
  }

  button {
    background: none;
    border: none;
    border-radius: 2px;
  }

  button:hover {
    background: var(--surface0);
    color: var(--teal);
  }
</style>
