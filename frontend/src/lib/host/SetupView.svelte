<script>
  import CharacterForm from "./CharacterForm.svelte";
  import CharacterSetupList from "./CharacterSetupList.svelte";
  import StartFightButton from "./StartFightButton.svelte";
  import {sessionIdStore} from "../common/sessionStore";
  import SaveFightButton from "./SaveFightButton.svelte";
</script>

<div class="p-5">
    <h1 class="text-2xl pb-5">Session: {$sessionIdStore}</h1>

    <CharacterForm/>

    <CharacterSetupList/>

    <div class="join grid justify-items-stretch grid-cols-2 gap-4">
        <SaveFightButton/>
        <StartFightButton/>
    </div>
</div>
