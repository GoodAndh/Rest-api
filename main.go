package main

import (
	"net/http"
	"restlast/app"
	"restlast/controller"
	"restlast/helper"
	"restlast/middlerware"
	"restlast/repository"
	"restlast/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	db := app.NewDb()

	categoryRepository := repository.NewRepository()
	categoryService := service.NewService(categoryRepository, db)
	categoryController := controller.NewController(categoryService)

	router.GET("/api", categoryController.GetAll)
	router.GET("/api/:id", categoryController.Getid)
	router.DELETE("/api/:id", categoryController.Delete)
	router.POST("/api", categoryController.Create)
	router.PUT("/api/:id", categoryController.Update)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.PageNotFound(w, r, nil)
	})
	router.PanicHandler=helper.PanicMessage

	http.ListenAndServe(":3000", middlerware.NewAuthMiddleware(router))
	
}
