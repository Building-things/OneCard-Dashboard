<script lang="ts">
    import {page} from "./ts/store";

    import Login from "./pages/Login.svelte";
    import Error from "./pages/Error.svelte";
    import Home from "./pages/Home.svelte"

    const pages = {
        "login": Login,
        "home": Home,
    }
</script>

{#if $page in pages}
    <svelte:component this={pages[$page]}></svelte:component>
{:else}
    <Error/>
{/if}