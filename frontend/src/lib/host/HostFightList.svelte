<script lang="ts">

  import { byInitiativeDesc, characterStore } from './characterStore'
  import { currentIdxStore } from './fightStore'
  import ConfirmationDialog from '../common/ConfirmationDialog.svelte'
  import { sessionIdStore } from '../common/sessionStore'
  import { sendUpdateToClients } from '../api/client'
  import { save } from '../api/client.js'
  import EyeClosed from '../../assets/EyeClosed.svelte'
  import Eye from '../../assets/Eye.svelte'

  $: entries = [...$characterStore].sort(byInitiativeDesc)
  $: entries, sendUpdateToClients(entries, $currentIdxStore, $sessionIdStore)

  $: isTheTurnOf = (idx: number): boolean => {
    return $currentIdxStore === idx
  };

  let idToDelete = ''
  let dialog: HTMLDialogElement

  const deleteCharacter = (id: string) => {
    $characterStore = $characterStore.filter(character => character.id !== id)
    save($characterStore, $sessionIdStore)
  }

  const handleDeleteClicked = (id: string) => {
    idToDelete = id
    dialog.showModal()
  }

  const handleHideClicked = (id: string) => {
    $characterStore = $characterStore.map(character => {
      if (character.id === id) {
        character.hidden = !character.hidden
      }
      return character
    })

    save($characterStore, $sessionIdStore)
  }
</script>

<table class="table mb-5">
  <!-- head -->
  <thead>
  <tr>
    <th></th>
    <th>Name</th>
    <th class="w-40"></th>
  </tr>
  </thead>
  <tbody>
  <!-- row 1 -->
  {#each entries as character, i (character.id)}
    <tr class:bg-accent={isTheTurnOf(i)}>
      <th class:text-accent-content={isTheTurnOf(i)}>{character.initiative}</th>
      <td class="text-xl" class:text-accent-content={isTheTurnOf(i)}>{character.name}</td>
      <td class="grid grid-flow-col justify-between">
        <button class="btn btn-ghost"
                class:text-accent-content={isTheTurnOf(i)}
                on:click={() => {handleHideClicked(character.id)}}>
          {#if (character.hidden)}
            <EyeClosed/>
          {:else }
            <Eye/>
          {/if}
        </button>
        <button class="btn btn-ghost"
                class:text-accent-content={isTheTurnOf(i)}
                on:click={() => {handleDeleteClicked(character.id)}}>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
            <path fill-rule="evenodd"
                  d="M16.5 4.478v.227a48.816 48.816 0 013.878.512.75.75 0 11-.256 1.478l-.209-.035-1.005 13.07a3 3 0 01-2.991 2.77H8.084a3 3 0 01-2.991-2.77L4.087 6.66l-.209.035a.75.75 0 01-.256-1.478A48.567 48.567 0 017.5 4.705v-.227c0-1.564 1.213-2.9 2.816-2.951a52.662 52.662 0 013.369 0c1.603.051 2.815 1.387 2.815 2.951zm-6.136-1.452a51.196 51.196 0 013.273 0C14.39 3.05 15 3.684 15 4.478v.113a49.488 49.488 0 00-6 0v-.113c0-.794.609-1.428 1.364-1.452zm-.355 5.945a.75.75 0 10-1.5.058l.347 9a.75.75 0 101.499-.058l-.346-9zm5.48.058a.75.75 0 10-1.498-.058l-.347 9a.75.75 0 001.5.058l.345-9z"
                  clip-rule="evenodd"/>
          </svg>
        </button>
      </td>
    </tr>
  {/each}
  </tbody>
</table>

<ConfirmationDialog bind:dialog
                    id="confirm-delete"
                    msg="Do you really want to delete this entry?"
                    on:confirm={() => {
                      deleteCharacter(idToDelete)
                      dialog.close()
                      idToDelete = ''
                    }}/>
