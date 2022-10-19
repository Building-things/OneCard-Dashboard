<script>
    import Ring from "../components/Ring.svelte";
    import {oneCardData, page} from "../ts/store";
    import RecentTransactions from "../components/RecentTransactions.svelte";
    
    
    function logout(){
        page.set("login")        
    }
</script>

<main>
    <h1 id="title">Home Page</h1>

    <input type="button" class="button logout-button" value="Logout" on:click={logout}>
    <div class="ring">
        <Ring value_of_current={$oneCardData.Balances?.StandardMealPlan} value_of_total=1291.75></Ring> 
    </div>

    <div class='analytics'>
        <p>Space for graph / data showing reccomended spending per day</p>
    </div>

    <div class="recent-transactions">
        {#each Array(3) as _, i}
            <RecentTransactions index={i} ></RecentTransactions>
        {/each}
    </div>
</main>

<style>
    #title {
        display: flex;
        justify-content: center;
        padding: 25px;
    }

    .ring {
        display: flex;
        justify-content: center;
        background-color: var(--background-accent-colour);
        margin-left: 10%;
        margin-right: 10%;
        border-radius: 15px;
        height: fit-content;
    }

    main {
        height: 100vh;
    }

    .logout-button {
        position: absolute;
        margin: 25px;
        top: 0;
    }

    .recent-transactions {
        display: flex;
        justify-content: center;
    }

    .analytics {
        display: flex;
        justify-content: center;
        background-color: var(--background-accent-colour);
        width: 80%;
        height: 200px;
        border-radius: 15px;
        margin: 10%;
    }
        
    
</style>


