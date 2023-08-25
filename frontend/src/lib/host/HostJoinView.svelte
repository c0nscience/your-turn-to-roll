<script lang="ts">

    import { apiUrl } from '../api/config'
    import { sessionIdStore } from '../common/sessionStore'
    import { modeStore } from '../common/modeStore'
    import ConfirmationDialog from '../common/ConfirmationDialog.svelte'

    let sessionId: number | null = null
    let dialog: HTMLDialogElement
    const handleJoin = async () => {
        if (!sessionId) {
            const resp = await fetch(`${apiUrl}/session/start`)
            const json = await resp.json()
            $sessionIdStore = json.id
            $modeStore = 'setup'
        } else {
            try {
                const resp = await fetch(`${apiUrl}/session/${sessionId}/continue`)
                const json = await resp.json()
                $sessionIdStore = json.id
                $modeStore = 'setup'
            } catch {
                dialog.showModal()
            }
        }
    }
</script>

<div class="grid grid-cols-1 p-5 h-full place-content-center">
    <h1 class="text-4xl text-accent col-span-1 text-center mb-5">Join the session</h1>
    <div class="join col-span-1 place-self-center">
        <input bind:value={sessionId}
               placeholder="Session Id"
               type="number"
               inputmode="numeric"
               min="0"
               class="input input-bordered join-item">
        <button class="btn btn-success join-item rounded-r-full" on:click={handleJoin}>
            {#if !sessionId}
                New
            {:else}
                Continue
            {/if}
        </button>
    </div>
</div>

<ConfirmationDialog bind:dialog id="connection-failed"
                    headline="Not found"
                    msg="Session does not exist."
                    showCancelBtn={false}
                    on:confirm={() => {
                      dialog.close()
                    }}/>
