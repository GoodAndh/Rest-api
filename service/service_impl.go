package service

import (
	"database/sql"
	"fmt"
	"restlast/helper"
	"restlast/model"
	"restlast/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
	Db   *sql.DB
}

func NewService(repo repository.Repository, db *sql.DB) Service {
	return &ServiceImpl{
		Repo: repo,
		Db:   db,
	}
}
func (service *ServiceImpl) GetAll() []model.InteraksiClient {
	category := service.Repo.GetAll(service.Db)
	fmt.Println("SLICEEEEEEEE", category)
	return helper.ToSlice(category)
}

func (service *ServiceImpl) Getid(id int) model.InteraksiClient {
	category, err := service.Repo.Getid(service.Db, id)
	helper.ErrorFatal(err)
	return helper.DbToClient(category)
}

func (service *ServiceImpl) Delete(id int) {
	service.Repo.Delete(service.Db, id)
}

func (service *ServiceImpl) Create(input model.Create) model.InteraksiClient {
	create := model.InteraksiDb{
		Name: input.Name,
	}
	category := service.Repo.Create(service.Db, create)
	return helper.DbToClient(category)
}

func (service *ServiceImpl) Update(input model.Update) model.InteraksiClient {
	getid, err := service.Repo.Getid(service.Db, input.Id)
	helper.ErrorFatal(err)
	getid.Name = input.Name
	category := service.Repo.Update(service.Db, getid)
	return helper.DbToClient(category)
}
