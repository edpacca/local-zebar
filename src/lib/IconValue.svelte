<script lang="ts">
    export let iconClass: string;
    export let value: number;
    export let isPercentage: boolean = false;
    export let warningValue: number | undefined = undefined;
    export let alertValue: number | undefined = undefined;
    export let warnAbove = true;

    $: isWarning =
        warningValue &&
        ((warnAbove && value >= warningValue) ||
            (!warnAbove && value <= warningValue));

    $: isAlerting =
        alertValue &&
        ((warnAbove && value >= alertValue) ||
            (!warnAbove && value <= alertValue));
</script>

<div class="icon-value">
    <i class={iconClass} class:warn={isWarning} class:alert={isAlerting}></i>
    <span class:warn={isWarning} class:alert={isAlerting}>
        {value}{#if isPercentage}%{/if}
    </span>
</div>

<style>
    .icon-value {
        gap: 2px;
        display: flex;
        flex-direction: row;
    }
</style>
