<script>

    import Ring from "c:/Users/leona/Github repos/OneCard-Dashboard/web/src/components/Ring.svelte";
    import {oneCardData, page} from "../ts/store";
    
    function logout(){
        page.set("login")        
    }
</script>

<main>
    <h1 id="title">Home Page</h1>

    <input type="button" class="button logout-button" value="Logout" on:click={logout}>
    <div class="ring">
        <Ring value_of_current={$oneCardData.balances.standardMealPlan} value_of_total=1291.75></Ring>
    </div>
</main>

<style>
    #title {
        display: flex;
        justify-content: center;
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
        display: fixed;
    }
        
    
</style>


