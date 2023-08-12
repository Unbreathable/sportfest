<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    let ranAlready = false;
    let loading = false;
    let working = false;
    let error = "";

    // Results of an attempt
    let result: string[] = [];

    export let close: () => void;

    async function startAlgorithm() {
        if(working) return;
        working = true;

        const res = await postRqAuthorized("/algorithm/start", "");
        if(!res.success) {
            error = res.message;
            working = false;
            return;
        }

        if(!res.algorithm) {
            error = res.message;
            result = res.logs;
            working = false;
            return;
        }

        error = "hallo welt";
        result = res.logs;
        working = false;
    }

</script>

<FormDialog title={loading ? "Wird geladen.." : "Algorithmus"} {error} {close}>
    {#if !loading}
    <div class="section">
        {#if !ranAlready}
            <p class="text-medium">Der Algorithmus sollte <strong>nur von Administratoren</strong> ausgeführt werden, da es sonst zu komischen Fehlern kommen kann!</p>

            {#if result.length == 0}
                <button class="button" on:click={startAlgorithm}>{working ? "Wird durchgeführt.." : "Versuch durchführen"}</button>
            {:else}
                {#each result as line}
                    <p class="text-medium">{line}</p>
                {/each}
                <div class="buttons">
                    <button class="button">{working ? "Wird akzeptiert.." : "Versuch akzeptieren"}</button>
                    <button class="button button-secondary" on:click={startAlgorithm}>{working ? "Wird durchgeführt.." : "Nochmal durchführen"}</button>
                </div>
            {/if}

        {:else}
            <p class="text-medium">Der Algorithmus wurde bereits ausgeführt. Bitte benachrichtige einen Administrator, um nochmal alles zurückzusetzen und den Algorithmus neu lauf zu lassen.</p>
        {/if}
    </div>
    {/if}
</FormDialog>

<style lang="scss">
    @import '$lib/styles/comp.scss';

    .section {
        display: flex;
        flex-direction: column;
        gap: var(--def-spacing);
    }

    .buttons {
        display: flex;
        gap: var(--def-spacing);
    }

</style>