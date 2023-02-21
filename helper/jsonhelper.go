package helper

import "net/http"

func Jsonhelper(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	return
}
