package controller

import (
	"AKO/ch/teko/ako/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type overviewDataJson struct {
	Headers []string      `json:"headers"`
	Rows    []rowDataJson `json:"rows"`
}

type rowDataJson struct {
	Ada     adaJson            `json:"ada"`
	Lessons []lessonStatusJson `json:"lesson"`
}

func (h Handler) GetOverviewData(context *gin.Context) {
	adas := h.repository.GetAllAdAs()

	lessons := h.repository.GetAllLessons()
	rowData := make([]rowDataJson, len(adas))

	for index, ada := range adas {
		company, _ := h.repository.GetCompanyById(ada.CompanyId)
		adaData := adaJson{Id: ada.Id, Name: ada.Name, Rank: ada.Rank, CompanyName: company.Name}
		lessonData := make([]lessonStatusJson, len(lessons))
		for index, lesson := range lessons {
			lessStatus, _ := h.repository.GetLessonStatus(ada.Id, lesson.Id)
			lessonData[index] = lessonStatusJson{Name: lesson.Name, Status: lessStatus}
		}
		rowData[index] = rowDataJson{Ada: adaData, Lessons: lessonData}
	}
	adaHeaders := model.GetAdAHeaderNames()
	lessonHeaders := make([]string, len(lessons))
	lessonHData := make([]string, len(lessonHeaders))

	for index, lesson := range lessons {
		lessonHData[index] = lesson.Name
	}

	headerData := append(adaHeaders, lessonHData...)

	overviewData := overviewDataJson{Headers: headerData, Rows: rowData}
	context.JSON(http.StatusOK, overviewData)
}
