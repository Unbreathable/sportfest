<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    let loading = false;
    let error = "";

    export let game = {
        id: 0,
        name: ""
    };
    export let refresh: () => void;
    export let close: () => void;

    async function del() {
        if(loading) return;
        loading = true;

        const res = await postRqAuthorized("/games/delete", JSON.stringify({
            "id": game.id
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

<FormDialog title="{game.name} löschen" {error} {close}>

    <div class="content">
        <p class="text-medium">Wenn du dieses Spiel löschst, werden auch <strong>ALLE Matches, Auswahlen und Teams</strong> im Spiel gelöscht.</p>
        <button class="button" on:click={del}>{loading ? "Wird gelöscht.." : "Löschen"}</button>
    </div>
</FormDialog>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';

    input {
        width: auto;
    }

</style>