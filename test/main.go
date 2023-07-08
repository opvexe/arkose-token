/*
Copyright 2022 The deepauto-io LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var chromeArgs = []string{
	"--no-sandbox",
	"--disable-gpu",
	"--disable-dev-shm-usage",
	"--disable-blink-features=AutomationControlled",
	"--incognito",
	"--headless=new",
}

func main() {
	webDriver, err := selenium.NewRemote(selenium.Capabilities{
		"chromeOptions": chrome.Capabilities{
			Args:            chromeArgs,
			ExcludeSwitches: []string{"enable-automation"},
		},
	}, "http://127.0.0.1:9515")

	if err != nil {
		panic(err)
	}

	if err = webDriver.Get("http://tms.im/f.html"); err != nil {
		panic(err)
	}

	initJS := `
		window.token = '';
		window.console.log = (message) => {
			token = message;
		};
	`
	if _, err = webDriver.ExecuteScript(initJS, nil); err != nil {
		panic(err)
	}

	element, _ := webDriver.FindElement(selenium.ByID, "enforcement-trigger")
	element.Click()

	time.Sleep(1 * time.Second)
	token, _ := webDriver.ExecuteScript("return token;", nil)

	fmt.Println(token)
}
