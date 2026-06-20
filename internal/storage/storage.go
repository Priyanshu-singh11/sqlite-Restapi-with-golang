package storage

import (
	"errors"

	"github.com/Priyanshu-singh11/Restapi/internal/types"
)

var ErrNotFound = errors.New("student not found")

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentById(id int64, name string, email string, age int) error
	DeleteStudentById(id int64) error
}
