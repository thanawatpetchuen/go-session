package gos

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Format struct {
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/set", SetSession)
	router.HandleFunc("/get", GetSession)
	return router
}

func NewServer(h http.Handler) *http.Server {
	var logFormatter handlers.LogFormatter = func(writer io.Writer, params handlers.LogFormatterParams) {
		fields := Format{
			URL:       params.URL.Path,
			Timestamp: params.TimeStamp,
		}

		json.NewEncoder(writer).Encode(fields)
	}
	h = handlers.CustomLoggingHandler(log.Writer(), h, logFormatter)
	h = handlers.RecoveryHandler()(h)
	server := http.Server{
		Addr:    ":3000",
		Handler: h,
	}
	return &server
}
