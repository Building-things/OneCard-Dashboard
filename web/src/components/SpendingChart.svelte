<script lang="ts">
    import ApexCharts from 'apexcharts'
    import {onMount} from 'svelte'
    import {oneCardData} from "../ts/store";

    function createSpendingChart(){
        //for some reason array map didnt work :( no clue why.
        const amounts:Array<number> = [];
        $oneCardData?.Transactions?.Recent.forEach(v => {
            amounts.push(v.Amount)
            
        })
        const xaxis:Array<string> = [];
        $oneCardData?.Transactions?.Recent.forEach(v => {
            xaxis.push(v.Date)
            
        })
        //console.log(amounts, xaxis);
        
        const options = {
            chart: {
                type: 'bar'
            },
            series: [
            {
                name: 'spending',
                data: amounts,
            }
            ],
            xaxis: {
                categories: xaxis,
                labels:{
                    style: {
                        colors: "white",
                        fontSize: "14px",
                    },
                }
            },
            yaxis: {
                labels:{
                    style: {
                        colors: "white",
                        fontSize: "14px",
                    },
                },
            },
            fill: {
                colors: 'var(--accent-colour)'
            }
        };

        let chart = new ApexCharts(document.querySelector("#chart2"), options);
        
        chart.render();
    }
    onMount(createSpendingChart);
</script>

<div id="chart2">

</div>
