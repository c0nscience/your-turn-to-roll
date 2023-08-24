<script>
    import CharacterForm from "./CharacterForm.svelte";
    import CharacterSetupList from "./CharacterSetupList.svelte";
    import StartFightButton from "./StartFightButton.svelte";
    import {sessionIdStore} from "../common/sessionStore";
</script>

<div class="p-5">
    <h1 class="text-2xl pb-5">Session: {$sessionIdStore}</h1>

    <CharacterForm/>

    <CharacterSetupList/>

    <StartFightButton/>
</div>
