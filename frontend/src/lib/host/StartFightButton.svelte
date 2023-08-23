<script lang="ts">

    import { characterStore } from './characterStore'
    import { apiUrl } from '../api/config'
    import { sessionIdStore } from './sessionStore'
    import { modeStore } from './modeStore'

    const handleClick = async () => {
        const payload = [...$characterStore]
            .sort((a, b) => b.initiative - a.initiative)
            .slice(0, 2)
            .map(c => c.name)
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

<button disabled={$characterStore.length === 0} on:click={handleClick} class="btn btn-accent w-full">Start</button>
