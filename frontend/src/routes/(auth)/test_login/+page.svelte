<script lang="ts">
	import { slide } from "svelte/transition";


    let step = 0;

    function nextStep() {
        step++;
    }

</script>

<div class="center">
    <div class="form">
        <h3 class="font-transition {step == 0 ? "headline" : "text-large"}">1. Email eingeben</h3>
        {#if step == 0}
        <div transition:slide class="comp">
            <p class="text-small">Gebe bitte deine Schul E-Mail Adresse an, damit du dich für das Sportfest eintragen kannst.</p>
            <input type="email" placeholder="example@mail.com">
            <button on:click={nextStep} class="button">Weiter</button>
        </div>
        {/if}
    </div>

    <div class="form">
        <h3 class="font-transition {step == 1 ? "headline" : "text-large"}">2. Email bestätigen</h3>
        {#if step == 1}
        <div transition:slide class="comp">
            <p class="text-small">Gebe bitte deine Schul E-Mail Adresse an, damit du dich für das Sportfest eintragen kannst.</p>
            <input type="email" placeholder="example@mail.com">
            <button on:click={nextStep} class="button">Weiter</button>
        </div>
        {/if}
    </div>

    <div class="form">
        <h3 class="font-transition {step == 2 ? "headline" : "text-large"}">3. Jahrgang auswählen</h3>
        {#if step == 2}
        <div transition:slide class="comp">
            <p class="text-small">Gebe bitte deine Schul E-Mail Adresse an, damit du dich für das Sportfest eintragen kannst.</p>
            <input type="email" placeholder="example@mail.com">
            <button on:click={() => step = 0} class="button">Weiter</button>
        </div>
        {/if}
    </div>
</div>

<style lang="scss">
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

    .form {
        gap: 0;
    }

    .headline {
        color: var(--accent);
    }

    .comp {
        margin-top: var(--section-spacing);
    }

    .font-transition {
        transition: 250ms font-size ease, 250ms color ease;
    }

</style>