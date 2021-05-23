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

func (cr *ClientRouter) CreateHandler(w http.ResponseWriter, r *http.Request)  {
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

// UpdateHandler ToDo: implement update handler
func (cr *ClientRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

// DeleteHandler ToDo: implement delete handler
func (cr *ClientRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (cr *ClientRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", cr.CreateHandler)
	r.Get("/{id}", cr.GetOneHandler)

	return r
}