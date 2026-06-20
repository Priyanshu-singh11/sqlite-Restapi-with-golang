package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Priyanshu-singh11/Restapi/internal/storage"
	"github.com/Priyanshu-singh11/Restapi/internal/types"
	"github.com/Priyanshu-singh11/Restapi/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// maxBodyBytes caps incoming JSON bodies so a malformed or hostile request
// can't exhaust server memory.
const maxBodyBytes = 1 << 20 // 1 MB

func New(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating a student")

		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			var validateErrs validator.ValidationErrors
			if errors.As(err, &validateErrs) {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		lastId, err := store.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		slog.Info("User created succesfully", slog.String("userId", fmt.Sprint(lastId)))
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student", slog.String("id", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("ERROR GETTING USER", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := store.GetStudentById(intId)

		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
				return
			}
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("List of all students")

		students, err := store.GetStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateStudentById(store storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

		var student types.Student

		err = json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		if err := validator.New().Struct(student); err != nil {
			var validateErrs validator.ValidationErrors
			if errors.As(err, &validateErrs) {
				response.WriteJson(
					w,
					http.StatusBadRequest,
					response.ValidationError(validateErrs),
				)
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		err = store.UpdateStudentById(
			intId,
			student.Name,
			student.Email,
			student.Age,
		)

		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
				return
			}
			response.WriteJson(
				w,
				http.StatusInternalServerError,
				response.GeneralError(err),
			)
			return
		}

		response.WriteJson(
			w,
			http.StatusOK,
			map[string]string{
				"message": "student updated successfully",
			},
		)
	}
}

func DeleteStudentById(store storage.Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GeneralError(err),
			)
			return
		}

		err = store.DeleteStudentById(intId)

		if err != nil {
			if errors.Is(err, storage.ErrNotFound) {
				response.WriteJson(w, http.StatusNotFound, response.GeneralError(err))
				return
			}
			response.WriteJson(
				w,
				http.StatusInternalServerError,
				response.GeneralError(err),
			)
			return
		}

		response.WriteJson(
			w,
			http.StatusOK,
			map[string]string{
				"message": "student deleted successfully",
			},
		)
	}
}
