package main

import (
	"log"
	"net/http"

	"go.course.todo-api/controller"
	"go.course.todo-api/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
