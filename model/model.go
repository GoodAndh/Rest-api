package model

type MainWeb struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type InteraksiClient struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type InteraksiDb struct {
	Id   int
	Name string
}

type Create struct {
	Name string `json:"name"`
}


type Update struct{
	Id int `json:"id"`
	Name string `json:"name"`
}