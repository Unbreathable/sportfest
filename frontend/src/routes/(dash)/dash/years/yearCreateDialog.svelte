<script lang="ts">
	import Dialog from "$lib/comp/dialog.svelte";

    export let close: () => void;

</script>

<Dialog>
    <div class="header">
        <h3 class="headline">Jahrgang erstellen</h3>
        <button on:click={close} class="icon-button"><span class="material-icons hover-accent icon-large">close</span></button>
    </div>

    <div class="content">
        <p>Gebe dem Jahrgang einen Namen:</p>
        <input placeholder="5. Klasse">
        <button class="button">Erstellen</button>
    </div>
</Dialog>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';

    input {
        width: auto;
    }

</style>