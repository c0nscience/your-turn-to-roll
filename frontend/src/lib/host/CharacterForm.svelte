<script lang="ts">

  import { characterStore } from './characterStore'
  import { v4 as uuidv4 } from 'uuid'

  let name: string = ''
  let initiative: number | null = null

  $: disabled = name === '' || !initiative || initiative < 0

  const addCharacter = () => {
    $characterStore = [...$characterStore, { id: uuidv4(), name, initiative: initiative! }]
    name = ''
    initiative = null
  }
</script>

<form on:submit|preventDefault={addCharacter}>
<div class="grid grid-cols-8 gap-2 mb-5">
  <input placeholder="Name" bind:value={name} type="text" class="input input-bordered col-span-4"/>
  <input placeholder="Init" bind:value={initiative} type="number" class="input input-bordered col-span-2"/>

  <button {disabled} class="btn btn-accent col-span-2">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
      <path fill-rule="evenodd"
            d="M12 3.75a.75.75 0 01.75.75v6.75h6.75a.75.75 0 010 1.5h-6.75v6.75a.75.75 0 01-1.5 0v-6.75H4.5a.75.75 0 010-1.5h6.75V4.5a.75.75 0 01.75-.75z"
            clip-rule="evenodd"/>
    </svg>
  </button>
</div>
</form>
