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
	"log"
	"net/http"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	"github.com/gin-gonic/gin"
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
	h := gin.Default()
	h.Any("/health", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusNoContent)
	})

	h.GET("/token", func(ctx *gin.Context) {
		webDriver, err := selenium.NewRemote(selenium.Capabilities{
			"chromeOptions": chrome.Capabilities{
				Args:            chromeArgs,
				ExcludeSwitches: []string{"enable-automation"},
			},
		}, "http://127.0.0.1:9515")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		if err := webDriver.Get("https://tms.im/f.html"); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		initJS := `
		window.token = '';
		window.console.log = (message) => {
			token = message;
		};
	`
		_, err = webDriver.ExecuteScript(initJS, nil)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}

		element, err := webDriver.FindElement(selenium.ByID, "enforcement-trigger")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		element.Click()

		time.Sleep(1 * time.Second)
		token, _ := webDriver.ExecuteScript("return token;", nil)

		ctx.JSON(http.StatusOK, struct {
			Token interface{} `json:"token"`
		}{
			Token: token,
		})
	})

	if err := h.Run(); err != nil {
		log.Panic(fmt.Errorf("http server run err: %s", err))
	}
}
