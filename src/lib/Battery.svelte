<script lang="ts">
    import type { BatteryOutput } from "zebar";

    export let battery: BatteryOutput;

    const iconClassIsCharging = "nf nf-md-lightning_bolt charging-icon"
    $: iconClassChargeAmount = `nf nf-fa-battery_${getBatteryIconValue(battery.chargePercent)}`

    const getBatteryIconValue = (percentage: number) => {
        if (percentage > 90)
            return "4";
        if (percentage > 70)
            return "3";
        if (percentage > 40)
            return "2";
        if (percentage > 20)
            return "1";
        return "0";
    }
</script>

<div class="flex-gap battery">
    {#if battery.isCharging || true}
        <i class={iconClassIsCharging}></i>
    {/if}
    <i class={iconClassChargeAmount}></i>
    {Math.round(battery.chargePercent)}
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
    }
</style>
