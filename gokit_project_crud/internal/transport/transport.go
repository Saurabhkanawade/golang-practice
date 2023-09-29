package transport

import (
	"context"
	"encoding/json"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/gokit_project_crud/internal/endpoints"
	"github.com/saurabhkanawade/gokit_project_crud/internal/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoint endpoints.Endpoints) http.Handler {
	route := mux.NewRouter()
	route.Use(middleWare)

	// swagger:operation POST /student create new portal
	// ---
	// summary: Creates the student with given payload
	// parameters:
	// - name: name
	//   in: body
	//   description: Name
	// - name: college
	//   in: body
	//   description: enter the college name
	// - name: gender
	//   in: body
	//   description: gender
	//   schema:
	//     "$ref": "#/definitions/Student"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/StudentResponse"
	//   "401":
	//     "$ref": "#/responses/httpErrorResponse"
	//   "404":
	//     "$ref": "#/responses/httpErrorResponse"
	//   "422":
	//     "$ref": "#/responses/httpErrorResponse"
	//   "500":
	//     "$ref": "#/responses/httpErrorResponse"

	route.Methods("POST").Path("/student").Handler(httpTransport.NewServer(
		endpoint.CreateStudent,
		decodeCreateStudentRequest,
		endcodeRequest,
	))
	return route
}

func endcodeRequest(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(response)
}

func decodeCreateStudentRequest(ctx context.Context, request2 *http.Request) (interface{}, error) {
	var student repository.StudentRequest
	json.NewDecoder(request2.Body).Decode(&student)
	logrus.Info("decoding the requrest", student)
	return student, nil
}

func middleWare(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/josn")
		handle.ServeHTTP(w, r)
	})
}
