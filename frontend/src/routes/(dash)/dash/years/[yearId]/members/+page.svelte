<script lang="ts">
	import { fade } from "svelte/transition";
	import { onMount } from "svelte";
	import ErrorDialog from "$lib/comp/errorDialog.svelte";
	import { page } from "$app/stores";
	import { currentYear, loadCurrentYear, yearError, yearLoading } from "$lib/scripts/year";
	import YearDeleteDialog from "../../yearDeleteDialog.svelte";
	import { goto } from "$app/navigation";

    onMount(() => loadData())

    let del = false;

    async function loadData() {
        await loadCurrentYear($page.params.yearId, $currentYear.id)
    }

</script>

{#if del}
<YearDeleteDialog year={$currentYear} refresh={() => goto("/dash/years")} close={() => del = false} />
{/if}

{#if $yearError != ""}
<ErrorDialog close={() => yearError.set("")} error={$yearError} />
{/if}

<div class="site">
    <div class="row width-100">
        <h3 class="headline">{$yearLoading ? "Wird geladen.." : $currentYear.name + " - Mitglieder"}</h3>
        <button class="button" on:click={() => del = true}>Löschen</button>
    </div>

    {#if !$yearLoading}
    <div transition:fade class="years">
        <div class="member">
            <div class="row">
                <span class="material-icons icon-accent icon-medium">person</span>
                <p class="text-medium">Deine Mutter</p>
            </div>
            <div class="row">
                <button class="button button-small button-secondary">Account bearbeiten</button>
            </div>
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

    .width-100 {
        width: 100%;
    }

    .row {
        display: flex;
        gap: var(--def-spacing);
        align-items: center;
        justify-content: space-between;
    }

    .member {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

</style>
