<script lang="ts">
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";
	import { page } from "$app/stores";

    onMount(() => loadData())

    let loading = true;
    let error = "";

    let match: any = {};
    let game: any = {};
    let team1: any[] = [], team2: any[] = [];

    async function loadData() {
        loading = true;

        const res = await postRqAuthorized("/matches/get", JSON.stringify({
            "match": parseInt($page.params.matchId)
        }))

        if(!res.success) {
            loading = false;
            return
        }

        match = res.match;
        game = res.game;
        team1 = res.team1;
        team2 = res.team2;

        console.log(res)

        loading = false;
    }

</script>

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100">
        <h3 class="headline">{loading ? "Match wird geladen.." : "Match " + match.id + " in " + game.name}</h3>
        {#if !loading}
        <div transition:fade class="row">
            <button class="button button-secondary">Match zusammenfügen</button>
        </div>
        {/if}
    </div>

    {#if !loading}
    <div transition:fade class="teams">
        <div class="team w-100">
            <p class="text-medium hl">Team 1</p>
            {#each team1 as acc}
                <div class="row">
                    <div class="title">
                        <span class="material-icons icon-small icon-accent">person</span>
                        <p class="text-small">{acc.name}</p>
                    </div>
                    <div class="buttons">
                        <button class="button button-small button-secondary">Profil</button>
                        <button class="button button-small">Entfernen</button>
                    </div>
                </div>
            {/each}
        </div>
        <div class="team w-100">
            <p class="text-medium hl">Team 2</p>
            {#each team2 as acc}
                <div class="row">
                    <div class="title">
                        <span class="material-icons icon-small icon-accent">person</span>
                        <p class="text-small">{acc.name}</p>
                    </div>
                    <div class="buttons">
                        <button class="button button-small button-secondary">Profil</button>
                        <button class="button button-small">Entfernen</button>
                    </div>
                </div>
            {/each}
        </div>
    </div>
    {/if}
</div>

<style lang="scss">
    @import "$lib/styles/comp.scss";

    .site {
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        gap: var(--section-spacing);
    }

    .years {
        display: flex;
        flex-direction: column;
        gap: var(--row-spacing);
    }

    .hl {
        font-weight: bold;
    }

    .w-100 {
        width: 100%;
    }

    .title {
        display: flex;
        gap: var(--line-spacing);
        align-items: center;
    }

    .team {
        display: flex;
        flex-direction: column;
        gap: var(--row-spacing);
    }

    .row {
        display: flex;
        gap: var(--def-spacing);
        align-items: center;
        justify-content: space-between;
    }

    .buttons {
        display: flex;
        gap: var(--def-spacing);
    }

    .teams {
        display: flex;
        gap: var(--section-spacing);
        align-items: center;
        justify-content: space-between;
    }

</style>
