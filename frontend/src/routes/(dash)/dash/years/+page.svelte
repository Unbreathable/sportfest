<script>
	import YearCreateDialog from "./yearCreateDialog.svelte";

    let showYearCreateDialog = false;

</script>

{#if showYearCreateDialog}
<YearCreateDialog close={() => showYearCreateDialog = false} />
{/if}

<div class="site">
    <div class="row">
        <h3 class="headline">Jahrgänge</h3>
        <button on:click={() => showYearCreateDialog = true} class="button">Jahrgang erstellen</button>
    </div>
</div>

<style lang="scss">
    @import "$lib/styles/comp.scss";

    .site {
        width: 100%;
        height: 100%;
    }

    .row {
        width: 100%;
        display: flex;
        gap: var(--section-spacing);
        align-items: center;
        justify-content: space-between;
    }

</style>
