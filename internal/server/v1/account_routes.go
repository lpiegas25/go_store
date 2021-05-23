package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
	"strconv"
)

type AccountRouter struct {
	Repository account.Repository
}

func (ar *AccountRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var a account.Account
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = ar.Repository.Create(ctx, &a)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), a.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"account": a})
}

func (ar *AccountRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	ac, err := ar.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"account": ac})
}

func (ar *AccountRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var ac account.Account
	err = json.NewDecoder(r.Body).Decode(&ac)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = ar.Repository.Update(ctx, uint(id), ac)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (ar *AccountRouter) DeleteHandler(w http.ResponseWriter, r *http.Request)  {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ar.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (ar *AccountRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", ar.CreateHandler)
	r.Get("/{id}", ar.GetOneHandler)
	r.Put("/{id}", ar.UpdateHandler)
	r.Delete("/{id}", ar.DeleteHandler)

	return r
}