<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    let loading = false;
    let error = "";

    export let year = {
        id: 0,
        name: ""
    };
    export let refresh: () => void;
    export let close: () => void;

    async function del() {
        if(loading) return;
        loading = true;

        const res = await postRqAuthorized("/years/delete", JSON.stringify({
            "id": year.id
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

<FormDialog title="{year.name} löschen" {error} {close}>

    <div class="content">
        <p class="text-medium">Wenn du diesen Jahrgang löschst, werden auch <strong>ALLE Benutzer und Spiele</strong> des Jahrgangs gelöscht.</p>
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