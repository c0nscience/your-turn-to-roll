<script lang="ts">

    import { characterStore } from './characterStore'
    import ConfirmationDialog from '../common/ConfirmationDialog.svelte'

    $: entries = [...$characterStore].sort((a, b) => b.initiative - a.initiative)

    let idToDelete = ''
    let dialog: HTMLDialogElement

    const deleteCharacter = (id: string) => {
        $characterStore = $characterStore.filter(character => character.id !== id)
    }

    const handleDeleteClicked = (id: string) => {
        idToDelete = id
        dialog.showModal()
    }
</script>

<table class="table mb-5">
    <!-- head -->
    <thead>
    <tr>
        <th></th>
        <th>Name</th>
        <th></th>
    </tr>
    </thead>
    <tbody>
    <!-- row 1 -->
    {#each entries as character (character.id)}
        <tr>
            <th>{character.initiative}</th>
            <td>{character.name}</td>
            <td>
                <button class="btn btn-error" on:click={() => {handleDeleteClicked(character.id)}}>del</button>
            </td>
        </tr>
    {/each}
    </tbody>
</table>


<ConfirmationDialog bind:dialog id="confirm-delete" msg="Do you really want to delete this entry?" on:confirm={() => {
    deleteCharacter(idToDelete)
    dialog.close()
    idToDelete = ''
}}/>
