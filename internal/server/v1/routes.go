package v1

import "net/http"

type Routeable interface {
	CreateHandler(w http.ResponseWriter, r *http.Request)
	GetOneHandler(w http.ResponseWriter, r *http.Request)
	UpdateHandler(w http.ResponseWriter, r *http.Request)
	DeleteHandler(w http.ResponseWriter, r *http.Request)
	Routes() http.Handler
}
