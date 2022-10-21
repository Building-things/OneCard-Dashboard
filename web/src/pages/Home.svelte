<script>
    import Ring from "../components/Ring.svelte";
    import {oneCardData, page} from "../ts/store";
    import Transaction from "../components/Transaction.svelte";
    
    
    function logout(){
        oneCardData.set({Balances: null, Transactions: null})
        page.set("login")        
    }
</script>


<main>
    <header>
        <h1 id="title">OCD</h1>
        <button class="button" on:click={logout}>Logout</button>
    </header>
    <section id="balances">
        <Ring value_of_current={$oneCardData.Balances?.StandardMealPlan} value_of_total=1291.75></Ring>
        <div>
            <p><b>Standard: </b>$</p>
            <p>{$oneCardData.Balances?.StandardMealPlan}</p>
        </div>
        <div>
            <p><b>Plus: </b>$</p>
            <p>{$oneCardData.Balances?.PlusMealPlan}</p>
        </div>
        <div>
            <p><b>Flex: </b>$</p>
            <p>{$oneCardData.Balances?.Flex}</p>
        </div>
    </section>

    <section>
        <p>Analytics</p>
    </section>

    <section id="transactions">
        {#if $oneCardData?.Transactions}
            {#each $oneCardData?.Transactions?.Recent.splice(0, 3) as t}
                <Transaction transaction={t}></Transaction>
            {/each}
        {/if}
    </section>
</main>

<style>
    main{
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        min-height: 100vh;
    }
    header{
        width: 100vw;
        height: 5vh;
        display: flex;
        align-items: center;
        justify-content: space-between;
        box-sizing: border-box;
        padding: 5% 5%;
    }   

    section{
        background-color: var(--background-accent-colour);
        margin-left: 8%;
        margin-right: 8%;
        border-radius: 15px;
        box-sizing: border-box;
        padding: 5%;
        margin-top: 5ch;
    }

    #transactions{
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    #balances{
        display: grid;
        height: 30%;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr 1fr 1fr;
        padding: 2% 0;
    }
    #balances > div{
        display: flex;
        align-items: center;
        font-size: 16px;
        grid-column: 2;
    }
    
</style>


