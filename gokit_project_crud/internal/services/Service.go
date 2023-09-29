package services

import (
	"context"
	"github.com/saurabhkanawade/gokit_project_crud/internal/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateStudent(ctx context.Context, student repository.StudentRequest) (repository.StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]repository.Student, error)
	GetStudentById(ctx context.Context, id int) (repository.Student, error)
	UpdateStudent(ctx context.Context, student repository.Student, studentId int) (repository.Student, error)
	DeleteStudentById(ctx context.Context, id int) error
}

type service struct {
	repository repository.Repository
	log        logrus.Logger
}

func (s service) CreateStudent(ctx context.Context, student repository.StudentRequest) (repository.StudentResponse, error) {
	data, err := s.repository.CreateStudent(ctx, student)
	s.log.Info("service () - created the new student into the database..")
	logrus.Info("service () - request :", data)
	if err != nil {
		s.log.Fatal(logrus.ErrorLevel, err)
	}

	return repository.StudentResponse{
		Id: data.Id,
	}, nil

}

func (s service) GetAllStudent(ctx context.Context) ([]repository.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetStudentById(ctx context.Context, id int) (repository.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateStudent(ctx context.Context, student repository.Student, studentId int) (repository.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteStudentById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.Repository, log logrus.Logger) Service {
	return &service{
		repository: repo,
		log:        log,
	}
}
