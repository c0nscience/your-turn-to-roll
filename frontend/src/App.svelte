<script lang="ts">

  import SetupView from './lib/host/SetupView.svelte'
  import ClientJoinView from './lib/client/ClientJoinView.svelte'
  import { modeStore } from './lib/common/modeStore'
  import ClientFightView from './lib/client/ClientFightView.svelte'
  import HostFightView from './lib/host/HostFightView.svelte'
  import HostJoinView from './lib/host/HostJoinView.svelte'

  const handleHostClicked = () => {
        $modeStore = 'join-setup'
    }

    const handleClientClicked = () => {
        $modeStore = 'join'
    }
    console.log('search', window.location.search)
    console.log('pathname', window.location.pathname)
</script>

<main class="container h-full mx-auto md:max-w-2xl">
    {#if !$modeStore}
        <div class="h-full w-full grid grid-cols-1 content-center gap-4 px-5">
            <button class="btn" on:click={handleClientClicked}>Join</button>

            <button class="btn" on:click={handleHostClicked}>Host</button>
        </div>
    {/if}
    {#if $modeStore === 'join-setup'}
        <HostJoinView/>
    {/if}
    {#if $modeStore === 'setup'}
        <SetupView/>
    {/if}
    {#if $modeStore === 'fight'}
        <HostFightView/>
    {/if}
    {#if $modeStore === 'join'}
        <ClientJoinView/>
    {/if}
    {#if $modeStore === 'client-fight'}
        <ClientFightView/>
    {/if}
</main>
