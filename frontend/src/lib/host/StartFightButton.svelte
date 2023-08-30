<script lang="ts">

    import { characterStore, charactersToSend } from './characterStore'
    import { apiUrl } from '../api/config'
    import { sessionIdStore } from '../common/sessionStore'
    import { modeStore } from '../common/modeStore'
    import { save } from '../api/client'

    const handleClick = async () => {
        const payload = charactersToSend($characterStore)
      await save($characterStore, $sessionIdStore)
        await fetch(`${apiUrl}/fight/start`, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                                     sessionId: $sessionIdStore,
                                     payload,
                                 }),
        })

        $modeStore = 'fight'
    }
</script>

<button disabled={$characterStore.length === 0} on:click={handleClick} class="btn btn-primary">Start</button>
