package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/role"
	"github.com/lpiegas25/go_store/pkg/response"
)

type RoleController struct {
	Repository *role.Repository
}

func (rr *RoleController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := rr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"roles": roles})
}

func (rr *RoleController) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var role role.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = rr.Repository.Create(ctx, &role)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), role.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"role": role})
}

func (rr *RoleController) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	role, err := rr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"role": role})
}

func (rr *RoleController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var role role.Role
	err = json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = rr.Repository.Update(ctx, uint(id), role)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (rr *RoleController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = rr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (rr *RoleController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", rr.CreateHandler)
	r.Get("/{id}", rr.GetOneHandler)
	r.Put("/{id}", rr.UpdateHandler)
	r.Delete("/{id}", rr.DeleteHandler)
	r.Get("/", rr.GetAllHandler)

	return r
}
