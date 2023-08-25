<script lang="ts">

    import { characterStore, charactersToSend } from './characterStore'
    import { apiUrl } from '../api/config'
    import { sessionIdStore } from '../common/sessionStore'
    import { modeStore } from '../common/modeStore'

    const handleClick = async () => {
        const payload = charactersToSend($characterStore)
        await fetch(`${apiUrl}/session/${$sessionIdStore}`, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                                     characters: $characterStore,
                                 }),
        })
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
