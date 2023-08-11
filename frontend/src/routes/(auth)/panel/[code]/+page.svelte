<script lang="ts">
	import { page } from "$app/stores";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import { postRq } from "$lib/scripts/requests";
    import { onMount } from "svelte";

    let error = "";
    let loadingGames = true;
    let loadingFriends = false
    let chosen: any[] = [];
    let games: any[] = [];

    onMount(() => {
        console.log($page.params.code);
        loadGames();
    })

    async function loadGames() {
        loadingGames = true;
        const res = await postRq("/users/choices", JSON.stringify({
            "code": $page.params.code
        }));

        if(!res.success) {
            error = res.message;
            return;
        }

        chosen = res.choices;
        games = res.games;
        games.forEach(game => {
            game.contains = chosen.find((c: any) => c == game.id) != undefined;
        })

        loadingGames = false;
    }


    let copied = false;
    let ownCode = "5-W03-ABCDEF";

    function copyOwnCode() {
        navigator.clipboard.writeText(ownCode);
        copied = true;
        setTimeout(() => {
            copied = false;
        }, 2000);
    }

    let choiceLoading = false;
    let currentChoice = 0;
    let requestError = "";

    function toggleChoice(game: any) {
        if(choiceLoading) return;
        choiceLoading = true;
        currentChoice = game.id;
     
        if(game.contains) {
            removeChoice(game);
        } else {
            addChoice(game);
        }
    }

    async function addChoice(game: any) {
        const res = await postRq("/users/addchoice", JSON.stringify({
            "code": $page.params.code,
            "choice": game.id
        }));

        if(!res.success) {
            requestError = res.message;
            choiceLoading = false;
            return;
        }

        game.contains = true;
        chosen.push(game.id);
        choiceLoading = false;
        games = games;
    }

    async function removeChoice(game: any) {
        const res = await postRq("/users/removechoice", JSON.stringify({
            "code": $page.params.code,
            "choice": game.id
        }));

        if(!res.success) {
            requestError = res.message;
            choiceLoading = false;
            return;
        }

        game.contains = false;
        chosen = chosen.filter(c => c != game.id);
        choiceLoading = false;
        games = games;
    }

</script>

<svelte:head>
    <title>Sportfest: Einstellungen</title>
</svelte:head>

{#if requestError != ""}
<ErrorDialog error={requestError} close={() => requestError = ""} />
{/if}

<div class="center">
    <div class="form">
        <h3 class="headline">{loadingGames || loadingFriends ? "Wird geladen.." : "Auswahl für das Sportfest"}</h3>

        <div class="comp">

            {#if error != ""}
            <p class="error">{error}</p>
            {/if}

            {#if !loadingGames && !loadingFriends && error == ""}
            <p class="text-medium text-hl">Spiele auswählen</p>
            <p class="text-small">Bitte wähle <strong>3-5 Spiele</strong> aus, in denen du im Sportfest teilnehmen willst. Falls du einen Freund angegeben hast, wird versucht euch in den Spielen, die ihr beide ausgewählt habt, in ein Team zu bringen.</p>
            <div class="games">
                {#each games as game}
                <div class="game {game.contains ? "selected" : ""}">
                    <p class="text-medium">{game.name}</p>
                    <p class="text-small">{game.description}</p>
                    <div class="buttons">
                        <button class="button button-small {game.contains ? "button-selected" : ""}" on:click={() => toggleChoice(game)}>{currentChoice == game.id && choiceLoading ? "Laden.." : (game.contains ? "Auswahl aufheben" : "Auswählen")}</button>
                        <button class="button button-small button-secondary" on:click={() => window.open(game.video)}>Video anschauen</button>
                    </div>
                </div>
                {/each}
            </div>

            <p class="text-medium text-hl">Freund definieren</p>
            <p class="text-small">Gebe hier den Code eines Freundes ein:</p>
            <div class="row">
                <input placeholder="5-W03-ABCDEF">
                <button class="button">Hinzufügen</button>
            </div>
            <p class="text-small">Wenn du den Code deines Freundes eingeben hast, <strong>muss</strong> dein Freund auch deinen Code eingeben:</p>
            <div class="row">
                <input bind:value={ownCode} disabled>
                <button class="button {copied ? "button-success" : ""}" on:click={copyOwnCode}>{copied ? "Kopiert!" : "Kopieren"}</button>
            </div>
            <p class="text-small">Wenn ihr beide eure Codes bei euch gegenseitig eingetragen habt, kannst du hier den Status überprüfen.</p>
            <div class="row sb">
                <p class="text-medium">Der Code ist noch nicht bei deinem Freund eingetragen.</p>
                <button class="button">Aktualisieren</button>
            </div>
            {/if}
        </div>
        <div class="spacer"></div>
    </div>
</div>

<style lang="scss">
    @import '$lib/styles/comp.scss';

    .center {
        width: 100vw;
        height: calc(100vh - 2 * var(--section-spacing));
        overflow-y: scroll;
        padding: var(--section-spacing) 0;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .headline {
        text-align: center;
    }
    
    .row {
        display: flex;
        gap: var(--def-spacing);
        align-items: center;
        width: 100%;

        input {
            width: 100%;
        }
    }

    .spacer {
        padding: var(--def-spacing);
    }

    .sb {
        justify-content: space-between;
    }

    .text-hl {
        margin-top: var(--def-spacing);
        font-weight: bold;
    }

    .form {
        display: flex;
        align-items: center;
        flex-direction: column;
        gap: var(--section-spacing);
        padding: var(--section-spacing);
        border-radius: var(--def-spacing);
        width: 90%;
        height: 100%;
        max-width: 700px;

        .comp {
            display: flex;
            flex-direction: column;
            gap: var(--def-spacing);
            width: 100%;
        }
    }

    .games {
        display: flex;
        flex-direction: column;
        gap: var(--def-spacing);

        .game {
            display: flex;
            align-items: start;
            flex-direction: column;
            background-color: transparent;
            padding: var(--button-padding);
            gap: var(--def-spacing);
            border: 2px solid var(--secondary);
            border-radius: var(--def-spacing);
            text-align: start;

            .buttons {
                display: flex;
                align-items: center;
                width: 100%;
                gap: var(--def-spacing);
                justify-content: space-between;
            }
        }

        .selected {
            border: 2px solid var(--accent);
        }
    }

</style>