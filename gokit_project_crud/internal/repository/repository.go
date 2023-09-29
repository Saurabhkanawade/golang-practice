package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/saurabhkanawade/gokit_project_crud/models"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"math/big"
)

type Repository interface {
	CreateStudent(ctx context.Context, student StudentRequest) (StudentResponse, error)
	GetAllStudent(ctx context.Context) ([]Student, error)
	GetStudentById(ctx context.Context, id int) (Student, error)
	UpdateStudent(ctx context.Context, student Student, studentId int) (Student, error)
	DeleteStudentById(ctx context.Context, id int) error
}

type repository struct {
	Db     sql.DB
	Logger logrus.Logger
}

func NewRespostiry(db sql.DB, log logrus.Logger) Repository {
	return &repository{
		Db:     db,
		Logger: log,
	}
}

//implementation of the functions which interact with db

func (r *repository) CreateStudent(ctx context.Context, student StudentRequest) (StudentResponse, error) {
	logrus.Info("repository () - repository called ")
	u, _ := uuid.NewRandom()
	studentId := uuidToInt64(u)

	create := models.Student{
		ID:      studentId,
		Name:    student.Name,
		College: student.College,
		Gender:  student.Gender,
	}

	err := create.Insert(ctx, &r.Db, boil.Infer())
	if err != nil {
		panic(err)
	}

	return StudentResponse{
		Id: create.ID,
	}, nil
}

func (r *repository) GetAllStudent(ctx context.Context) ([]Student, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetStudentById(ctx context.Context, id int) (Student, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateStudent(ctx context.Context, student Student, studentId int) (Student, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteStudentById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func uuidToInt64(u uuid.UUID) int64 {
	// Convert the UUID to a big.Int value
	uInt := new(big.Int)
	uInt.SetString(u.String(), 16)

	// Convert the big.Int to an int64 value
	return uInt.Int64()
}
