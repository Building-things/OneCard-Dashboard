<script>    
    import { oneCardData } from "../ts/store";

    export let index;
</script>

<main>
    <div class="rct-tr">
        <p>{$oneCardData.Transactions?.Recent[index]?.Date}</p>
        <p>{$oneCardData.Transactions?.Recent[index]?.Location}</p>
        <p>${$oneCardData.Transactions?.Recent[index]?.Amount}</p>
    </div>
</main>

<style>
    .rct-tr {
        background-color: var(--background-accent-colour);
        margin: 15px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        border-radius: 10px;
        padding: 15px;
        text-align: center;
    }
</style>