<script lang="ts">

  import { connect } from './displayStore'
  import { modeStore } from '../common/modeStore'
  import { sessionIdStore } from '../common/sessionStore'
  import ConfirmationDialog from '../common/ConfirmationDialog.svelte'

  let sessionId: number | null = null
  let dialog: HTMLDialogElement

  const handleJoin = () => {
    connect(sessionId!)
      .then(() => {
        $sessionIdStore = sessionId!
        $modeStore = 'client-fight'
      })
      .catch(() => {
        dialog.showModal()
      })
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
    <button disabled={!sessionId} class="btn btn-success join-item rounded-r-full" on:click={handleJoin}>Join
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
