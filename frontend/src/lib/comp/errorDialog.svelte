<script lang="ts">
	import Dialog from "$lib/comp/dialog.svelte";

    export let close = () => {};
    export let error = "Ein Fehler ist aufgetreten";

</script>

<Dialog index={300}>
    <div class="header">
        <h3 class="headline">Fehler</h3>
        <button on:click={close} class="icon-button"><span class="material-icons hover-accent icon-large">close</span></button>
    </div>

    <div class="content">
        <p class="text-medium">{error}</p>
        <button on:click={close} class="button">Verstanden</button>
    </div>
</Dialog>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';

    input {
        width: auto;
    }

</style>