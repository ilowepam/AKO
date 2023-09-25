package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type newLessonJson struct {
	Name string `json:"name"`
}

type lessonJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type lessonStatusJson struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type lessonStatusChangeJson struct {
	AdaId      int    `json:"adaId"`
	LessonName string `json:"lessonName"`
}

func (h Handler) CreateLessonStatus(context *gin.Context) {
	lessonStatus := lessonStatusChangeJson{}
	err := context.BindJSON(&lessonStatus)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	/*adaId, err := strconv.Atoi(lessonStatus.AdaId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "invalid adaId")
		return
	}*/
	adaId := lessonStatus.AdaId
	lesson := h.repository.GetLessonByName(lessonStatus.LessonName)
	commitStatus, err := h.repository.CreateLessonStatus(adaId, lesson.Id, true)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, commitStatus)
}

func (h Handler) CreateLesson(context *gin.Context) {
	lesson := newLessonJson{}
	err := context.BindJSON(&lesson)
	if err != nil {
		_ = context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	id, err := h.repository.CreateLesson(lesson.Name)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, lessonJson{id, lesson.Name})
}
