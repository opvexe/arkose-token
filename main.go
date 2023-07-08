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

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"

	"github.com/gin-gonic/gin"
)

var (
	proxy = selenium.Proxy{
		Type: selenium.Manual,
		HTTP: "http://hahaha:xixixi@202.182.71.232:17575", //   "http://username:password@proxy-host:proxy-port",
	}
)

func main() {
	h := gin.Default()
	chromeArgs := []string{
		"--no-sandbox",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--disable-blink-features=AutomationControlled",
		"--incognito",
		"--headless=new",
		//"--proxy-server=" + proxy.HTTP,
	}

	webDriver, err := selenium.NewRemote(selenium.Capabilities{
		"chromeOptions": chrome.Capabilities{
			Args:            chromeArgs,
			ExcludeSwitches: []string{"enable-automation"},
		},
	}, "http://127.0.0.1:9515")

	if err != nil {
		log.Panic("selenium: new remote web driver err: ", err)
	}

	h.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	h.Any("/health", func(ctx *gin.Context) {
		ctx.Writer.WriteHeader(http.StatusNoContent)
	})

	h.GET("/token", func(ctx *gin.Context) {
		if err := webDriver.Get("http://127.0.0.0:8080/"); err != nil {
			log.Printf("selenium: get web err: %s", err)
			ctx.JSON(http.StatusInternalServerError, err)
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
			log.Printf("selenium: execute script err: %s", err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		element, err := webDriver.FindElement(selenium.ByID, "enforcement-trigger")
		if err != nil {
			log.Printf("selenium: find element err: %s", err)
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		element.Click()

		token, _ := webDriver.ExecuteScript("return token;", nil)

		if token == "" {
			log.Printf("selenium: token is empty")
			ctx.JSON(http.StatusInternalServerError, fmt.Errorf("IP 被封"))
			return
		}

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
