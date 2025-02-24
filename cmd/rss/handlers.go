package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/amledigital/weather-map-rss/data"
)

var (
	feed          *rss
	weathMapItems []*WeatherMapItem
	weathMapJson  WeatherMapJsonList
)

func (a *AppConfig) HandleGetWeatherMapRSS(w http.ResponseWriter, r *http.Request) {

	prettyPrint := r.URL.Query().Get("pretty")

	w.Header().Set("Content-Type", "application/rss+xml")
	jsonFS := data.GetJsonFS()

	mapF, err := jsonFS.ReadFile("json/maps.json")

	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(mapF, &weathMapJson)

	if err != nil {
		log.Panic(err)
	}

	for i := range weathMapJson.WeatherMaps {

		var curr = weathMapJson.WeatherMaps[i]

		// get media link from json
		mediaLink := fmt.Sprintf("%s?timestamp=%d", curr.Link, a.CurrentTimestamp)
		// reset Link to weather page for RSS feed
		curr.Link = "https://www.9and10news.com/weather/"
		thumb := MediaThumbnail{}
		thumb.Height = "405"
		thumb.Width = "720"
		thumb.URL = mediaLink

		curr.MediaThumbnail = thumb

		curr.MediaContent = MediaContent{
			Medium: "image",
			URL:    mediaLink,
		}

		curr.PubDate = a.CurrentDate
		curr.Guid = Guid{
			Text:        mediaLink,
			IsPermalink: false,
		}
	}

	feed = NewRssFeed()

	feed.Channel.AtomLink.Href = app.BaseURL + r.URL.Path

	feed.Channel.Items = weathMapJson.WeatherMaps

	if prettyPrint == "true" {

		b, err := xml.MarshalIndent(feed, "", "\t")

		if err != nil {
			log.Fatalln(err)
		}

		// hack and cleanup formatting for Baron RSS feed ingress

		out := strings.ReplaceAll(string(b), "></media:thumbnail>", " />")
		out = strings.ReplaceAll(out, "></media:content>", " />")
		out = strings.ReplaceAll(out, "></atom:link>", " />")

		_, err = io.WriteString(w, out)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {

		err = xml.NewEncoder(w).Encode(feed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
