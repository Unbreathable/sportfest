<script lang="ts">
	import Dialog from "$lib/comp/dialog.svelte";
	import { slide } from "svelte/transition";

    let alert = false;
    let errBefore = "";
    export let error: string;
    $: {
        if(errBefore != "" && error != "") {
            alert = true;
            setTimeout(() => {
                alert = false;
            }, 500);
        }
        errBefore = error;
    }

    export let close: () => void;

</script>

<Dialog>
    <div class="header">
        <h3 class="headline">Jahrgang erstellen</h3>
        <button on:click={close} class="icon-button"><span class="material-icons hover-accent icon-large">close</span></button>
    </div>

    {#if error != ""}
    <p transition:slide class="error {alert ? "error-alert" : ""}">{error}</p>
    {/if}

    <slot />
</Dialog>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';

    .error {
        transition: 250ms transform ease;
    }

    .error-alert {
        transform: scale(1.1);
    }

</style>