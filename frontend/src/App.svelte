<script lang="ts">

    import SetupView from './lib/host/SetupView.svelte'

    let mode: 'host' | null = null
</script>

<main class="container h-full mx-auto">
    {#if !mode}
        <div class="h-full w-full grid grid-cols-1 content-center gap-4 px-5">
            <button class="btn">Client</button>

            <button class="btn" on:click={() => mode = 'host'}>Host</button>
        </div>
    {/if}
    {#if mode === 'host'}
        <SetupView/>
    {/if}
</main>
