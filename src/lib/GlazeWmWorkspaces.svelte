<script lang="ts">
  import type { GlazeWmOutput } from "zebar";

  export let glazewm: GlazeWmOutput;

  const focusWorkspace = (name: string) => {
    const command = `focus --workspace ${name}`;
    glazewm?.runCommand(command);
  };
</script>

<div class="workspaces">
  {#each glazewm.currentWorkspaces as workspace}
    <button
      onclick={() => focusWorkspace(workspace.name)}
      class="workspace"
      class:focused={workspace.hasFocus}
      class:displayed={workspace.isDisplayed}
    >
      {workspace.displayName ?? workspace.name}
    </button>
  {/each}
</div>

<style>
  .workspaces {
    display: flex;
    align-items: center;
    gap: 4px;
    background: red;
  }

  .workspace {
    background: var(--surface0);
    color: var(--text);
    padding: 4px 8px;
    border: none;
    border-radius: 2px;
    cursor: pointer;
  }

  .displayed {
    background: var(--surface1);
  }

  .focused,
  .workspace:hover {
    background: var(--lavender);
    color: var(--base);
  }
</style>
