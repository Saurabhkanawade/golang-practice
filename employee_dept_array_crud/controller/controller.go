package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/employee_dept_array_crud/model"
	"net/http"
)

var employees []model.Employee
var empUpdate model.Employee

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(&employees)

}
func GetEmployeesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ParamId := mux.Vars(r)

	for _, emp := range employees {
		if emp.EmpId == ParamId["id"] {
			_ = json.NewEncoder(w).Encode(&emp)
			return
		}
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var employee = model.Employee{
		EmpId: uuid.New().String(),
	}

	_ = json.NewDecoder(r.Body).Decode(&employee)

	employees = append(employees, employee)

	_ = json.NewEncoder(w).Encode(employee)

	return
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	paramId := mux.Vars(r)

	//for index, emp := range employees {
	//	if emp.EmpId == paramId["id"] {
	//		//deleting the existing emp
	//		employees = append(employees[:index], employees[index+1:]...)
	//		_ = json.NewDecoder(r.Body).Decode(&empUpdate)
	//		//adding updated one
	//		employees = append(employees, empUpdate)
	//		return
	//	}
	//}

	for _, emp := range employees {
		if emp.EmpId == paramId["id"] {
			_ = json.NewDecoder(r.Body).Decode(&empUpdate)
			empUpdate.EmpId = paramId["id"]
			employees = append(employees, empUpdate)
			_ = json.NewEncoder(w).Encode(empUpdate)
			return
		}
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	paramId := mux.Vars(r)

	for index, emp := range employees {
		if emp.EmpId == paramId["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			_ = json.NewEncoder(w).Encode("Deleted the employee")
			return
		}
	}

}
