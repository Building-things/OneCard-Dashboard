<script lang="ts">
    import Ring from "../components/Ring.svelte";
    import {oneCardData, page} from "../ts/store";
    import Transaction from "../components/Transaction.svelte";
    import { onMount } from "svelte";
    import SpendingChart from "../components/SpendingChart.svelte"

    let showMoreTransactions = false;
    
    let daily_budget = 0;
    onMount(()=>{
        const current_date = new Date()
        const current_month = new Date().getMonth()
        const current_year = new Date().getFullYear()
        let term_end_date;
        if(current_month > 6){
            term_end_date = new Date(`December 20, ${current_year} 23:59:000`)
        }else if(current_month < 5){
            term_end_date = new Date(`April 29, ${current_year} 23:59:000`)
        }else{
            alert("Summer Term Not Supported");
        }
        const time_between = term_end_date.getTime() - current_date.getTime();
        const time_between_days = Math.floor(time_between / (1000 * 60 * 60 * 24));
        daily_budget = parseFloat(($oneCardData?.Balances?.StandardMealPlan / time_between_days).toFixed(2));
        
    })
    function logout(){
        oneCardData.set({Balances: null, Transactions: null, Meta:null})
        page.set("login")        
    }
</script>


<main>
    <header>
        <h1>OCD</h1>
        <button class="button" on:click={logout}>Logout</button>
    </header>
    <section id="balances">
        <Ring value_of_current={$oneCardData.Balances?.StandardMealPlan} value_of_total={$oneCardData.Meta?.StandardTotal}></Ring>
        <div>
            <p><b>Standard: </b></p>
            <b class="dollar-amount">${$oneCardData.Balances?.StandardMealPlan}</b>
        </div>
        <div>
            <p><b>Plus: </b></p>
            <b class="dollar-amount">${$oneCardData.Balances?.PlusMealPlan}</b>
        </div>
        <div>
            <p><b>Flex: </b></p>
            <b class="dollar-amount">${$oneCardData.Balances?.Flex}</b>
        </div>
    </section>

    <section id="analytics">
        <b>You Can Spend <b class="dollar-amount">${daily_budget}</b>/day</b>
        <SpendingChart></SpendingChart>
    </section>

    <section id="transactions">
        {#if $oneCardData?.Transactions}
            {#each $oneCardData?.Transactions?.Recent as t}
                <Transaction transaction={t}></Transaction>
            {/each}
            {#if showMoreTransactions}
                {#each $oneCardData?.Transactions?.All as t}
                    <Transaction transaction={t}></Transaction>
                {/each}
            {/if}
        {/if}
        <button class="button" style="width: 100%" on:click={()=>{showMoreTransactions = true}}>Load More</button>

    </section>
</main>

<style>
    main{
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        height: fit-content;
    }
    header{
        width: 100vw;
        height: 8vh;
        display: flex;
        align-items: center;
        justify-content: space-between;
        box-sizing: border-box;
        padding: 5% 5%;
        margin-bottom: 2vh;
    }   

    section{
        background-color: var(--background-accent-colour);
        margin-left: 8%;
        margin-right: 8%;
        border-radius: 15px;
        box-sizing: border-box;
        padding: 5%;
        margin-bottom: 4vh;
    }

    #transactions{
        display: flex;
        flex-direction: column;
        justify-content: flex-start;
        align-items: center;
        overflow-y: scroll;
        max-height: 100%;
    }

    #balances{
        display: grid;
        height: 30%;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr 1fr 1fr;
        padding: 0 3%;
    }
    #balances > div{
        display: flex;
        align-items: center;
        justify-content: flex-end;
        font-size: 16px;
        grid-column: 2;
    }
    #analytics{
        display: flex;
        flex-direction: column;
        align-items: center;
    }
</style>


