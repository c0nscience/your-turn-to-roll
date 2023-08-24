<script lang="ts">
  import { currentIdxStore } from './fightStore'
  import { characterStore, charactersToSend } from './characterStore'
  import { apiUrl } from '../api/config'
  import { sessionIdStore } from '../common/sessionStore'
  import { modeStore } from '../common/modeStore'

  $: $currentIdxStore, currentIdxChanged()

  const currentIdxChanged = async () => {
    const payload = charactersToSend($characterStore, $currentIdxStore)

    await fetch(`${apiUrl}/fight/${$sessionIdStore}/update`, {
      method: 'POST',
      mode: 'cors',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({payload}),
    })

  }

  const handleNext = () => {
    $currentIdxStore = ($currentIdxStore + 1) % $characterStore.length
  }

  const handleBack = () => {
    if ($currentIdxStore === 0) {
      $currentIdxStore = $characterStore.length - 1
    } else {
      $currentIdxStore = $currentIdxStore - 1
    }
  }

  const handleStop = () => {
    $modeStore = 'setup'
  }
</script>

<div class="join grid justify-items-stretch grid-cols-6">
  <button on:click={handleBack} class="btn join-item col-span-2">&#9668;&#9668;</button>
  <button on:click={handleNext} class="btn join-item btn-primary col-span-3">&#9658;&#9658;</button>
  <button on:click={handleStop} class="btn join-item btn-error">&#9608;</button>
</div>
