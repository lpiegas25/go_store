package v1

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/pkg/model/delivery"
	"github.com/lpiegas25/go_store/pkg/response"
	"net/http"
	"strconv"
)

type DeliveryController struct {
	Repository *delivery.Repository
}

func (dc *DeliveryController) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var delivery delivery.Delivery
	err := json.NewDecoder(r.Body).Decode(&delivery)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err = dc.Repository.Create(ctx, &delivery)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), delivery.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"delivery": delivery})
}

func (dc *DeliveryController) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	deliveryDTO, err := dc.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"delivery": deliveryDTO})
}

func (dc *DeliveryController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", dc.CreateHandler)
	r.Get("/{id}", dc.GetOneHandler)
	//r.Put("/{id}", ic.UpdateHandler)
	//r.Delete("/{id}", ic.DeleteHandler)

	return r
}
