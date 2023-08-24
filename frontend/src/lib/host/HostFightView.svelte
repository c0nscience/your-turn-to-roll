<script lang="ts">
  import { sessionIdStore } from '../common/sessionStore'
  import HostFightList from './HostFightList.svelte'
  import HostFightActions from './HostFightActions.svelte'
</script>

<div class="p-5">
  <h1 class="text-2xl pb-5">Session: {$sessionIdStore}</h1>

  <HostFightList/>

  <HostFightActions/>
</div>
