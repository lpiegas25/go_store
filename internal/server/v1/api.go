package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/lpiegas25/go_store/internal/data"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/model/client"
	"github.com/lpiegas25/go_store/pkg/model/employee"
	"github.com/lpiegas25/go_store/pkg/model/role"
	"github.com/lpiegas25/go_store/pkg/model/warehouse"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	ar := &AccountRouter{Repository: &account.AccountRepository{Data: data.New()}}
	cr := &ClientRouter{Repository: &client.ClientRepository{Data: data.New()}}
	rr := &RoleRouter{Repository: &role.RoleRepository{Data: data.New()}}
	wr := &WarehouseRouter{Repository: &warehouse.WarehouseRepository{Data: data.New()}}
	er := &EmployeeRouter{Repository: &employee.EmployeeRepository{Data: data.New()}}

	r.Mount("/accounts", ar.Routes())
	r.Mount("/clients", cr.Routes())
	r.Mount("/roles", rr.Routes())
	r.Mount("/warehouses", wr.Routes())
	r.Mount("/employees", er.Routes())

	return r
}
