<script lang="ts">
    export let iconClass: string;
    export let value: number;
    export let isPercentage: boolean = false;
    export let warningValue: number | undefined = undefined;
    export let alertValue: number | undefined = undefined;
</script>

<div
    class:warning={warningValue && value > warningValue}
    class:alert={alertValue && value > alertValue}
>
    <i class={iconClass}></i>
        {value}
        {#if isPercentage}%{/if}
</div>
