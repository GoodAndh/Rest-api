package repository

import (
	"database/sql"
	"fmt"
	"restlast/helper"
	"restlast/model"
)

type RepositoryImpl struct {
}

func NewRepository() Repository {

	return &RepositoryImpl{}

}
func (repo *RepositoryImpl) GetAll(Db *sql.DB) []model.InteraksiDb {
	script := "select id,name from category;"
	var category []model.InteraksiDb
	rows, err := Db.Query(script)
	helper.ErrorFatal(err)
	for rows.Next() {
		Isian := model.InteraksiDb{}
		defer rows.Close()
		err := rows.Scan(&Isian.Id, &Isian.Name)
		helper.ErrorFatal(err)
		category = append(category, Isian)

	}
	fmt.Println("REPOOOOOOOOOOOOOOOOOOOO", category)
	return category
}

func (repo *RepositoryImpl) Getid(Db *sql.DB, input int) (model.InteraksiDb, error) {
	script := "select id,name from category where id = ? ;"
	var category model.InteraksiDb
	rows, err := Db.Query(script, input)

	helper.ErrorFatal(err)
	if rows.Next() {

		defer rows.Close()
		err := rows.Scan(&category.Id, &category.Name)
		helper.ErrorFatal(err)

		return category, nil

	} else {
		return category, err
	}

}

func (repo *RepositoryImpl) Delete(Db *sql.DB, input int) {
	script := "delete from category where id = ?"
	_, err := Db.Exec(script, input)
	helper.ErrorFatal(err)

}
func (repo *RepositoryImpl) Create(Db *sql.DB, input model.InteraksiDb) model.InteraksiDb {
	script := "insert into category(name) values(?) "
	result, err := Db.Exec(script, input.Name)
	helper.ErrorFatal(err)

	id, err := result.LastInsertId()
	helper.ErrorFatal(err)
	input.Id = int(id)
	return input

}

func (repo *RepositoryImpl) Update(Db *sql.DB, input model.InteraksiDb) model.InteraksiDb {
	script := "update category set name = ? where id = ?"
	result, err := Db.Exec(script, input.Name, input.Id)
	helper.ErrorFatal(err)
	_, err = result.LastInsertId()
	helper.ErrorFatal(err)
	return input

}
