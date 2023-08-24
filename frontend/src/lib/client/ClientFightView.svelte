<script lang="ts">

    import { sessionIdStore } from '../common/sessionStore'
    import { namesStore } from './displayStore'
</script>

<div class="grid grid-cols-1 p-5 h-full place-content-center">
    <h1 class="text-gray-500 text-2xl text-center">Session: {$sessionIdStore}</h1>

    <div class="overflow-x-auto">
        <table class="table">
            <!-- head -->
            <thead>
            <tr>
                <th></th>
            </tr>
            </thead>
            <tbody>
            {#each $namesStore as name, i (name)}
                <tr class={i === 0 ? 'bg-accent text-gray-900 text-6xl': 'text-2xl'}>
                    <th>{name}</th>
                </tr>
            {/each}
            </tbody>
        </table>
    </div>
</div>
