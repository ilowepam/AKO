package controller

import (
	"AKO/ch/teko/ako/model"
)

type newAdAJson struct {
	Name      string     `json:"name"`
	Rank      model.Rank `json:"rank"`
	CompanyId int        `json:"companyId"`
}

type adaJson struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Rank      model.Rank `json:"rank"`
	CompanyId int        `json:"companyId"`
}
