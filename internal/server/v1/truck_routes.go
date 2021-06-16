package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/truck"
	"github.com/lpiegas25/go_store/pkg/response"
)

type TruckRouter struct {
	Repository *truck.Repository
}

func (tr *TruckRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	trucks, err := tr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"trucks": trucks})
}

func (tr *TruckRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var truck truck.Truck
	err := json.NewDecoder(r.Body).Decode(&truck)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = tr.Repository.Create(ctx, &truck)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), truck.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"truck": truck})
}

func (tr *TruckRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	truck, err := tr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"truck": truck})
}

func (tr *TruckRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var truck truck.Truck
	err = json.NewDecoder(r.Body).Decode(&truck)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = tr.Repository.Update(ctx, uint(id), truck)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (tr *TruckRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = tr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (tr *TruckRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", tr.CreateHandler)
	r.Get("/{id}", tr.GetOneHandler)
	r.Put("/{id}", tr.UpdateHandler)
	r.Delete("/{id}", tr.DeleteHandler)
	r.Get("/", tr.GetAllHandler)

	return r
}
