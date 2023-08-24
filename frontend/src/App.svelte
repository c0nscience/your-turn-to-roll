<script lang="ts">

  import SetupView from './lib/host/SetupView.svelte'
  import ClientJoinView from './lib/client/ClientJoinView.svelte'
  import { apiUrl } from './lib/api/config'
  import { sessionIdStore } from './lib/common/sessionStore'
  import { modeStore } from './lib/common/modeStore'
  import ClientFightView from './lib/client/ClientFightView.svelte'
  import HostFightView from './lib/host/HostFightView.svelte'

  const handleHostClicked = async () => {
    const resp = await fetch(`${apiUrl}/session/start`)
    const json = await resp.json()

    $sessionIdStore = json.id
    $modeStore = 'host'
  }

  const handleClientClicked = () => {
    $modeStore = 'join'
  }
</script>

<main class="container h-full mx-auto">
  {#if !$modeStore}
    <div class="h-full w-full grid grid-cols-1 content-center gap-4 px-5">
      <button class="btn" on:click={handleClientClicked}>Join</button>

      <button class="btn" on:click={handleHostClicked}>Host</button>
    </div>
  {/if}
  {#if $modeStore === 'host'}
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
