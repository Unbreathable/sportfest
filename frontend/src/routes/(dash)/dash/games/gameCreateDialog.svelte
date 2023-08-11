<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    let name = "", description = "", video = "", minTeamSize: number, maxTeamSize: number;

    let loading = false;
    let error = "";

    export let refresh: () => void;
    export let close: () => void;

    async function create() {
        if(loading) return;
        loading = true;

        const res = await postRqAuthorized("/games/create", JSON.stringify({
            "name": name,
            "description": description,
            "video": video,
            "minTeamSize": minTeamSize,
            "maxTeamSize": maxTeamSize
        }))

        if(!res.success) {
            loading = false;
            error = res.message;
            return;
        }

        refresh();
        close();
    }

</script>

<FormDialog title="Spiel erstellen" {error} {close}>

    <div class="content">
        <p class="text-small">Gebe dem Spiel einen Name:</p>
        <input bind:value={name} placeholder="Völkerball">
        <p class="text-small">Das Spiel braucht auch noch eine Beschreibung, die direkt unter dem Namen steht:</p>
        <textarea bind:value={description} placeholder="In diesem Spiel geht es um.." />
        <p class="text-small">Neben dem Auswahl Knopf auf der Seite für Schüler gibt es auch noch einen Knopf, der zu einem Video weiterleitet:</p>
        <input bind:value={video} placeholder="https://youtube.com">
        <p class="text-small">Jetzt muss nur noch die minimale und maximale Teamgröße definiert werden:</p>
        <input bind:value={minTeamSize} type="number" placeholder="Minimale Teamgröße">
        <input bind:value={maxTeamSize} type="number" placeholder="Maximale Teamgröße">

        <button on:click={create} class="button">{loading ? "Wird erstellt.." : "Erstellen"}</button>
    </div>
</FormDialog>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';

    input {
        width: auto;
    }

</style>