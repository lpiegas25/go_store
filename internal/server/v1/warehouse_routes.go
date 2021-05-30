package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/warehouse"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
	"strconv"
)

type WarehouseRouter struct {
	Repository warehouse.Repository
}

func (wr *WarehouseRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	warehouses, err := wr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"warehouses": warehouses})
}

func (wr *WarehouseRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var warehouse warehouse.Warehouse
	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = wr.Repository.Create(ctx, &warehouse)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), warehouse.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"warehouse": warehouse})
}

func (wr *WarehouseRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	warehouse, err := wr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"warehouse": warehouse})
}

func (wr *WarehouseRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var warehouse warehouse.Warehouse
	err = json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = wr.Repository.Update(ctx, uint(id), warehouse)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (wr *WarehouseRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = wr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (wr *WarehouseRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", wr.CreateHandler)
	r.Get("/{id}", wr.GetOneHandler)
	r.Put("/{id}", wr.UpdateHandler)
	r.Delete("/{id}", wr.DeleteHandler)
	r.Get("/", wr.GetAllHandler)

	return r
}
