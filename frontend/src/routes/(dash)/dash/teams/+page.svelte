<script lang="ts">
	import { goto } from "$app/navigation";
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import AlgorithmControls from "./algorithmControls.svelte";

    let loading = false;
    let error = "";

    let algorithmControls = false;

    onMount(() => loadData())

    async function loadData() {
        
    }

</script>

{#if algorithmControls}
<AlgorithmControls close={() => algorithmControls = false} />
{/if}

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100 row-sb">
        <h3 class="headline">{loading ? "Teams werden geladen.." : "Teameinteilung"}</h3>
        <button on:click={() => algorithmControls = true} class="button">Algorithmus</button>
    </div>

    {#if !loading}
    <div transition:fade class="years">

    </div>
    {/if}
</div>

<style lang="scss">
    @import "$lib/styles/comp.scss";

    .site {
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: var(--section-spacing);
    }

    .years {
        display: flex;
        flex-direction: column;
        gap: var(--def-spacing);
    }

    .width-100 {
        width: 100%;
    }

    .row {
        display: flex;
        gap: var(--def-spacing);
        align-items: center;
    }

    .row-sb {
        gap: var(--section-spacing);
        justify-content: space-between;
    }

    .row-start {
        align-items: start;
    }

    .row-line {
        gap: var(--line-spacing);
    }

    .col-line {
        display: flex;
        flex-direction: column;
        gap: var(--line-spacing);
    }

</style>
