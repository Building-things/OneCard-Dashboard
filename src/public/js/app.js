function load_styles(){
    const  head = document.getElementsByTagName('HEAD')[0];
    
    let link = document.createElement('link');
    let link2  = document.createElement('link');

    let script  = document.createElement('script');

    script.src = "https://cdn.jsdelivr.net/npm/apexcharts";
    script.type = "module";

    link.rel = 'stylesheet';
    link.type = 'text/css';
    link.href = './static/css/global.css';

    link2.rel = 'stylesheet';
    link2.type = 'text/css';
    link2.href = './static/css/home.css';

    head.appendChild(link);
    head.appendChild(link2);
    head.appendChild(script);
}

function remove_elements(){
    document.querySelector("div.quick-link-container").remove()
    document.querySelectorAll("a").forEach(node => node.remove())
    document.querySelector("div.column-left > section:not(.transactions)").remove()
    document.querySelector('script[type="text/javascript"]').remove()
}
function add_elements(){
    
    const chartContainer = document.createElement("section")
    chartContainer.innerText = "CHART HERE"
    chartContainer.id = "radialChart"

    var options = {
        series: [70],
        chart: {
        height: 350,
        type: 'radialBar',
      },
      plotOptions: {
        radialBar: {
          hollow: {
            size: '70%',
          }
        },
      },
      labels: ['Cricket'],
    };



    document.querySelector("div.column-left").appendChild(chartContainer)
}

function spending_per_day(){
    const res_amt_str = document.querySelectorAll("section.summary.transactions > ul > li > .transaction-amt")[0].innerHTML
    const res_amt_f = parseFloat(res_amt_str.substring(1).replace(",", ""))
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
    const daily_budget = (res_amt_f / time_between_days).toFixed(2);


    const spending_node = document.createElement("p")
    spending_node.innerHTML = `You can spend $${daily_budget} per day.`
    document.querySelector("div.column-right > section").prepend(spending_node)
}

load_styles()
remove_elements()
add_elements()
spending_per_day()