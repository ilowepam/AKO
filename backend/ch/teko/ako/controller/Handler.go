package controller

import (
	"AKO/ch/teko/ako/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Handler struct {
	engine     *gin.Engine
	repository *repository.PostgresqlStorageImpl
}

func MakeApi() *Handler {
	h := &Handler{
		engine:     gin.Default(),
		repository: repository.NewPostgresStorageImpl(),
	}

	configureCors(h.engine)

	router := h.engine.Group("/api/ako/")

	trainingLessonGroup := router.Group("/traininglesson")
	trainingLessonGroup.GET("/all", h.getLessons)
	trainingLessonGroup.GET("/:id", h.getLessonByID)
	trainingLessonGroup.POST("/", h.createLesson)
	trainingLessonGroup.PUT("/:id", h.updateLesson)
	trainingLessonGroup.DELETE("/:id", h.deleteLesson)

	return h
}
func (h Handler) Run() {
	err := h.engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func configureCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	engine.Use(cors.New(config))
}
