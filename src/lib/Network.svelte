<script lang="ts">
    import type { NetworkGateway, NetworkOutput } from "zebar";

    export let networkOutput: NetworkOutput | null;

    const iconNameEthernet = "ethernet-cable";
    const iconNameWifi = "wifi_strength_";
    const iconNameNoConnection = "wifi_strength_off_outline";

    let iconClass: string = iconNameNoConnection;

    $: () => {
        switch (networkOutput?.defaultInterface?.type) {
            case "ethernet":
                iconClass = iconNameEthernet;
            case "wifi":
                iconClass = wifiIcon(networkOutput.defaultGateway);
            default:
                iconClass = iconNameNoConnection;
        }
    }

    $: ssid = networkOutput?.defaultGateway?.ssid;
    $: ipaddr = networkOutput?.defaultGateway?.ipv4Addresses ? networkOutput?.defaultGateway?.ipv4Addresses[0] : "";

    const wifiIconValue = (signalStrength: number | null): string => {
        if (!signalStrength || signalStrength < 20) {
            return "outline";
        } else if (signalStrength > 80) {
            return "4";
        } else if (signalStrength > 60) {
            return "3"
        } else if (signalStrength > 40) {
            return "2"
        } else {
            return "1"
        }
    }

    const wifiIcon = (defaultGateway: NetworkGateway | null) => {
        if (!defaultGateway) {
            return `${iconNameWifi}_alert_outline`;
        }
        return `${iconNameWifi}_${wifiIconValue(defaultGateway.signalStrength)}`
    }
</script>

{#if networkOutput}
    <div class="network">
        <i class={`nf nf-md-${iconClass}`}></i>
        {ssid} | {ipaddr}
    </div>
{/if}
