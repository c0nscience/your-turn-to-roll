<script lang="ts">

    import { characterStore } from './characterStore'

    $: entries = [...$characterStore].sort((a, b) => b.initiative - a.initiative)

    const deleteCharacter = (id: string) => {
        $characterStore = $characterStore.filter(character => character.id !== id)
    }
</script>

{#each entries as character (character.id)}
    <div role="listitem">{character.name}</div><button on:click={() => {deleteCharacter(character.id)}}>del</button>
{/each}
