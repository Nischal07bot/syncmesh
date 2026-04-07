package router

import "net/http"

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
