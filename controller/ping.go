package controller

import (
	"encoding/json"
	"net/http"

	"go.course.todo-api/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(data.Code)
			json.NewEncoder(w).Encode(data)
		}
	}
}
