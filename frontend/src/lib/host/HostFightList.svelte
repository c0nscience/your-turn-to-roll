<script lang="ts">

  import { characterStore } from './characterStore'
  import { currentIdxStore } from './fightStore'

  $: entries = [...$characterStore].sort((a, b) => b.initiative - a.initiative)

  $: isTheTurnOf = (idx: number): boolean => {
    return $currentIdxStore === idx
  }
</script>

<table class="table mb-5">
  <!-- head -->
  <thead>
  <tr>
    <th></th>
    <th>Name</th>
  </tr>
  </thead>
  <tbody>
  <!-- row 1 -->
  {#each entries as character, i (character.id)}
    <tr class:bg-accent={isTheTurnOf(i)}>
      <th class:text-accent-content={isTheTurnOf(i)}>{character.initiative}</th>
      <td class="text-xl" class:text-accent-content={isTheTurnOf(i)}>{character.name}</td>
    </tr>
  {/each}
  </tbody>
</table>
