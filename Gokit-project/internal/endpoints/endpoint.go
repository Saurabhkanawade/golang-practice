package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/saurabhkanawade/internal/models"
	"github.com/saurabhkanawade/internal/services"
	"github.com/sirupsen/logrus"
	"log"
)

type Endpoints struct {
	Hello          endpoint.Endpoint
	GetStudents    endpoint.Endpoint
	GetStudentById endpoint.Endpoint
	CreateStudent  endpoint.Endpoint
	UpdateStudent  endpoint.Endpoint
	DeleteStudent  endpoint.Endpoint
}

func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		Hello:          MakeHelloEndpoint(s),
		GetStudents:    MakeGetStudentEndpoint(s),
		GetStudentById: MakeGetStudentByIdEndpoint(s),
		CreateStudent:  MakeCreateStudentEndpoint(s),
		UpdateStudent:  MakeUpdateStudentEndpoint(s),
		DeleteStudent:  MakeDeleteStudentEndpoint(s),
	}
}

func MakeGetStudentEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - MakeGetStudentEndpoint endpoint has been called.")
	return func(ctx context.Context, response interface{}) (interface{}, error) {

		students, err := s.GetStudents(ctx)
		handleNilError(err)
		return students, nil
	}
}

func MakeGetStudentByIdEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - MakeGetStudentByIdEndpoint endpoint has been called.")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "", nil
	}
}

func MakeCreateStudentEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - MakeCreateStudentEndpoint endpoint has been called.")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "saurabh", nil
	}
}

func MakeUpdateStudentEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - MakeUpdateStudentEndpoint endpoint has been called.")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "", nil
	}
}

func MakeDeleteStudentEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - MakeDeleteStudentEndpoint endpoint has been called.")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return "", nil
	}

}

func MakeHelloEndpoint(s services.Service) endpoint.Endpoint {
	logrus.Info("endpoint() - endpoint has been called.")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloRequest)
		return s.Hello(ctx, req.Name)
	}
}

type HelloRequest struct {
	Name string `json:"name"`
}

type GetStudentsRequest struct {
	StudentId   int64  `json:"studentId"`
	CollegeName string `json:"collegeName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetStudentResponse struct {
	Students []models.Student `json:"students"`
}

type GetStudentByIdResponse struct {
	StudentId int64 `json:"studentId"`
}

func handleNilError(err error) {
	if err != nil {
		log.Fatalln("error")
	}
}
