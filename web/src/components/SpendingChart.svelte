<script lang="ts">
    import ApexCharts from 'apexcharts'
    import {onMount} from 'svelte'
    import {oneCardData} from "../ts/store";

    function createSpendingChart(){
        //for some reason array map didnt work :( no clue why.
        
        const spent_per_day = {};
        const days_to_show = 7;
        $oneCardData?.Transactions?.All.forEach(v => {
            if(!spent_per_day[v.Date]) spent_per_day[v.Date] = parseFloat(v.Amount.toFixed(2));
            spent_per_day[v.Date] += parseFloat(v.Amount.toFixed(2))
            if(Object.keys(spent_per_day).length >= days_to_show) return;
            
        })

        const options = {
            chart: {
                type: 'bar',
                toolbar: 'false',
                height: '120%'
            },
            series: [
            {
                name: 'spending',
                data: Object.values(spent_per_day),
            }
            ],
            xaxis: {
                categories: Object.keys(spent_per_day),
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
