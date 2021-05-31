package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"strconv"

	"github.com/lpiegas25/go_store/pkg/model/client"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
)

type ClientRouter struct {
	Repository client.Repository
}

func (cr *ClientRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	clients, err := cr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"clients": clients})
}

func (cr *ClientRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	ac, err := cr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"client": ac})
}

func (cr *ClientRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var c client.Client
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = cr.Repository.Create(ctx, &c)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), c.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"client": c})
}

func (cr *ClientRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var c client.Client
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = cr.Repository.Update(ctx, uint(id), c)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (cr *ClientRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = cr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (cr *ClientRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", cr.CreateHandler)
	r.Get("/{id}", cr.GetOneHandler)
	r.Get("/", cr.GetAllHandler)
	r.Put("/{id}", cr.UpdateHandler)
	r.Delete("/{id}", cr.DeleteHandler)

	return r
}
