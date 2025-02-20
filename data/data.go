package data

import "embed"

//go:embed json
var jsonFS embed.FS

func GetJsonFS() embed.FS {
	return jsonFS
}
