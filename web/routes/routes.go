package routes

import (
	"net/http"
	"todo-client/web/controllers"
)

func RegisterRoutes() {
	http.Handle("/", controllers.RootController)
	http.HandleFunc("/list", controllers.ListTasks)
	http.HandleFunc("/create", controllers.CreateTasks)
}
