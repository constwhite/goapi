package main

import (
	"fmt"
	"net/http"

	"github.com/constwhite/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
}
