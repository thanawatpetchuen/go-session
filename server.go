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
	Timestamp time.Time
	URL       string
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", MyHandler)
	router.HandleFunc("/get", func(rw http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(session.Values["foo"]); err != nil {
			log.Fatal(err)
		}
	})
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
