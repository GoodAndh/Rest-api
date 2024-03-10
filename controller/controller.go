package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	GetAll(w http.ResponseWriter,r *http.Request,params httprouter.Params)
	Getid(w http.ResponseWriter,r *http.Request,params httprouter.Params)
	Delete(w http.ResponseWriter,r *http.Request,params httprouter.Params)
	Create(w http.ResponseWriter,r *http.Request,params httprouter.Params)
	Update(w http.ResponseWriter,r *http.Request,params httprouter.Params)
}