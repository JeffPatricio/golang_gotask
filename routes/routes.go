package routes

import (
	"gotask/controllers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{task_id}", controllers.UpdateTask).Methods("PATCH")
	router.HandleFunc("/tasks/{task_id}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", controllers.PostTask).Methods("POST")

	return router
}
