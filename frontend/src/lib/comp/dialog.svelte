<script>

	import { fade, scale } from "svelte/transition";
    export const index = 100;

</script>
<div style="z-index: {index};" transition:fade class="dialog-outer">
    <div transition:scale class="dialog">
        <slot />
    </div>
</div>

<style lang="scss">
    .dialog-outer {
        position: absolute;
        top: 0px;
        left: 0px;
        width: 100vw;
        height: 100vh;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: rgba(0, 0, 0, 0.5);
    }

    .dialog {
        width: 100%;
        max-width: 500px;
        max-height: 80%;
        overflow-y: scroll;
        background-color: var(--background);
        border-radius: 1rem;
        padding: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.5);
    }
</style>