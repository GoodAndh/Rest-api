package helper

import "restlast/model"

func DbToClient(input model.InteraksiDb) model.InteraksiClient {
	return model.InteraksiClient{
		Id:   input.Id,
		Name: input.Name,
	}

}

func ToSlice(input []model.InteraksiDb) []model.InteraksiClient {
	var Toclient []model.InteraksiClient
	for _, v := range input {
		Toclient = append(Toclient, DbToClient(v))
	}
	return Toclient
}
