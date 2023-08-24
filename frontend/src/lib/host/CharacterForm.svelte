<script lang="ts">

  import { characterStore } from './characterStore'
  import { v4 as uuidv4 } from 'uuid'

  let name: string = ''
  let initiative: number | null = null

  $: disabled = name === '' || !initiative || initiative < 0

  const addCharacter = () => {
    $characterStore = [...$characterStore, {id: uuidv4(), name, initiative: initiative!}]
    name = ''
    initiative = null
  }
</script>

<div class="grid grid-cols-8 gap-2 mb-5">
  <input placeholder="Name" bind:value={name} type="text" class="input input-bordered col-span-3"/>
  <input placeholder="Initiative" bind:value={initiative} type="number" class="input input-bordered col-span-3"/>

  <button {disabled} class="btn btn-accent col-span-2" on:click={addCharacter}>Add</button>
</div>
