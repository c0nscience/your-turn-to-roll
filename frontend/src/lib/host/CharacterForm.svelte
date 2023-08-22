<script lang="ts">

    import { characterStore } from './characterStore'
    import {v4 as uuidv4} from 'uuid'

    let name: string = ''
    let initiative: number | null = null

    $: disabled = name === '' || !initiative || initiative < 0

    const addCharacter = () => {
        $characterStore = [{id: uuidv4(), name, initiative: initiative!}]
        name = ''
        initiative = null
    }
</script>

<div class="form-control w-full max-w-xs">
    <label class="label">
        <span class="label-text">Name</span>
        <input bind:value={name} type="text" class="input input-bordered w-full max-w-xs"/>
    </label>
</div>

<div class="form-control w-full max-w-xs">
    <label class="label">
        <span class="label-text">Initiative</span>
        <input bind:value={initiative} type="number" class="input input-bordered w-full max-w-xs"/>
    </label>
</div>
<span>name: '{name}'</span>
<span>init: '{initiative}'</span>
{disabled}
<button {disabled} class="btn" on:click={addCharacter}>Add</button>
