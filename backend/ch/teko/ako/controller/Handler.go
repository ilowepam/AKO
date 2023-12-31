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

	overviewGroup := router.Group("/overview")
	overviewGroup.GET("/", h.GetOverviewData)

	lessonGroup := router.Group("/lesson")
	lessonGroup.POST("/status/", h.CreateLessonStatus)
	lessonGroup.POST("/", h.CreateLesson)

	adaGroup := router.Group("/ada")
	adaGroup.POST("/", h.CreateAda)

	companyGroup := router.Group("/company")
	companyGroup.POST("/", h.CreateCompany)
	companyGroup.GET("/all", h.GetAllCompanies)

	rankGroup := router.Group("/rank")
	rankGroup.GET("/all", h.getAllRanks)

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
