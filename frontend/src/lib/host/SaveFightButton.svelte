<script lang="ts">
  import { characterStore } from './characterStore'
  import { sessionIdStore } from '../common/sessionStore'
  import { save } from '../api/client'

  const handleClick = async () => {
    await save($characterStore, $sessionIdStore)
  }
</script>

<button disabled={$characterStore.length === 0} on:click={handleClick} class="btn btn-secondary">Save</button>
