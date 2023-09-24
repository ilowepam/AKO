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
