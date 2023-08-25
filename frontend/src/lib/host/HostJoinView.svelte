<script lang="ts">

    import { apiUrl } from '../api/config'
    import { sessionIdStore } from '../common/sessionStore'
    import { modeStore } from '../common/modeStore'
    import ConfirmationDialog from '../common/ConfirmationDialog.svelte'
    import { characterStore } from './characterStore'

    let sessionId: number | null = null
    let dialog: HTMLDialogElement
    let keyDialog: HTMLDialogElement
    let key: string | null = null
    const handleJoin = async () => {
        if (!sessionId || !key) {
            const resp = await fetch(`${apiUrl}/session/start`)
            const json = await resp.json()
            $sessionIdStore = json.id
            key = json.key
            keyDialog.showModal()
        } else {
            const resp = await fetch(`${apiUrl}/session/${sessionId}/continue`, {
                method: 'POST',
                mode: 'cors',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({key}),
            })

            if (resp.status === 404 || resp.status === 400) {
                dialog.showModal()
                return
            }

            const json = await resp.json()
            $sessionIdStore = json.id
            $characterStore = json.characters
            $modeStore = 'setup'
        }
    }
</script>

<div class="grid grid-cols-1 p-5 h-full place-content-center">
    <h1 class="text-4xl text-accent col-span-1 text-center mb-5">Join the session</h1>
    <div class="join col-span-1 place-self-center">
        <input bind:value={sessionId}
               placeholder="Id"
               type="number"
               inputmode="numeric"
               min="0"
               class="input input-bordered join-item w-24">
        <input bind:value={key}
               placeholder="Key"
               class="input input-bordered join-item w-28">
        <button disabled={(!sessionId || !key) && !(!sessionId && !key)}
                class="btn btn-success join-item rounded-r-full w-24"
                on:click={handleJoin}>
            {#if !sessionId && !key}
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

<ConfirmationDialog bind:dialog={keyDialog} id="key-prompt"
                    headline="This is your secret key:"
                    msg="'{key}' << Note it down!!!"
                    showCancelBtn={false}
                    on:confirm={() => {
                        $modeStore = 'setup'
                        keyDialog.close()
                    }}/>
