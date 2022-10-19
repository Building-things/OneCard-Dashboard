<script>    
    import { oneCardData } from "../ts/store";

    export let index;
</script>

<main>
    <div>
        <p>{$oneCardData.Transactions?.Recent[index]?.Date}</p>
        <p>{$oneCardData.Transactions?.Recent[index]?.Location}</p>
        <p>${$oneCardData.Transactions?.Recent[index]?.Amount}</p>
    </div>
</main>

<style>
    
</style>