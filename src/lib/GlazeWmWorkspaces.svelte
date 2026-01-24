<script lang="ts">
  import type { GlazeWmOutput } from "zebar";
  import type { Workspace } from "glazewm";
  import WorkspaceNameIcon from "./WorkspaceNameIcon.svelte";

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
      class:displayed={workspace.isDisplayed}>
      <WorkspaceNameIcon workspace={workspace}/>
    </button>
  {/each}
</div>

<style>
  .workspaces {
    display: flex;
    align-items: center;
    gap: 4px;
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

