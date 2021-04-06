package gos

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var name = "key"
var store = sessions.NewCookieStore([]byte(name))

func SetSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "session-name")

	body := LoginBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["user"] = body.Username
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSession(rw http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	user := User{
		Username: session.Values["user"].(string),
	}
	ContentJSON(rw)
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(user); err != nil {
		log.Fatal(err)
	}
}
