<script lang="ts">
	import { goto } from "$app/navigation";
	import { fade } from "svelte/transition";
	import YearCreateDialog from "./yearCreateDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import YearDeleteDialog from "./yearDeleteDialog.svelte";

    let showYearCreateDialog = false;
    const yearToDeleteDef = {
        id: 0,
        name: ""
    }
    let yearToDelete = yearToDeleteDef;

    let loading = true;
    let error = "";

    let years: any[] = [];

    onMount(() => loadData())

    async function loadData() {
        loading = true;
        
        const res = await postRqAuthorized("/years/list", "");
        if(!res.success) {
            error = res.message
            return
        }

        loading = false;
        years = res.obj;
    }

</script>

{#if yearToDelete.id != yearToDeleteDef.id}
<YearDeleteDialog year={yearToDelete} refresh={() => loadData()} close={() => yearToDelete = yearToDeleteDef} />
{/if}

{#if showYearCreateDialog}
<YearCreateDialog close={() => showYearCreateDialog = false} refresh={() => loadData()} />
{/if}

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100">
        <h3 class="headline">{loading ? "Jahrgänge werden geladen.." : "Jahrgänge"}</h3>
        <button on:click={() => showYearCreateDialog = true} class="button">Jahrgang erstellen</button>
    </div>

    {#if !loading}
    <div transition:fade class="years">
        {#if years.length == 0}
        <p class="text-medium">Keine Jahrgänge vorhanden</p>
        {/if}
        {#each years as year}
        <div class="year">
            <div class="row">
                <span class="material-icons icon-accent icon-medium">school</span>
                <p class="text-medium">{year.name}</p>
            </div>
            <div class="row">
                <button class="button button-small button-secondary" on:click={() => goto("/dash/years/" + year.id + "/members")}>Überblick</button>
                <button class="button button-small button-secondary" on:click={() => yearToDelete = year}>Löschen</button>
            </div>
        </div>
        {/each}
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
        gap: var(--row-spacing);
    }

    .width-100 {
        width: 100%;
    }

    .row {
        display: flex;
        gap: var(--def-spacing);
        align-items: center;
        justify-content: space-between;
    }

    .year {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

</style>
