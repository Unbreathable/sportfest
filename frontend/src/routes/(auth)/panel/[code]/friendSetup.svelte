<script lang="ts">
	import { page } from "$app/stores";
	import { postRq } from "$lib/scripts/requests";
	import { onMount } from "svelte";

    let step = 0;
    let loading = true;
    let friendName = ""

    onMount(() => loadStatus())

    async function loadStatus() {
        step = 0;
        loading = true;

        const res = await postRq("/users/friends/status", JSON.stringify({
            "code": $page.params.code
        }))

        if(!res.success) {
            error(res.message);
            return;
        }

        if(res.set) {
            ownCode = res.code;
            code = res.enteredCode
            step = 1;
        }

        if(res.friends) {
            friendName = res.friend
            code = res.enteredCode
            step = 2
        }

        console.log(res)
    
        loading = false;
    }

    let copied = false;
    let ownCode = "";

    export let error: (msg: string) => void;

    function copyOwnCode() {
        navigator.clipboard.writeText(ownCode);
        copied = true;
        setTimeout(() => {
            copied = false;
        }, 2000);
    }

    let generateLoading = false;

    async function generateCode() {
        if(generateLoading) return;
        generateLoading = true;
        const res = await postRq("/users/friends/generate", JSON.stringify({
            "code": $page.params.code
        }))
        generateLoading = false;

        if(!res.success) {
            error(res.message);
            return;
        }

        loadStatus();
    }

    let saveLoading = false;
    let code = "";

    async function saveCode() {
        if(saveLoading) return;
        saveLoading = true;
        const res = await postRq("/users/friends/set", JSON.stringify({
            "code": $page.params.code,
            "friend_code": code
        }))
        saveLoading = false;

        if(!res.success) {
            error(res.message);
            return;
        }

        loadStatus();
    }

</script>

<!-- * Defining friends: Step 1 -->
<p class="text-medium text-hl">{loading ? "Laden.." : "Freund hinzufügen (optional)"}</p>
{#if step == 0}
    <p class="text-small">Wenn du mit einem Freund, der in deinem Jahrgang ist, zusammen in einem Team sein willst, könnt ihr euch hier gegenseitig hinzufügen. Wir würden es dir aber empfehlen zuerst das Video dazu anzuschauen!</p>
    <div class="row">
        <button class="button" on:click={generateCode}>{generateLoading ? "Wird generiert.." : "Starten"}</button>
        <button class="button button-secondary" on:click={() => window.open("https://youtube.com")}>Video anschauen</button>
    </div>
{/if}

<!-- * Defining friends: Step 2 -->
{#if step == 1}
    <p class="text-small">Hier ist dein Freundescode:</p>
    <div class="row">
        <input bind:value={ownCode} disabled>
        <button class="button {copied ? "button-success" : ""}" on:click={copyOwnCode}>{copied ? "Kopiert!" : "Kopieren"}</button>
    </div>
    <p class="text-small">Um deinen Freund jetzt hinzuzufügen, gebe bitte in diesem Feld <strong>den Freundescode deines Freundes</strong> an:</p>
    <div class="row">
        <input bind:value={code} placeholder="Freundescode deines Freundes">
        <button class="button" on:click={saveCode}>{saveLoading ? "Wird gespeichert.." : "Speichern"}</button>
    </div>

    <p class="text-small">Du kannst jetzt mit dem Knopf unter diesem Text nochmal überprüfen ob alles funktioniert hat, wenn ihr euch beide gegenseitig eingetragen habt. Du kannst den eingetragenen Freundescode jederzeit ändern, falls du doch jemand anderen hinzufügen willst.</p>
    <div class="row">
        <button class="button" on:click={loadStatus}>Status überprüfen</button>
        <button class="button button-secondary" on:click={() => window.open("https://youtube.com")}>Video anschauen</button>
    </div>
{/if}

<!-- * Defining friends: Step 3 -->
{#if step == 2}
    <p class="text-small">Du hast deinen Freund <strong>{friendName}</strong> erfolgreich hinzugefügt. Falls du deine Freundschaft ändern willst, kannst du nochmal den Knopf unter diesem Text drücken.</p>
    <div class="row">
        <button class="button" on:click={() => step = 1}>Freundschaft ändern</button>
        <button class="button button-secondary" on:click={() => window.open("https://youtube.com")}>Video anschauen</button>
    </div>
{/if}

<style lang="scss">
    @import '$lib/styles/comp.scss';

    .text-hl {
        margin-top: var(--def-spacing);
        font-weight: bold;
    }

    button {
        width: max-content;
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
</style>