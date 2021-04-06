package gos

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

func ContentJSON(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
}

func JSONResponse(rw http.ResponseWriter, v interface{}) {
	ContentJSON(rw)
	if err := json.NewEncoder(rw).Encode(v); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
}

func GetUser(rw http.ResponseWriter, session *sessions.Session) (string, bool) {
	userSession := session.Values["user"]
	if userSession == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("no user"))
		return "", false
	}
	userName, ok := userSession.(string)
	if !ok {
		http.Error(rw, strconv.FormatBool(ok), http.StatusInternalServerError)
		return "", false
	}
	return userName, true
}

func GetSessionId(rw http.ResponseWriter, session *sessions.Session) (string, bool) {
	idSession := session.Values["id"]
	if idSession == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("no session id"))
		return "", false
	}
	id, ok := idSession.(string)
	if !ok {
		http.Error(rw, strconv.FormatBool(ok), http.StatusInternalServerError)
		return "", false
	}
	return id, true
}
