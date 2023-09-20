package controller

import (
	"AKO/ch/teko/ako/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h Handler) createAdA(context *gin.Context) {
	ada := newAdAJson{}
	err := context.BindJSON(&ada)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.CreateAdA(ada.Name, ada.Rank)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, adaJson{id, ada.Name, ada.Rank, ada.CompanyId})
}

func (h Handler) getAllAdAByCompanyId(context *gin.Context) {
	param := context.Param("companyId")
	companyId, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "invalid companyId")
		return
	}
	adas := h.repository.GetAllAdAByCompanyId(companyId)
	remindersResponses := make([]adaJson, len(adas))
	for index, ada := range adas {
		remindersResponses[index] = adaJson{Id: ada.Id, Name: ada.Name, Rank: ada.Rank, CompanyId: ada.CompanyId}
	}
	context.JSON(http.StatusOK, remindersResponses)
}

func (h Handler) deleteAdA(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	deleted := h.repository.DeleteAdA(id)
	if deleted {
		context.Status(http.StatusOK)
	} else {
		context.JSON(http.StatusBadRequest, "Could not delete AdA")
	}
}
