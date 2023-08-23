<script lang="ts">

    import { createEventDispatcher, onMount } from 'svelte'

    export let id: string
    export let msg: string

    let dialog: any

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
        <h3 class="font-bold text-lg">Confirmation</h3>
        <p data-testid="modal-text" class="py-4">{msg}</p>
        <div class="modal-action">
            <button data-testid="confirm-btn" class="btn">Ok</button>
            <button data-testid="cancel-btn" class="btn" on:click={handleCancel} type="button">Cancel</button>
        </div>
    </form>
</dialog>
