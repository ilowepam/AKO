package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type lessonResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Time int    `json:"time"`
	Desc string `json:"desc"`
}

func (h Handler) getLessons(context *gin.Context) {

}

func (h Handler) getLessonByID(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	lesson, err := h.repository.GetLessonById(id)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, lessonResponse{Id: lesson.Id, Name: lesson.Name, Time: lesson.Time, Desc: lesson.Desc})
}

func (h Handler) createLesson(context *gin.Context) {

}

func (h Handler) updateLesson(context *gin.Context) {

}

func (h Handler) deleteLesson(context *gin.Context) {

}
