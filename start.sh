#!/bin/sh

chromedriver --port=9515 --url-base=/wd/hub &
./arkosed
