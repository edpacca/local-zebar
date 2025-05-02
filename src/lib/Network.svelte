<script lang="ts">
    import type { NetworkGateway, NetworkOutput } from "zebar";

    export let network: NetworkOutput;

    const iconNameEthernet = "ethernet-cable";
    const iconNameWifi = "wifi_strength";
    const iconNameNoConnection = "wifi_strength_off_outline";
    const iconNameError = "wifi_strength_alert_outline";

    $: iconClass = getIconClass(network);
    $: ssid = network?.defaultGateway?.ssid;
    $: ipaddr = network?.defaultGateway?.ipv4Addresses
        ? network?.defaultGateway?.ipv4Addresses[0]
        : "";

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
</script>

<div class="network">
    <i class={`nf nf-md-${iconClass}`}></i>
    {ssid} |
    <span>
        {ipaddr}
    </span>
</div>

<style>
    span {
        color: var(--subtext0);
    }
</style>
