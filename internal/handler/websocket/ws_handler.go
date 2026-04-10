package websocket

import "net/http"

// Handle is the HTTP entrypoint for websocket upgrades.
func Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	_, _ = w.Write([]byte("websocket handler not implemented"))
}
