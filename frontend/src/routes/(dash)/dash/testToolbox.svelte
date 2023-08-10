<script>
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    export let close = () => {};
    let doing = false;

    async function createStandardYears() {
        doing = true;
        await postRqAuthorized("/years/default", "");
        doing = false;
        location.assign("/dash/years");
    }

    async function createStandardGames() {
        doing = true;
        await postRqAuthorized("/games/default", "");
        doing = false;
        location.assign("/dash/games");
    }

</script>

<FormDialog error="" title="Testwerkzeuge" {close}>
    <div class="section">
        <p class="text-medium">Standards erstellen</p>
        <button class="button" on:click={createStandardYears}>{doing ? "Laden.." : "Standard Jahrgänge erstellen"}</button>
        <button class="button" on:click={createStandardGames}>{doing ? "Laden.." : "Standard Spiele erstellen"}</button>
    </div>

    <div class="section">
        <p class="text-medium">Sportfest Simulation</p>
        <input type="number" placeholder="Minimale Anzahl pro Jahrgang">
        <input type="number" placeholder="Maximale Anzahl pro Jahrgang">
        <input type="number" placeholder="Anzahl an Freundschaften">
        <button class="button">Simulation starten</button>
    </div>
</FormDialog>

<style lang="scss">
    @import "$lib/styles/comp.scss";

    .section {
        display: flex;
        flex-direction: column;
        gap: var(--row-spacing);
    }
</style>