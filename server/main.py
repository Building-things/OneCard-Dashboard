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
        content = request.json
        if content.get("usr", None) == None or content.get("usr", None) == None:
            return jsonify(data), 400
        # retrieve data

        response = s.get(URL, allow_redirects=True)
        if response.status_code == 400:
            return jsonify(data), 400

        login_html = BeautifulSoup(response.text, features="html.parser")
        execution = login_html.find("input", {"name":"execution"})['value']

        response = s.post(URL,data={"username":content["usr"], "password":content["pwd"], "execution": execution, "_eventId": "submit"})
        if response.status_code == 400:
            return jsonify(data), 400

        
        response = s.get("https://www.uvic.ca/MyCard/account/summary")
        if response.status_code == 400:
            return jsonify(data), 400
        
        onecard_html = BeautifulSoup(response.text, features="html.parser")

        balances = {}
        balance_names = onecard_html.find_all("span", attrs={"class": "transaction-desc"})
        balance_values = onecard_html.find_all("span", attrs={"class": "transaction-amt"})
        
        for i, balance_name in enumerate(balance_names):
            if balance_name.text not in {"Market", "Servery"}:
                balances[balance_name.text] = balance_values[i].text
                data["balances"] = balances

        return jsonify(data)


app.run()