package middlewares

import (
	"net/htt"
)

func setmiddlewareJSON(next http.HandleFunc) ttp.HandleFunc {
	return func(w http.ResponseWriter, r http.Request {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

