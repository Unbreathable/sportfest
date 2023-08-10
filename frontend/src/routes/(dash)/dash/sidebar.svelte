<script lang="ts">
	import { page } from "$app/stores";
	import { slide } from "svelte/transition";

    const teamLinks = [
        {
            icon: "construction",
            name: "Algorithmus",
            link: "/dash/algorithm"
        },
        {
            icon: "sports_volleyball",
            name: "Spiele",
            link: "/dash/matches"
        },
        {
            icon: "problem",
            name: "Fehler",
            link: "/dash/problems"
        },
        {
            icon: "calendar_month",
            name: "Zeitplan (Kommt noch)",
            link: "/dash/schedule"
        }
    ]

</script>

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

    <p class="text-small hl">Das Fest erstellen</p>
    {#each teamLinks as link}
        <a href={link.link} class="link">
            <span class="material-icons icon-medium icon-accent">{link.icon}</span>
            <p class="text-medium">{link.name}</p>
        </a>
    {/each}
</div>

<style lang="scss">
    @import '$lib/styles/comp.scss';

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