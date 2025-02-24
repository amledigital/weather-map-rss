package main

import "net/http"

func redirectToFeedExt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/rss/v1/weather-maps" {

			pretty := r.URL.Query().Get("pretty")

			if pretty == "true" {

				http.Redirect(w, r, "/rss/v1/weather-maps/feed.rss?pretty=true", http.StatusSeeOther)

			} else {
				http.Redirect(w, r, "/rss/v1/weather-maps/feed.rss", http.StatusSeeOther)
			}

			return
		}

		next.ServeHTTP(w, r)

	})
}
