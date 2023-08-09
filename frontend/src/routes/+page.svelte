<script lang="ts">
    import { onMount } from "svelte";
    let dialog: HTMLDialogElement;

    onMount(() => {
        dialog.showModal();
    });
</script>

<div class="test">
    <dialog bind:this={dialog} on:close={() => console.log("hi")}>
        <p class="text-medium">HELLO WORLD</p>
    </dialog>
</div>
<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<style lang="scss">
    .test {
        width: 100vw;
        height: 100vh;
    }
</style>