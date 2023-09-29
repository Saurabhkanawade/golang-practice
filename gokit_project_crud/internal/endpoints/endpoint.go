package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/saurabhkanawade/gokit_project_crud/internal/repository"
	"github.com/saurabhkanawade/gokit_project_crud/internal/services"
	"github.com/sirupsen/logrus"
)

type Endpoints struct {
	CreateStudent     endpoint.Endpoint
	GetAllStudent     endpoint.Endpoint
	GetStudentById    endpoint.Endpoint
	UpdateStudent     endpoint.Endpoint
	DeleteStudentById endpoint.Endpoint
}

func MakeNewEndpoints(s services.Service) Endpoints {
	return Endpoints{
		CreateStudent: MakeCreateStudentEndpoint(s),
	}

}

func MakeCreateStudentEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("endpoint () - create student endpoint.")
		student := request.(repository.StudentRequest)
		logrus.Info("endpoint () - request :", request)
		res, err := s.CreateStudent(ctx, student)
		if err != nil {
			logrus.Fatal("error in create endpoint", err)
		}

		return repository.StudentResponse{
			Id: res.Id,
		}, nil
	}
}

// swagger:model Student
type Student struct {
	Name    string `json:"name"`
	College string `json:"college"`
	Gender  string `json:"gender"`
}

// Student
// swagger:response Student
type StudentResponse struct {
	//in:body
	ID string `json:"ID"`
}

