<script lang="ts">

    import SetupView from './lib/host/SetupView.svelte'
    import { apiUrl } from './lib/api/config'
    import { sessionIdStore } from './lib/host/sessionStore'
    import { modeStore } from './lib/host/modeStore'

    const handleHostClicked = async () => {
        const resp = await fetch(`${apiUrl}/session/start`)
        const json = await resp.json()

        $sessionIdStore = json.id
        $modeStore = 'host'
    }

    const handleClientClicked = () =>{
        $modeStore = 'client'
    }
</script>

<main class="container h-full mx-auto">
    {#if !$modeStore}
        <div class="h-full w-full grid grid-cols-1 content-center gap-4 px-5">
            <button class="btn" on:click={handleClientClicked}>Client</button>

            <button class="btn" on:click={handleHostClicked}>Host</button>
        </div>
    {/if}
    {#if $modeStore === 'host'}
        <SetupView/>
    {/if}
    {#if $modeStore === 'fight'}
        <h1 class="text-6xl">FIGHT</h1>
    {/if}
    {#if $modeStore === 'client'}
        <h1 class="text-6xl">Client</h1>
    {/if}
</main>
