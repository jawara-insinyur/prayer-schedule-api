package main

import (
	"log"
	"net/http"

	"github.com/jawara-insinyur/prayer-schedule-api/handler"
	"github.com/ringsaturn/tzf"
)

var f tzf.F

func init() {
	var err error
	f, err = tzf.NewDefaultFinder()
	if err != nil {
		panic(err)
	}
}

func main() {
	h := &handler.PrayScheduleHandler{Finder: f}
	router := initializeRoutes(h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening in port 8080...")
	server.ListenAndServe()
}

func initializeRoutes(h *handler.PrayScheduleHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /api/prayer-schedule", handler.RootHandler(h.PrayScheduleHandler))
	return mux
}
