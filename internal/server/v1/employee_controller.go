package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/employee"
	"github.com/lpiegas25/go_store/pkg/response"
)

type EmployeeController struct {
	Repository *employee.Repository
}

func (er *EmployeeController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	employees, err := er.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"employees": employees})
}

func (er *EmployeeController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = er.Repository.Create(ctx, &e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), e.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"employee": e})
}

func (er *EmployeeController) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	e, err := er.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"employee": e})
}

func (er *EmployeeController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var e employee.Employee
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = er.Repository.Update(ctx, uint(id), e)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (er *EmployeeController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = er.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (er *EmployeeController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", er.CreateHandler)
	r.Get("/{id}", er.GetOneHandler)
	r.Get("/", er.GetAllHandler)
	r.Put("/{id}", er.UpdateHandler)
	r.Delete("/{id}", er.DeleteHandler)

	return r
}
