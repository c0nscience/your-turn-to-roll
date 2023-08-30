<script lang="ts">

  import { sessionIdStore } from '../common/sessionStore'
  import { namesStore } from './displayStore'
  import { receive, send } from './transition'
  import { flip } from 'svelte/animate'
</script>

<div class="grid grid-cols-1 p-5 h-full place-content-center">
  <h1 class="text-gray-500 text-2xl text-center">Session: {$sessionIdStore}</h1>

  <div class="overflow-x-auto">
    {#if $namesStore.length > 0}
      <table class="table">
        <!-- head -->
        <thead>
        <tr>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {#each $namesStore as name, i (name)}
          <tr class={i === 0 ? 'bg-accent text-gray-900 text-5xl': 'text-2xl'}
              in:receive={{key: name}}
              out:send={{key: name}}
              animate:flip={{duration: 400}}
          >
            <th>{name}</th>
          </tr>
        {/each}
        </tbody>
      </table>
    {:else}
      <p class="text-center text-primary-content text-6xl overflow-hidden">Roll for initiative.</p>
    {/if}
  </div>
</div>
