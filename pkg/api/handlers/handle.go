package handlers

import "net/http"

func HandleAdapter(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		hf(w, r)
	}
}
