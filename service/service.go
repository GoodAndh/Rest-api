package service

import "restlast/model"

//interaksi client
type Service interface {
	GetAll() []model.InteraksiClient
	Getid(id int) model.InteraksiClient
	Delete(id int) 
	Create(input model.Create)model.InteraksiClient
	Update(input model.Update)model.InteraksiClient

}
