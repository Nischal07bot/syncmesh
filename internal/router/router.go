package router

import (
	"net/http"
	"github.com/Nischal07bot/syncmesh/internal/handler/websocket"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", websocket.Handle)
	return mux
}
