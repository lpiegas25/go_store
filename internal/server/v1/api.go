package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/lpiegas25/go_store/internal/data"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/model/client"
	"github.com/lpiegas25/go_store/pkg/model/employee"
	"github.com/lpiegas25/go_store/pkg/model/payment"
	"github.com/lpiegas25/go_store/pkg/model/role"
	"github.com/lpiegas25/go_store/pkg/model/truck"
	"github.com/lpiegas25/go_store/pkg/model/warehouse"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	ar := &AccountRouter{Repository: &account.Repository{Data: data.New()}}
	cr := &ClientRouter{Repository: &client.Repository{Data: data.New()}}
	rr := &RoleRouter{Repository: &role.Repository{Data: data.New()}}
	wr := &WarehouseRouter{Repository: &warehouse.Repository{Data: data.New()}}
	er := &EmployeeRouter{Repository: &employee.Repository{Data: data.New()}}
	pr := &PaymentRouter{Repository: &payment.Repository{Data: data.New()}}
	tr := &TruckRouter{Repository: &truck.Repository{Data: data.New()}}

	r.Mount("/accounts", ar.Routes())
	r.Mount("/clients", cr.Routes())
	r.Mount("/roles", rr.Routes())
	r.Mount("/warehouses", wr.Routes())
	r.Mount("/employees", er.Routes())
	r.Mount("/payments", pr.Routes())
	r.Mount("/trucks", tr.Routes())

	return r
}
