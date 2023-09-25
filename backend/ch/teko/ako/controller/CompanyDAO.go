package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type newCompanyJson struct {
	Name string `json:"name"`
}

type companyJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (h Handler) getAllCompanies(context *gin.Context) {
	companies := h.repository.GetAllCompanies()
	allCompanies := make([]companyJson, len(companies))
	for index, company := range companies {
		allCompanies[index] = companyJson{Id: company.Id, Name: company.Name}
	}
	context.JSON(http.StatusOK, allCompanies)
}

func (h Handler) CreateCompany(context *gin.Context) {
	company := newCompanyJson{}
	err := context.BindJSON(&company)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.CreateCompany(company.Name)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, companyJson{id, company.Name})
}
