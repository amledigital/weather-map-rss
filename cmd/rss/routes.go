package main

import (
	"net/http"
)

func routes() *http.ServeMux {

	rtr := http.NewServeMux()

	rtr.HandleFunc("/rss/v1/weather-maps/rss.xml", app.HandleGetWeatherMapRSS)

	return rtr

}
