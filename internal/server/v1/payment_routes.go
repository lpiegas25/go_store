package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/payment"
	"github.com/lpiegas25/go_store/pkg/response"
)

type PaymentRouter struct {
	Repository *payment.Repository
}

func (pr *PaymentRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payments, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"payments": payments})
}

func (pr *PaymentRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var payment payment.Payment
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = pr.Repository.Create(ctx, &payment)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), payment.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"payment": payment})
}

func (pr *PaymentRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	payment, err := pr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"payment": payment})
}

func (pr *PaymentRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var payment payment.Payment
	err = json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Update(ctx, uint(id), payment)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, nil)
}

func (pr *PaymentRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = pr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (pr *PaymentRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", pr.CreateHandler)
	r.Get("/{id}", pr.GetOneHandler)
	r.Put("/{id}", pr.UpdateHandler)
	r.Delete("/{id}", pr.DeleteHandler)
	r.Get("/", pr.GetAllHandler)

	return r
}
