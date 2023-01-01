let value_of_current = document.querySelector("#standard").innerHTML.split("$")[1]
let value_of_total = document.querySelector("input[type='hidden']").getAttribute("value")
value_of_total = parseFloat(value_of_total).toFixed(2)
value_of_current = parseFloat(value_of_current).toFixed(2)

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
    height: (document.querySelector("#balances").getBoundingClientRect().width) / 1.5,
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
      colors: 'var(--accent-color)'
  }
};

let chart = new ApexCharts(document.querySelector("#chart"), options);
  
chart.render();