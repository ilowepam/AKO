package controller

import (
	"AKO/ch/teko/ako/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newCompanyJson struct {
	Name      string    `json:"name"`
	Commander model.AdA `json:"commander"`
}

type companyJson struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Commander model.AdA `json:"commander"`
}

func (h Handler) createCompany(context *gin.Context) {
	company := newCompanyJson{}
	err := context.BindJSON(&company)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.CreateCompany(company.Name, company.Commander)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, companyJson{id, company.Name, company.Commander})
}
