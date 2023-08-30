<script>
  import CharacterForm from "./CharacterForm.svelte";
  import CharacterSetupList from "./CharacterSetupList.svelte";
  import StartFightButton from "./StartFightButton.svelte";
  import SaveFightButton from "./SaveFightButton.svelte";
  import HostHeader from "./HostHeader.svelte";
</script>

<div class="p-5">
  <HostHeader/>

  <CharacterForm/>

  <CharacterSetupList/>

  <div class="join grid justify-items-stretch grid-cols-2 gap-4">
    <SaveFightButton/>
    <StartFightButton/>
  </div>
</div>
