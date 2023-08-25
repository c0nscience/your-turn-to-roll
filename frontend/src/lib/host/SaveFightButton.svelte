<script>
    import {characterStore} from "./characterStore";
    import {apiUrl} from "../api/config";
    import {sessionIdStore} from "../common/sessionStore";

    const handleClick = async () => {
        await fetch(`${apiUrl}/session/${$sessionIdStore}`, {
            method: 'POST',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                characters: $characterStore,
            }),
        })
    }
</script>

<button disabled={$characterStore.length === 0} on:click={handleClick} class="btn btn-secondary">Save</button>
