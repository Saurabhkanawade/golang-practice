package router

import (
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/employee_dept_array_crud/controller"
	"log"
	"net/http"
)

func RouterFunc() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/employees", controller.CreateEmployee).Methods("POST")
	route.HandleFunc("/employees", controller.GetAllEmployees).Methods("GET")
	route.HandleFunc("/employees/{id}", controller.GetEmployeesById).Methods("GET")
	route.HandleFunc("/employees/{id}", controller.UpdateEmployee).Methods("PUT")
	route.HandleFunc("/employees/{id}", controller.DeleteEmployee).Methods("Delete")

	http.ListenAndServe(":9000", route)

	log.Println("Started the server on 9000")

	return route
}
