package repository

import (
	"database/sql"
	"restlast/model"
)

type Repository interface {
	GetAll(Db *sql.DB) []model.InteraksiDb
	Getid(Db *sql.DB,input int) (model.InteraksiDb, error)
	Delete(Db *sql.DB,input int)
	Create(Db *sql.DB,input model.InteraksiDb)model.InteraksiDb
	Update(Db *sql.DB,input model.InteraksiDb)model.InteraksiDb
}
