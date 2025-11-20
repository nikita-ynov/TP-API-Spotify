package router

import (
	"net/http"
	"spotify/controller"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/album/damso", controller.AlbumDamso)
	mux.HandleFunc("/track/laylow", controller.TrackLaylow)

	return mux
}
