<script lang="ts">

  import { createEventDispatcher } from 'svelte'

  export let id: string
  export let msg: string
  export let headline: string = 'Confirmation'
  export let showCancelBtn: boolean = true
  export let dialog: HTMLDialogElement

  const dispatcher = createEventDispatcher()

  const handleSubmit = () => {
    dispatcher('confirm')
  }

  function handleCancel() {
    dialog.close()
  }
</script>

<dialog bind:this={dialog} data-testid={id} id={id} class="modal">
  <form method="dialog" class="modal-box" on:submit|preventDefault={handleSubmit}>
    <h3 class="font-bold text-lg">{headline}</h3>
    <p data-testid="modal-text" class="py-4">{msg}</p>
    <div class="modal-action">
      <button data-testid="confirm-btn" class="btn btn-success">Ok</button>
      {#if showCancelBtn}
        <button data-testid="cancel-btn" class="btn btn-error" on:click={handleCancel} type="button">Cancel</button>
      {/if}
    </div>
  </form>
</dialog>
