<script lang="ts">
    let capsLockOn = false;

    const onkeydown = (event: KeyboardEvent) => {
        capsLockOn = event.getModifierState('CapsLock');
    }
</script>

<svelte:window on:keydown={onkeydown}/>
{#if capsLockOn}
    <button>CAPS</button>
{/if}

<style>
    button {
        background-color: var(--peach);
    }
</style>
