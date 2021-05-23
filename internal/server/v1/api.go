package v1

import (
	"github.com/go-chi/chi"
	"github.com/lpiegas25/go_store/internal/data"
	"github.com/lpiegas25/go_store/pkg/model/account"
	"github.com/lpiegas25/go_store/pkg/model/client"
	"net/http"
)

func New() http.Handler {
	r := chi.NewRouter()

	ar := &AccountRouter{Repository: &account.AccountRepository{Data: data.New()}}
	cr := &ClientRouter{Repository: &client.ClientRepository{Data: data.New()}}

	r.Mount("/accounts", ar.Routes())
	r.Mount("/clients", cr.Routes())

	return r
}
