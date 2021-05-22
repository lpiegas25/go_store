package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
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

func (ar *AccountRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", ar.CreateHandler)

	return r
}