<script lang="ts">
	import { goto } from "$app/navigation";
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import { currentYear, yearLoading } from "$lib/scripts/year";
	import { postRqAuthorized } from "$lib/scripts/requests";

    onMount(() => loadData())

    let loading = true;
    let error = "";

    let currentGame = 0;
    let matches: Map<number, any[]> = new Map();
    let games: any[] = []
    let count = 0;

    async function loadData() {
        currentGame = parseInt(localStorage.getItem("currentGame") ?? "0")
        loading = true;

        const res = await postRqAuthorized("/matches/list", JSON.stringify({
            "year": 0
        }))

        if(!res.success) {
            loading = false;
            return
        }

        res.games.forEach((game: any) => {
            res.matches.forEach((match: any) => {
                if(match.game == game.id) {
                    if(!matches.has(game.id)) matches.set(game.id, [])
                    matches.get(game.id)!.push(match)
                }
            })
        })
        games = res.games
        count = res.matches.length

        console.log(res)

        loading = false;
    }

    function setCurrentGame(id: number) {
        currentGame = id;
        localStorage.setItem("currentGame", id.toString())
    }

</script>

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100">
        <h3 class="headline">{loading ? "Spiele werden geladen.." : "Spiele (" + count + ")"}</h3>
        {#if !loading}
        <div transition:fade class="row">
            <button class="button button-small button-secondary {currentGame == 0 ? "button-selected" : ""}" on:click={() => setCurrentGame(0)}>Alles</button>
            {#each games as game}
            <button class="button button-small button-secondary {currentGame == game.id ? "button-selected" : ""}" on:click={() => setCurrentGame(game.id)}>{game.name}</button>
            {/each}
        </div>
        {/if}
    </div>

    {#if !loading}
    <div transition:fade class="years">
        {#each games as game}
            {#if currentGame == 0 || currentGame == game.id}
            <p class="text-medium hl">{game.name} ({(matches.get(game.id) ?? []).length})</p>
            {/if}
            {#if currentGame == 0 || currentGame == game.id}
            {#each matches.get(game.id) ?? [] as match}
            <div class="year">
                <div class="row">
                    <span class="material-icons icon-accent icon-medium">sports_volleyball</span>
                    <p class="text-medium">Match {match.id}, Teamgröße: {match.teamSize}</p>
                </div>
                <div class="row">
                    <button class="button button-small button-secondary" on:click={() => goto("/dash/matches/" + match.id)}>Überblick</button>
                    <button class="button button-small button-secondary">Löschen</button>
                </div>
            </div>
            {/each}
            {/if}
        {/each}
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
        margin-top: var(--def-spacing);
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
