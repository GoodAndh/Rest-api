package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restlast/helper"
	"restlast/model"
	"restlast/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ControllerImpl struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &ControllerImpl{
		service: service,
	}
}

func (controller *ControllerImpl) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	categoryResponse := controller.service.GetAll()
	webResponse := model.MainWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	fmt.Println("=========", categoryResponse)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.ErrorFatal(err)

}

func (controller *ControllerImpl) Getid(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	lastUrl := params.ByName("id")
	id, err := strconv.Atoi(lastUrl)
	// fmt.Println("LAST URL", id)
	helper.ErrorFatal(err)
	categoryResponse := controller.service.Getid(id)
	if categoryResponse.Id == 0 {
		webResponse := model.MainWeb{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "data yang diminta tidak ditemukan / tidak ada",
		}
		encoder := json.NewEncoder(w)
		err = encoder.Encode(webResponse)
		helper.ErrorFatal(err)
		return
	}
	webResponse := model.MainWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.ErrorFatal(err)

}

func (controller *ControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	lastUrl := params.ByName("id")
	id, err := strconv.Atoi(lastUrl)
	helper.ErrorFatal(err)
	controller.service.Delete(id)
	webResponse := model.MainWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Selected id deleted",
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.ErrorFatal(err)
}

func (controller *ControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	create := model.Create{}

	err := decoder.Decode(&create)
	helper.ErrorFatal(err)
	category := controller.service.Create(create)
	fmt.Println("CREATE :",create)
	fmt.Println(category)

	webResponse := model.MainWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   category,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.ErrorFatal(err)

}

func (controller *ControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	lastUrl := params.ByName("id")
	id, err := strconv.Atoi(lastUrl)
	helper.ErrorFatal(err)
	decoder := json.NewDecoder(r.Body)
	update := model.Update{
		Id: id,
	}
	err = decoder.Decode(&update)
	helper.ErrorFatal(err)
	category := controller.service.Update(update)
	webResponse := model.MainWeb{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   category,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.ErrorFatal(err)
}
