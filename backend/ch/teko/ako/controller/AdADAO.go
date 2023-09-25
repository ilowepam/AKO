package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type newAdAJson struct {
	Name        string `json:"name"`
	Rank        string `json:"rank"`
	CompanyName string `json:"companyName"`
}

type adaJson struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Rank        string `json:"rank"`
	CompanyName string `json:"companyName"`
}

func (h Handler) CreateAda(context *gin.Context) {
	ada := newAdAJson{}
	err := context.BindJSON(&ada)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.CreateAdA(ada.Name, ada.Rank, ada.CompanyName)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, adaJson{id, ada.Name, ada.Rank, ada.CompanyName})
}
