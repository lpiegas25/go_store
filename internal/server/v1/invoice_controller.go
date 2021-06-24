package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/invoice"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
	"strconv"
)

type InvoiceController struct {
	Repository *invoice.Repository
}

func (ic *InvoiceController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	invoices, err := ic.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}
	response.JSON(w, r, http.StatusOK, response.Map{"invoices": invoices})
}

func (ic *InvoiceController) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var invoice invoice.Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = ic.Repository.Create(ctx, &invoice)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), invoice.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"invoice": invoice})
}

func (ic *InvoiceController) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	invoice, err := ic.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"invoice": invoice})
}

//
//func (ic *InvoiceController) UpdateHandler(w http.ResponseWriter, r *http.Request) {
//	idStr := chi.URLParam(r, "id")
//
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
//		return
//	}
//	var payment payment.Payment
//	err = json.NewDecoder(r.Body).Decode(&payment)
//	if err != nil {
//		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
//		return
//	}
//	defer r.Body.Close()
//
//	ctx := r.Context()
//	err = ic.Repository.Update(ctx, uint(id), payment)
//	if err != nil {
//		response.HTTPError(w, r, http.StatusNotFound, err.Error())
//		return
//	}
//	response.JSON(w, r, http.StatusOK, nil)
//}

func (ic *InvoiceController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", ic.CreateHandler)
	r.Get("/{id}", ic.GetOneHandler)
	//r.Put("/{id}", ic.UpdateHandler)
	//r.Delete("/{id}", ic.DeleteHandler)

	return r
}
