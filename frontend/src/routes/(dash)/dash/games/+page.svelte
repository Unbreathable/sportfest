<script lang="ts">
	import { goto } from "$app/navigation";
	import { fade } from "svelte/transition";
	import { postRqAuthorized } from "$lib/scripts/requests";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import GameDeleteDialog from "./gameDeleteDialog.svelte";
	import GameCreateDialog from "./gameCreateDialog.svelte";

    let showYearCreateDialog = false;
    const gameToDeleteDef = {
        id: 0,
        name: ""
    }
    let gameToDelete = gameToDeleteDef;

    let loading = true;
    let error = "";

    let games: any[] = [];

    onMount(() => loadData())

    async function loadData() {
        loading = true;
        
        const res = await postRqAuthorized("/games/list", "");
        if(!res.success) {
            error = res.message
            return
        }

        loading = false;
        games = res.obj;
    }

</script>

{#if gameToDelete.id != gameToDeleteDef.id}
<GameDeleteDialog game={gameToDelete} refresh={() => loadData()} close={() => gameToDelete = gameToDeleteDef} />
{/if}

{#if showYearCreateDialog}
<GameCreateDialog close={() => showYearCreateDialog = false} refresh={() => loadData()} />
{/if}

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100 row-sb">
        <h3 class="headline">{loading ? "Auswählbare Spiele werden geladen.." : "Auswählbare Spiele"}</h3>
        <button on:click={() => showYearCreateDialog = true} class="button">Spiel erstellen</button>
    </div>

    {#if !loading}
    <div transition:fade class="years">
        {#if games.length == 0}
        <p class="text-medium">Keine Auswählbaren Spiele vorhanden</p>
        {/if}
        {#each games as game}
        <div class="row w-100 row-sb row-start">
            <div class="col-line">
                <div class="row">
                    <span class="material-icons icon-accent icon-medium">sports_volleyball</span>
                    <p class="text-medium"><strong>{game.name}</strong></p>
                </div>
                <div class="row row-start row-line">
                    <p class="text-small icon-primary">Beschreibung:</p>
                    <p class="text-small">{game.description}</p>
                </div>
                <div class="row row-start row-line">
                    <p class="text-small icon-primary">Minimale Teamgröße:</p>
                    <p class="text-small">{game.minTeamSize}</p>
                </div>
                <div class="row row-start row-line">
                    <p class="text-small icon-primary">Maximale Teamgröße:</p>
                    <p class="text-small">{game.maxTeamSize}</p>
                </div>
            </div>
            <div class="row">
                <button class="button button-small button-secondary" on:click={() => goto("/dash/games/" + game.id + "/members")}>Überblick</button>
                <button class="button button-small button-secondary" on:click={() => window.open(game.video)}>Video</button>
                <button class="button button-small button-secondary" on:click={() => gameToDelete = game}>Löschen</button>
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
