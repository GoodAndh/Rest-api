package helper

import (
	"encoding/json"
	"net/http"
	"restlast/model"

	"github.com/julienschmidt/httprouter"
)

func ErrorFatal(err error) {
	if err != nil {
		panic(err)
	}
}

func PageNotFound(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	response := model.MainWeb{
		Code:   http.StatusNotFound,
		Status: "error not found",
		Data:   "url yang diminta tidak ada",
	}
	bro, err := json.Marshal(response)
	ErrorFatal(err)
	w.Write(bro)
}

func PanicMessage(w http.ResponseWriter, r *http.Request, i interface{}) {
	w.Header().Add("Content-Type", "application/json")

	response := model.MainWeb{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
		Data:   i,
	}
	bro, err := json.Marshal(response)
	ErrorFatal(err)
	w.Write(bro)
}

