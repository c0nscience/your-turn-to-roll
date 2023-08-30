<script lang="ts">
  import { currentIdxStore } from './fightStore'
  import { characterStore } from './characterStore'
  import { modeStore } from '../common/modeStore'

  const handleNext = () => {
    $currentIdxStore = ($currentIdxStore + 1) % $characterStore.length
  };

  const handleBack = () => {
    if ($currentIdxStore === 0) {
      $currentIdxStore = $characterStore.length - 1
    } else {
      $currentIdxStore = $currentIdxStore - 1
    }
  };

  const handleStop = () => {
    $modeStore = 'setup'
  };
</script>

<div class="join grid justify-items-stretch grid-cols-6">
  <button on:click={handleBack} class="btn join-item col-span-2">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
      <path
        d="M9.195 18.44c1.25.713 2.805-.19 2.805-1.629v-2.34l6.945 3.968c1.25.714 2.805-.188 2.805-1.628V8.688c0-1.44-1.555-2.342-2.805-1.628L12 11.03v-2.34c0-1.44-1.555-2.343-2.805-1.629l-7.108 4.062c-1.26.72-1.26 2.536 0 3.256l7.108 4.061z"/>
    </svg>
  </button>
  <button on:click={handleNext} class="btn join-item btn-primary col-span-3">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
      <path
        d="M5.055 7.06c-1.25-.714-2.805.189-2.805 1.628v8.123c0 1.44 1.555 2.342 2.805 1.628L12 14.471v2.34c0 1.44 1.555 2.342 2.805 1.628l7.108-4.061c1.26-.72 1.26-2.536 0-3.256L14.805 7.06C13.555 6.346 12 7.25 12 8.688v2.34L5.055 7.06z"/>
    </svg>
  </button>
  <button on:click={handleStop} class="btn join-item btn-error">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
      <path fill-rule="evenodd" d="M4.5 7.5a3 3 0 013-3h9a3 3 0 013 3v9a3 3 0 01-3 3h-9a3 3 0 01-3-3v-9z"
            clip-rule="evenodd"/>
    </svg>
  </button>
</div>
