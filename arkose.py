from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
import time
from flask import Flask, jsonify, render_template

app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/token")
def get_token():
    chrome_options = Options()
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-gpu")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--disable-blink-features=AutomationControlled")
    chrome_options.add_argument("--incognito")
    chrome_options.add_argument("--headless")

    driver = webdriver.Chrome(options=chrome_options)
    driver.get("http://127.0.0.1:8080")

    init_js = """
            window.token = '';
            window.console.log = function(message) {
                token = message;
            };
        """
    driver.execute_script(init_js)

    element = driver.find_element(By.ID, "enforcement-trigger")
    element.click()
    time.sleep(1)

    token = driver.execute_script("return token;")
    return jsonify({"token": token})


if __name__ == "__main__":
    app.run()
