<script>
    import Ring from "../components/ring.svelte";

    export let user_data;
</script>

{#if user_data.balances}
<main>
    <header>
        <h1>OCD</h1>
    </header>
    <section class="container">
        <!--TODO get 1292 from API not hard code it-->
        <Ring size=120 value_of_total=1292 value_of_current={user_data.balances["Residence Meal Plan"]}/>
        <p>{user_data.balances["Residence Meal Plan"]}</p>
    </section>
</main>
{:else}
    <h1>TODO LOADING</h1>
{/if}

<style>
    main{
        width: 100vw;
        color: white;
    }
    h1{
        margin: 0;
    }
    .container{
        width: 90%;
        margin: auto;
        background-color: var(--background-accent-colour);
        border-radius: 10px;
    }
</style>