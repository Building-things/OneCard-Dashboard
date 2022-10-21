<script lang="ts">    
    import type {Transaction} from "../ts/store"

    export let transaction:Transaction;
</script>

<div>
    <p id="date">{transaction.Date}</p>
    <p id="amount">{transaction.Amount}</p>
    <p id="location">{transaction.Location}</p>
</div>

<style>
    div{
        width: 100%;
        box-sizing: border-box;
        padding: 5% 5%;
        background-color: var(--background-main-colour);
        margin-bottom: 1vh;
        border-left: 4px solid var(--accent-colour)
    }
    #date{
        font-weight: bold;
        
    }
</style>