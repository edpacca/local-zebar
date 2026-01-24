<script lang="ts">
    import { checkThreshold } from "../utils";

    export let iconClass: string;
    export let value: number;
    export let isPercentage: boolean = false;
    export let warningValue: number | undefined = undefined;
    export let severeWarningValue: number | undefined = undefined;
    export let warnAbove = true;

    $: isWarning =
        warningValue &&
        checkThreshold(value, warningValue, warnAbove) &&
        !isSevereWarning;

    $: isSevereWarning =
        severeWarningValue &&
        checkThreshold(value, severeWarningValue, warnAbove);
</script>

<div class="icon-value">
    <i
        class={iconClass}
        class:warn={isWarning}
        class:severe-warning={isSevereWarning}
    ></i>
    <span class:warn={isWarning} class:severe-warning={isSevereWarning}>
        {value}{#if isPercentage}%{/if}
    </span>
</div>

<style>
    .icon-value {
        gap: 2px;
        display: grid 1fr auto;
        align-items: center;
        /* display: flex;
        flex-direction: row; */
    }
</style>
