<script lang="ts">
  import type { Character } from './character.type'
  import { createEventDispatcher } from 'svelte'

  export let character: Character

  const dispatcher = createEventDispatcher()

  const handleDelete = () => {
      dispatcher('delete', {
          id: character.id
      })
  }
</script>

<td>
    <input placeholder="Initiative" bind:value={character.initiative} type="number" step="0.1" inputmode="numeric" class="input w-24"/>
</td>
<td><input placeholder="Name" bind:value={character.name} type="text" class="input w-full"/></td>
<td>
    <button class="btn btn-error" on:click={handleDelete}>del</button>
</td>
