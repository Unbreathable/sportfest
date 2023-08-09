<script lang="ts">
	import { goto } from "$app/navigation";
	import { postRq } from "$lib/config/requests";
	import { fade, scale } from "svelte/transition";

    let showPassword = false;
    let username: string;
    let password: string;

    let loading = false;

    let error = false;
    let errorMessage = "";

    async function login() {
        if(loading || error) return;
        loading = true;

        console.log("login");

        const res = await postRq("/account/login", JSON.stringify({
            "name": username,
            "password": password
        }));

        if(!res.success) {
            loading = false;
            error = true;
            errorMessage = res.message ?? "Ein unterwarteter Fehler ist aufgetreten. Bitte versuche es später erneut.";
        } else {
            localStorage.setItem("token", res.token);
            goto("/dash")
        }
    }

    function close() {
        error = false;
    }

    function handlePassword(e: any) {
        password = e.target.value;
    }

</script>

<svelte:head>
    <title>Als Admin einloggen</title>
</svelte:head>

{#if error}
<div transition:fade class="dialog-outer">
    <div transition:scale class="dialog">
        <p class="text-medium">{errorMessage}</p>
        <button on:click={close} class="button">Verstanden</button>
    </div>
</div>
{/if}

<div class="center">
    <div class="form">
        <h3 class="headline">Einloggen</h3>
        <div class="comp">
            <p class="text-small">Falls du ein Admin-Konto hast, kannst du dich hier damit einloggen.</p>
            <input bind:value={username} type="username" placeholder="Benutzername">
            <input on:input={handlePassword} type="{!showPassword ? "password" : "text"}" placeholder="Passwort">
            <div class="row">
                <button class="button" on:click={login}>{loading ? "Einloggen.." : "Login"}</button>
                <button class="button button-secondary {showPassword ? "button-selected" : ""}" on:click={() => showPassword = !showPassword}>Passwort anzeigen</button>
            </div>
        </div>
    </div>

</div>

<style lang="scss">
    @import '$lib/styles/dialog.scss';
    @import '$lib/styles/comp.scss';
    @import '$lib/styles/form.scss';

    .center {
        width: 100vw;
		height: 100vh;
        display: flex;
        gap: var(--section-spacing);
        align-items: center;
        justify-content: center;
        flex-direction: column;
    }

    .row {
        display: flex;
        align-items: center;
        justify-content: space-between;
    }

    .form {
        gap: 0;
    }

    .comp {
        margin-top: var(--section-spacing);
    }

    .font-transition {
        transition: 250ms font-size ease, 250ms color ease;
    }

</style>