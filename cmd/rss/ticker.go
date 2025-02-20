package main

import (
	"time"
)

func (a *AppConfig) UpdateTimeStamp() {

	for {
		time.Sleep(time.Minute * 5)
		a.CurrentDate = time.Now().Format(time.RFC1123Z)
		a.CurrentTimestamp = time.Now().UnixNano()
	}

}
