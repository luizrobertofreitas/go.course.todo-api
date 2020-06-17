package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.course.todo-api/model"
	"go.course.todo-api/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.RequestURI)
		jsonEncoder := json.NewEncoder(w)
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				fmt.Println(err)
				errorData := views.ErrorResponse{
					Message: "Error when creating a new todo",
				}
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(errorData)
			} else {
				w.WriteHeader(http.StatusCreated)
				jsonEncoder.Encode(data)
			}
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.Read(name)
			if err != nil {
				fmt.Println(err)
				errorData := views.ErrorResponse{
					Message: "Cannot read todo data",
				}
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(errorData)
			} else {
				w.WriteHeader(http.StatusOK)
				jsonEncoder.Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				log.Fatal(err)
				errorData := views.ErrorResponse{
					Message: "Error when deleting the todo for " + name,
				}
				w.WriteHeader(http.StatusInternalServerError)
				jsonEncoder.Encode(errorData)
			} else {
				w.WriteHeader(http.StatusAccepted)
			}
		}
	}
}
