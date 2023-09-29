package transport

import (
	"context"
	"encoding/json"
	newHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/internal/endpoints"
	"github.com/saurabhkanawade/internal/models"
	"net/http"
	"strconv"
)

func MakeHTTPHandlers(ctx context.Context, endpoints endpoints.Endpoints) http.Handler {
	route := mux.NewRouter()
	route.Use(middleWare)

	route.Methods("GET").Path("/students").Handler(newHttp.NewServer(
		endpoints.GetStudents,
		decodeGetStudentRequest,
		encodeRequest,
	))

	route.Methods("POST").Path("/hello").Handler(newHttp.NewServer(
		endpoints.Hello,
		decodeHelloRequest,
		encodeRequest,
	))

	route.Methods("GET").Path("/student/{studentId}").Handler(newHttp.NewServer(
		endpoints.GetStudentById,
		decodeGetStudentByIdRequest,
		encodeRequest,
	))

	route.Methods("POST").Path("/student").Handler(newHttp.NewServer(
		endpoints.CreateStudent,
		decodeCreateStudent,
		encodeRequest,
	))

	route.Methods("/PUT").Path("/student/{id}").Handler(newHttp.NewServer(
		endpoints.UpdateStudent,
		decodeUpdateStudent,
		encodeRequest,
	))

	route.Methods("DELETE").Path("/student/{studentId}").Handler(newHttp.NewServer(
		endpoints.DeleteStudent,
		decodeDeleteStudent,
		encodeRequest,
	))
	return route
}

func decodeDeleteStudent(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	var vars = mux.Vars(request2)
	studentId, err := strconv.Atoi(vars["studentId"])

	if err != nil {
		return nil, err
	}
	student := &models.Student{
		StudentId: int64(studentId),
	}
	return student, nil

}

func decodeUpdateStudent(ctx context.Context, request2 *http.Request) (interface{}, error) {
	var student models.Student
	var vars = mux.Vars(request2)
	studentId, _ := strconv.Atoi(vars["studentId"])
	json.NewDecoder(request2.Body).Decode(&student)
	student.StudentId = int64(studentId)
	return student, nil
}

func decodeCreateStudent(ctx context.Context, request2 *http.Request) (interface{}, error) {
	var student endpoints.GetStudentsRequest
	json.NewDecoder(request2.Body).Decode(&student)
	return student, nil
}

func decodeHelloRequest(ctx context.Context, w *http.Request) (interface{}, error) {
	var request endpoints.HelloRequest
	if err := json.NewDecoder(w.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeGetStudentRequest(ctx context.Context, request2 *http.Request) (interface{}, error) {
	var students models.Student
	if err := json.NewDecoder(request2.Body).Decode(&students); err != nil {
		return nil, err
	}
	return students, nil
}

func decodeGetStudentByIdRequest(ctx context.Context, request2 *http.Request) (interface{}, error) {
	var request endpoints.GetStudentsRequest
	if err := json.NewDecoder(request2.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeRequest(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
	return json.NewEncoder(writer).Encode(&i)
}

func middleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(writer, request)
	})
}
