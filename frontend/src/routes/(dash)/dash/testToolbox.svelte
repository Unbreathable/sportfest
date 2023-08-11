<script lang="ts">
	import FormDialog from "$lib/comp/formDialog.svelte";
	import { postRqAuthorized } from "$lib/scripts/requests";

    export let close = () => {};
    let doing = false;

    async function createStandardYears() {
        if(doing) return;
        doing = true;
        await postRqAuthorized("/years/default", "");
        doing = false;
        location.assign("/dash/years");
    }

    async function createStandardGames() {
        if(doing) return;
        doing = true;
        await postRqAuthorized("/games/default", "");
        doing = false;
        location.assign("/dash/games");
    }

    // Simulation
    let minPerYear: Number, maxPerYear: Number;
    async function createUsers() {
        if(doing) return;
        doing = true;
        await postRqAuthorized("/users/simulation", JSON.stringify({
            "minPerYear": minPerYear,
            "maxPerYear": maxPerYear,
        }));
        doing = false;
        location.assign("/dash/users");
    }

</script>

<FormDialog error="" title="Testwerkzeuge" {close}>
    <div class="section">
        <p class="text-medium">Standards erstellen</p>
        <button class="button" on:click={createStandardYears}>{doing ? "Laden.." : "Standard Jahrgänge erstellen"}</button>
        <button class="button" on:click={createStandardGames}>{doing ? "Laden.." : "Standard Spiele erstellen"}</button>
    </div>

    <div class="section">
        <p class="text-medium">Test Teilnehmer erstellen</p>
        <input type="number" bind:value={minPerYear} placeholder="Minimale Anzahl pro Jahrgang">
        <input type="number" bind:value={maxPerYear} placeholder="Maximale Anzahl pro Jahrgang">
        <button class="button" on:click={createUsers}>{doing ? "Laden.." : "Teilnehmer erstellen"}</button>
    </div>

    <div class="section">
        <p class="text-medium">Test Auswahlen treffen</p>
        <p class="text-small">Wenn du den Knopf drückst, wählt ein Programm für jeden Benutzer zufällig zwischen allen Auswählbaren Spielen 3-5 Spiele aus.</p>
        <button class="button">Auswahlen treffen</button>
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