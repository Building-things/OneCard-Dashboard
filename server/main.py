import requests
from bs4 import BeautifulSoup
from flask import Flask, request, jsonify
from flask_cors import CORS


app = Flask(__name__)
CORS(app)



URL = "https://www.uvic.ca/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php"


@app.route("/", methods=["POST", "GET"])
def data():
    data = {}

    if request.method != "POST" and request.method != "OPTIONS":
        return jsonify(data), 400

    with requests.Session() as s:
        content = request.form
        if content.get("username", None) == None or content.get("password", None) == None:
            return jsonify(data), 400
        # retrieve data

        response = s.get(URL, allow_redirects=True)
        if not response.ok:
            return jsonify(data), 400

        login_html = BeautifulSoup(response.text, features="html.parser")
        execution = login_html.find("input", {"name":"execution"})['value']

        response = s.post(URL,data={"username":content["username"], "password":content["password"], "execution": execution, "_eventId": "submit"})
        if not response.ok:
            return jsonify(data), 400

        
        response = s.get("https://www.uvic.ca/MyCard/account/summary")
        if not response.ok:
            return jsonify(data), 400
       
        onecard_html = BeautifulSoup(response.text, features="html.parser")
        try:
            execution = onecard_html.find("input", {"name":"execution"})['value']
            if execution:
                return jsonify(data), 400
        except TypeError:
            pass

        balances = {}
        transactions = {}
        balance_names = onecard_html.find_all("span", attrs={"class": "transaction-desc"})
        balance_values = onecard_html.find_all("span", attrs={"class": "transaction-amt"})
        for i, balance_name in enumerate(balance_names):
                key = ""
                if balance_name.text == "Residence Meal Plan":
                    key = "standardMealPlan"
                elif balance_name.text == "ONECard Flex":
                    key = "flex"
                elif balance_name.text == "Plus Meal Plan":
                    key = "plusMealPlan"
                else:
                    key = balance_name.text
                    transactions[key] = float(balance_values[i].text[1:])
                    continue
                #the text[1:] removes the dollar sign
                balances[key] = float(balance_values[i].text[1:])
                
        data["balances"] = balances    
        data["transactions"] = {"recent": transactions}

        return jsonify(data)


app.run()