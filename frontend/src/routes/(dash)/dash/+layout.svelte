<script>
	import Sidebar from "./sidebar.svelte";

</script>

<div class="content">
    <Sidebar />
    
    <div class="settings">
        <slot />
    </div>
</div>

<style lang="scss">

    .content {
        display: flex;
        width: 100vw;
        height: 100vh;
    }
    
    .settings {
        padding: var(--section-spacing);
        width: 100%;
    }

</style>