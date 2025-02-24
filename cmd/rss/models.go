package main

import (
	"fmt"
	"time"
)

type AppConfig struct {
	Port             string
	BaseURL          string
	CurrentDate      string
	CurrentTimestamp int64
	DoneChan         chan bool
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		CurrentDate:      time.Now().Format(time.RFC1123Z),
		CurrentTimestamp: time.Now().UnixNano(),
		DoneChan:         make(chan bool),
	}
}

type rss struct {
	XmlnsAtom  string   `xml:"xmlns:atom,attr"`
	XmlnsMedia string   `xml:"xmlns:media,attr"`
	Version    string   `xml:"version,attr"`
	Channel    *Channel `xml:"channel"`
}

func NewRssFeed() *rss {
	var feed = &rss{
		XmlnsAtom:  "http://www.w3.org/2005/Atom",
		XmlnsMedia: "http://search.yahoo.com/mrss/",
		Version:    "2.0",
		Channel:    NewChannel("", "", ""),
	}
	return feed
}

type Channel struct {
	AtomLink      *AtomLink         `xml:"atom:link"`
	Title         string            `xml:"title"`
	Link          string            `xml:"link"`
	Description   string            `xml:"description"`
	Copyright     string            `xml:"copyright"`
	Language      string            `xml:"language"`
	PubDate       string            `xml:"pubDate"`
	LastBuildDate string            `xml:"lastBuildDate"`
	Generator     string            `xml:"generator"`
	Items         []*WeatherMapItem `xml:"item"`
}

func NewChannel(title, link, description string) *Channel {

	if title == "" {
		title = "910 Media Group - Weather Maps"
	}

	if link == "" {
		link = "https://www.9and10news.com/weather/"
	}

	if description == "" {
		description = "Northern Michigan Local Weather Maps Updated every 5 minutes"
	}

	return &Channel{
		AtomLink:      NewAtomLink(),
		Title:         title,
		Link:          link,
		Description:   description,
		Copyright:     fmt.Sprintf("910 Media Group %d", time.Now().Year()),
		Language:      "en-US",
		PubDate:       time.Now().Format(time.RFC1123Z),
		LastBuildDate: time.Now().Format(time.RFC1123Z),
		Generator:     "https://www.9and10news.com",
		Items:         NewWeatherMapItemList(),
	}
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

func NewAtomLink() *AtomLink {
	return &AtomLink{
		Rel:  "self",
		Type: "application/rss+xml",
		Href: "",
	}
}

type WeatherMapItem struct {
	Title          string         `xml:"title"`
	Link           string         `xml:"link"`
	Description    string         `xml:"description"`
	PubDate        string         `xml:"pubDate"`
	Guid           string         `xml:"guid"`
	MediaThumbnail MediaThumbnail `xml:"media:thumbnail"`
	MediaContent   MediaContent   `xml:"media:content"`
}

func NewWeatherMapItem(title, link, description, pubdate, guid string) *WeatherMapItem {
	return &WeatherMapItem{
		Title:       title,
		Link:        link,
		Description: description,
		PubDate:     pubdate,
		Guid:        guid,
	}
}

type MediaThumbnail struct {
	Height string `xml:"height,attr"`
	Width  string `xml:"width,attr"`
	URL    string `xml:"url,attr"`
}

type MediaContent struct {
	Medium string `xml:"medium,attr"`
	URL    string `xml:"url,attr"`
}

func NewWeatherMapItemList() []*WeatherMapItem {

	return []*WeatherMapItem{}

}

type WeatherMapJson struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Link    string `json:"url"`
	PubDate string `json:"pub_date,omitempty"`
	Guid    string `json:"guid,omitempty"`
}

type WeatherMapJsonList struct {
	WeatherMaps []*WeatherMapItem `json:"weather_maps"`
}
