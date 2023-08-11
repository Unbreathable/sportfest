<script lang="ts">
	import { fade } from "svelte/transition";
	import { postRqAuthorized } from "$lib/scripts/requests";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import { goto } from "$app/navigation";

    let loading = true;
    let error = "";

    let searchInput: string = "";
    let users: any[] = [];

    onMount(() => loadData(""))

    async function loadData(query: string) {
        loading = true;
        
        const res = await postRqAuthorized("/users/search", JSON.stringify({
            "query": query
        }));
        if(!res.success) {
            error = res.message
            return
        }

        loading = false;
        users = res.obj;
    }

    function search(e: any) {
        if(e.target.value == "") {
            loadData("");
        }
    }

    function resetSearch() {
        searchInput = "";
        loadData("");
    }

    function inputKeyEvent(e: any) {
        if(!e) e = window.event;
        var keyCode = e.code || e.key;
        if(keyCode == "Enter") {
            loadData(searchInput);
        }
    }

</script>

{#if error != ""}
<ErrorDialog close={() => error = ""} {error} />
{/if}

<div class="site">
    <div class="row width-100">
        <h3 class="headline">{loading ? "Teilnehmer werden geladen.." : "Teilnehmer"}</h3>
        <div class="row">
            {#if searchInput != ""}	
            <button on:click={resetSearch} class="button button-secondary">Zurücksetzen</button>
            {/if}
            <input bind:value={searchInput} on:input={search} on:keypress={inputKeyEvent} placeholder="Name" />
            <button on:click={() => loadData(searchInput)} class="button">Suchen</button>
        </div>
    </div>

    {#if !loading}
    <div transition:fade class="years">
        {#if users.length == 0}
        <p class="text-medium">Keine Teilnehmer vorhanden</p>
        {/if}
        {#each users as user}
        <div class="year">
            <div class="row">
                <span class="material-icons icon-accent icon-medium">person</span>
                <p class="text-medium">{user.name}</p>
            </div>
            <div class="row">
                <button class="button button-small button-secondary">Profil</button>
                <button class="button button-small button-secondary" on:click={() => window.open("/panel/" + user.code)}>Auswahlseite</button>
            </div>
        </div>
        {/each}
        {#if users.length != 0}
        <div class="center">
            <p class="text-medium">Benutze die Suche um mehr Benutzer zu finden..</p>
        </div>
        {/if}
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

    .center {
        width: 100%;
        display: flex;
        justify-content: center;
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
