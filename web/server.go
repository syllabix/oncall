package web

import (
	"net/http"

	"github.com/syllabix/oncall/api/rest"
	"github.com/syllabix/oncall/config"
	"github.com/syllabix/oncall/web/home"
)

type Server struct {
	*http.Server
}

func NewServer(settings config.ServerSettings, api rest.Router, home home.Page) Server {
	mux := http.NewServeMux()
	mux.Handle("/", root(api, home))

	return Server{
		Server: &http.Server{
			Addr:         settings.Host + ":" + settings.Port,
			Handler:      mux,
			ReadTimeout:  settings.ReadTimeout,
			WriteTimeout: settings.WriteTimeout,
		},
	}
}

func root(api, home http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			home.ServeHTTP(w, r)
			return
		}
		api.ServeHTTP(w, r)
	})
}
