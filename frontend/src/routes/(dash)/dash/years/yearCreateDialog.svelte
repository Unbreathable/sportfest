<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    let name = "";
    let loading = false;
    let error = "";

    export let refresh: () => void;
    export let close: () => void;

    async function create() {
        if(loading) return;
        loading = true;

        const res = await postRqAuthorized("/years/create", JSON.stringify({
            "name": name
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

<FormDialog {error} {close}>

    <div class="content">
        <p>Gebe dem Jahrgang einen Namen:</p>
        <input bind:value={name} placeholder="5. Klasse">
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