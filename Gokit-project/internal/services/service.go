package services

import (
	"context"
	"github.com/saurabhkanawade/internal/models"
	"github.com/sirupsen/logrus"
	"log"
)

type Service interface {
	Hello(ctx context.Context, name string) (string, error)
	GetStudents(ctx context.Context) ([]models.Student, error)
	GetStudentById(ctx context.Context, getStudentById string) (string, error)
	CreateStudent(ctx context.Context, createStudent string) (string, error)
	UpdateStudent(ctx context.Context, updateStudent string) (string, error)
	DeleteStudent(ctx context.Context, deleteStudent string) (string, error)
}

type MyService struct {
	//repository repository.Repository
	//log        logrus.Logger
}

// function of service implementation of interfact

func (s *MyService) Hello(ctx context.Context, name string) (string, error) {
	logrus.Info("service () - hello called")
	if name == "" {
		log.Println("name cannot be empty")
	}
	return "Hello ," + name + "", nil
}

func (s *MyService) GetStudents(ctx context.Context) ([]models.Student, error) {
	logrus.Info("service () - getStudent called")
	//student array
	student := []models.Student{
		{
			StudentId: 2, StudentName: "saurabh", CollegeName: "smbst", Address: "sangamner", PhoneNumber: "8779122363",
		},
	}

	return student, nil
}

func (s *MyService) GetStudentById(ctx context.Context, studentId string) (string, error) {
	return "", nil
}

func (s *MyService) CreateStudent(ctx context.Context, student models.Student) (string, error) {
	return "", nil
}

func (s *MyService) UpdateStudent(ctx context.Context, studentId string) (string, error) {
	return "", nil
}

func (s *MyService) DeleteStudent(ctx context.Context) (string, error) {
	return "", nil
}
