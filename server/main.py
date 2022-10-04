import requests
from bs4 import BeautifulSoup
from flask import Flask, request
from flask_cors import CORS
import os


template_dir = os.path.abspath('../frontend/public/')
app = Flask(__name__, template_folder=template_dir)
CORS(app)



URL = "https://www.uvic.ca/cas/login?service=https%3A%2F%2Fwww.uvic.ca%2Ftools%2Findex.php"


@app.route("/", methods=["POST", "GET"])
def data():
    content = request.json
    with requests.Session() as s:
        response = s.get(URL, allow_redirects=True)
        html = BeautifulSoup(response.text)
        execution = html.find("input", {"name":"execution"})['value']
        response = s.post(URL,data={"username":content["usr"], "password":content["pwd"], "execution": execution, "_eventId": "submit"})
        response = s.get("https://www.uvic.ca/MyCard/account/summary")

        return response.text


app.run()