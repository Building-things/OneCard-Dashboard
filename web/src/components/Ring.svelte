<script>
    export let value_of_total;
    export let value_of_current;

    import ApexCharts from 'apexcharts'
    import {onMount} from 'svelte'



    function CreateRing(){
        const options = {
          series: [value_of_current / value_of_total * 100 ],  //this has to be a percent
          chart: {
            foreColor: '#FFFFFF',
            animations: {
            enabled: true,
            easing: 'easeinout',
            speed: 800,
            animateGradually: {
                enabled: true,
                delay: 150
            },
            dynamicAnimation: {
                enabled: true,
                speed: 350
            }
        },
          height: ((window.innerHeight) / 4),
          type: 'radialBar',
          
        },
        plotOptions: {
          radialBar: {
            dataLabels: {
                name: {
                    fontSize: '30px',
                },
                    value: {
                    fontSize: '16px',
                },
                total: {
                    show: true,
                    label: 'Balance: ',
                    
                    formatter: function () { return `$${value_of_current}`}
                },
                            
            },
            hollow: {
              size: '60%',
            }
          }
        },
        fill: {
            colors: 'var(--accent-colour)'
        }
        };

        let chart = new ApexCharts(document.querySelector("#chart"), options);
        
        chart.render();
    }


    onMount(CreateRing);
</script>


<div id="chart">
</div>

<style>
    #chart{
        display: block;
        grid-row: 1 / 4;
        grid-column: 1;
    }
</style>