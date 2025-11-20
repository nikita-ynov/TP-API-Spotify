package main

import (
	"fmt"
	"net/http"
	initTemp "spotify/pages"
	"spotify/router"
)

func main() {
	initTemp.Init()

	r := router.New()

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	blue := "\033[34m"
	reset := "\033[0m"

	fmt.Println("Serveur démarré sur", blue+"http://localhost:8080"+reset)

	http.ListenAndServe(":8080", r)
}
