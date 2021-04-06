package gos

import "net/http"

func ContentJSON(rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
}
