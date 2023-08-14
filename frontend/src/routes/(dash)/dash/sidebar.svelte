<script lang="ts">
	import { page } from "$app/stores";
	import { slide } from "svelte/transition";
	import TestToolbox from "./testToolbox.svelte";

    let testToolkit = false;

</script>

{#if testToolkit}
<TestToolbox close={() => testToolkit = false} />
{/if}

<div class="content shadow">
    <h3 class="text-large">Sportfest: Manager</h3>

    <p class="text-small hl">Plattform einrichten</p>
    <a href="/dash/users" class="link {$page.url.pathname == "/dash/users" ? "link-selected" : ""}">
        <span class="material-icons icon-medium icon-accent">group</span>
        <p class="text-medium">Teilnehmer</p>
    </a>

    <a href="/dash/years" class="link {$page.url.pathname.includes("/dash/years") ? "link-selected" : ""}">
        <span class="material-icons icon-medium icon-accent">school</span>
        <p class="text-medium">Jahrgänge</p>
    </a>
    {#if $page.url.pathname != "/dash/years" && $page.url.pathname.startsWith("/dash/years")}
        <div transition:slide class="link-below">
            <a href="/dash/years/{$page.params.yearId}/members" class="link {$page.url.pathname.endsWith("/members") ? "link-secondary" : ""}">
                <span class="material-icons icon-medium icon-accent">group</span>
                <p class="text-medium">Mitglieder</p>
            </a>

            <a href="/dash/years/{$page.params.yearId}/matches" class="link {$page.url.pathname.endsWith("/matches") ? "link-secondary" : ""}">
                <span class="material-icons icon-medium icon-accent">sports_volleyball</span>
                <p class="text-medium">Spiele</p>
            </a>
        </div>
    {/if}

    <a href="/dash/games" class="link {$page.url.pathname == "/dash/games" ? "link-selected" : ""}">
        <span class="material-icons icon-medium icon-accent">joystick</span>
        <p class="text-medium">Auswählbare Spiele</p>
    </a>

    <p class="text-small hl">Teams und Zeitplan</p>
    <a href="/dash/teams" class="link {$page.url.pathname == "/dash/teams" ? "link-selected" : ""}">
        <span class="material-icons icon-medium icon-accent">hardware</span>
        <p class="text-medium">Teameinteilung</p>
    </a>

    <a href="/dash/matches" class="link {$page.url.pathname.startsWith("/dash/matches") ? "link-selected" : ""}">
        <span class="material-icons icon-medium icon-accent">sports_volleyball</span>
        <p class="text-medium">Spiele</p>
    </a>

    <div class="button-row">
        <button class="button button-small button-secondary">Neues Adminkonto</button>
        <button class="button button-small button-secondary" on:click={() => testToolkit = true}>Testwerkzeuge</button>
    </div>
</div>

<style lang="scss">
    @import '$lib/styles/comp.scss';

    .button-row {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        margin-top: auto;
    }

    .content {
        padding: var(--def-spacing);
        min-width: 400px;
        height: calc(100vh - 2 * var(--def-spacing));
        display: flex;
        flex-direction: column;
    }

    .hl {
        font-weight: bold;
        margin-top: var(--def-spacing);
        margin-bottom: calc(var(--def-spacing)/2);
    }

    .hl-small {
        margin-bottom: calc(var(--def-spacing)/2);
    }

    .link {
        text-decoration: none;
        margin-bottom: calc(var(--def-spacing)/2);
        color: var(--text-color);
        display: flex;
        align-items: center;
        gap: var(--row-spacing);
        padding: var(--row-spacing);
        border-radius: var(--def-spacing);
        transition: 250ms all ease;
    
        &:hover {
            background-color: var(--secondary);
        }
    }

    .link-secondary {
        background-color: var(--secondary);
        color: var(--text-color);

        .material-icons {
            color: var(--accent);
        }

        &:hover {
            background-color: var(--secondary);
        }
    }

    .link-selected {
        background-color: var(--primary);
        color: var(--button-text);

        .material-icons {
            color: var(--secondary);
        }

        &:hover {
            background-color: var(--primary);
        }
    }

    .link-below {
        margin: 0 0 var(--def-spacing) var(--def-spacing);
    }

</style>