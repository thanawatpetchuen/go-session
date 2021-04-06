package gos

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

const COOKIE_STORE_NAME = "key"
const SESSION_NAME = "gos"
const COOKIE_TTL = 10

var store = sessions.NewCookieStore([]byte(COOKIE_STORE_NAME))

func SetSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, SESSION_NAME)

	body := LoginBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["user"] = body.Username
	session.Values["id"] = uuid.New().String()
	// Save it before we write to the response/return from the handler.
	// session.Options.Secure = true
	session.Options.MaxAge = COOKIE_TTL

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSession(rw http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SESSION_NAME)
	userName, ok := GetUser(rw, session)
	if !ok {
		return
	}
	sessionId, ok := GetSessionId(rw, session)
	if !ok {
		return
	}
	user := User{
		Username: userName,
	}
	response := NewSessionResponse(user, sessionId)
	JSONResponse(rw, response)
}
