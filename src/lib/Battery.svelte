<script lang="ts">
    import type { BatteryOutput } from "zebar";
    import IconValue from "./IconValue.svelte";

    export let battery: BatteryOutput;

    const iconClassIsCharging = "nf nf-md-lightning_bolt charging-icon";
    $: iconClassChargeAmount = `nf nf-fa-battery_${getBatteryIconValue(battery.chargePercent)}`;

    const batteryTimeString = (time: number | null): string => {
        if (time) {
            return `(${(time / (1000 * 60)).toFixed(0)} min)`;
        }
        return "";
    };

    const getBatteryIconValue = (percentage: number) => {
        if (percentage > 90) return "4";
        if (percentage > 70) return "3";
        if (percentage > 40) return "2";
        if (percentage > 20) return "1";
        return "0";
    };
</script>

<div class="flex-gap battery">
    {#if battery.isCharging}
        <i class={iconClassIsCharging}></i>
    {/if}
    <IconValue
        iconClass={iconClassChargeAmount}
        warningValue={30}
        severeWarningValue={15}
        warnAbove={false}
        isPercentage={true}
        value={Math.round(battery.chargePercent)}
    />
    <span class="subtext small"
        >{batteryTimeString(
            battery.isCharging ? battery.timeTillFull : battery.timeTillEmpty
        )}</span
    >
</div>

<style>
    .battery {
        gap: 2px;
        position: relative;
    }

    .charging-icon {
        position: absolute;
        font-size: 8px;
        left: -7px;
        top: -1px;
    }

    .small {
        font-size: 10px;
    }
</style>
