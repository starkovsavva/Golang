package api

import (
	"net/http"
)

type api struct {
	addr string
	r    *http.ServeMux
}

func New(addr string, r *http.ServeMux) *api {
	return &api{addr: addr, r: r}
}

func (api *api) FillEndPoints() {
	api.r.HandleFunc("/api/hello", hello)
	api.r.HandleFunc("/api/goodbye", goodbye)
}

func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("goodbye!\n"))
}
