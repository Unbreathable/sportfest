<script lang="ts">
	import { onMount } from "svelte";

    let copied = false;
    let ownCode = "5-W03-ABCDEF";

    onMount(() => {
        console.log("Hello World!");
    })

    function copyOwnCode() {
        navigator.clipboard.writeText(ownCode);
        copied = true;
        setTimeout(() => {
            copied = false;
        }, 2000);
    }
</script>

<div class="center">
    <div class="form shadow">
        <h3 class="headline">Sportfest: Einstellungen</h3>
        <div class="comp">
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

            <p class="text-medium text-hl">Spiele auswählen</p>
            <p class="text-small">Bitte wähle <strong>3-5 Spiele</strong> aus, in denen du im Sportfest teilnehmen willst. Falls du einen Freund angegeben hast, wird versucht euch in den Spielen, die ihr beide ausgewählt habt, in ein Team zu bringen.</p>
            <div class="games">
                <div class="game">
                    <p class="text-medium">Völkerball</p>
                    <p class="text-small">Ganzes normales Völkerball ohne irgendwelche Veränderungen. Wenn du das Spiel nicht kennst, kannst du trotzdem das Video dazu anschauen!</p>
                    <div class="buttons">
                        <button class="button button-small">Auswählen</button>
                        <button class="button button-small button-secondary">Video anschauen</button>
                    </div>
                </div>

                <div class="game selected">
                    <p class="text-medium">Battle Box</p>
                    <p class="text-small">Falls dir der Name des Spiel etwas sagt, weil er aus einem Minecraft-Event kommt, weißt du ja schon in welche Richtung es geht. Wir haben es allerdings noch ein bisschen verändert! Empfohlen ist auf jeden Fall vor dem Auswählen das Video anzuschauen!</p>
                    <div class="buttons">
                        <button class="button button-selected button-small">Auswahl aufheben</button>
                        <button class="button button-small button-secondary">Video anschauen</button>
                    </div>
                </div>

                <div class="game">
                    <p class="text-medium">Capture The Cone</p>
                    <p class="text-small">Eine veränderte Version von dem bekannten Spiel "Capture The Flag", die mal bei uns im Sportunterricht gespielt wurde. Hier auch am besten das Video für die Erklärung anschauen!</p>
                    <div class="buttons">
                        <button class="button button-small">Auswählen</button>
                        <button class="button button-small button-secondary">Video anschauen</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<style lang="scss">
    @import '$lib/styles/comp.scss';

    .center {
        margin: var(--section-spacing) 0;
        width: 100vw;
        display: flex;
        align-items: center;
        justify-content: center;
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
        justify-content: center;
        flex-direction: column;
        gap: var(--section-spacing);
        padding: var(--modal-padding);
        border-radius: var(--def-spacing);
        width: 90%;
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