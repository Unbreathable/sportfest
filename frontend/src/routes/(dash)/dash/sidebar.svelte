<script lang="ts">
	import { page } from "$app/stores";


    const platformLinks = [
        {
            icon: "group",
            name: "Teilnehmer",
            link: "/dash/users"
        },
        {
            icon: "school",
            name: "Jahrgänge",
            link: "/dash/years"
        },
        {
            icon: "joystick",
            name: "Auswählbare Spiele",
            link: "/dash/games"
        }
    ]

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
    {#each platformLinks as link}
        <a href={link.link} class="link {$page.url.pathname.includes(link.link) ? "link-selected" : ""}">
            <span class="material-icons icon-medium icon-accent">{link.icon}</span>
            <p class="text-medium">{link.name}</p>
        </a>
    {/each}

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
        width: 400px;
        max-width: 35%;
        height: 100vh;
        display: flex;
        flex-direction: column;
        gap: calc(var(--row-spacing)/2);
    }

    .hl {
        font-weight: bold;
        margin-top: var(--def-spacing);
    }

    .link {
        text-decoration: none;
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

</style>