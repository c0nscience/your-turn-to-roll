<script lang="ts">

    import { connect } from './displayStore'
    import { modeStore } from '../common/modeStore'
    import { sessionIdStore } from '../common/sessionStore'

    let sessionId: number | null = null

    const handleJoin = () => {
        connect(sessionId!)
        $sessionIdStore = sessionId!
        $modeStore = 'client-fight'
    }
</script>

<div class="grid grid-cols-1 p-5 h-full place-content-center">
    <h1 class="text-4xl text-accent col-span-1 text-center mb-5">Join the session</h1>
    <div class="join col-span-1 place-self-center">
        <input bind:value={sessionId} placeholder="Session Id" type="number"
               class="input input-bordered join-item">
        <button disabled={!sessionId} class="btn btn-success join-item rounded-r-full" on:click={handleJoin}>Join
        </button>
    </div>
</div>
