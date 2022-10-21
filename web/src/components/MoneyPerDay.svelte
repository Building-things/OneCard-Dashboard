<script>

    import {onMount} from 'svelte'
    import ApexCharts from 'apexcharts';


    function SpendingChart(){
        let options = {
            chart: {
            height: 200,
            width: 150,
            type: "line",
            foreColor: 'var(--background-accent-colour)'
        },
        fill: {
            colors: 'var(--accent-colour)'
        },
        series: [
            {
            name: "Temperature",
            data: [15, 20, 25, 29, 33, 36, 37, 32, 27, 23]
            }
        ],
        stroke: {
            curve: 'smooth'
        },
        yaxis: {
            min: 0,
            max: 45
        },
        }

        var chart = new ApexCharts(document.querySelector("#transactions"), options);
        
        chart.render();
    }
    
    function MoneyPerDay(){
        
    }

    onMount(SpendingChart);
</script>

<main>
    <div id="transactions">

    </div>

</main>

<style>
    #transactions {
        background-color: var(--background-accent-colour);
        width: 200px;
        display: flex;
        justify-content: center;
        text-align: center;
    }
</style>