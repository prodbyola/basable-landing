package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	ch "basable/internal/core/handlers"
	cs "basable/internal/domain/collaborator"
	cr "basable/internal/infrastructure/persistence/collaborator"
)

type RouterInterface interface {
	IndexRoutes(router *gin.Engine)
	CollaboratorRoutes()
}

type Routes struct {
	db     *gorm.DB
	router *gin.RouterGroup
}

func NewRouter(db *gorm.DB, router *gin.RouterGroup) RouterInterface {
	return &Routes{
		db:     db,
		router: router,
	}
}

func (r *Routes) IndexRoutes(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"statusText": "success",
			"statusCode": http.StatusOK,
			"message":    "Welcome to Basable Landing API",
		})
	})

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"statusText": "success",
			"statusCode": http.StatusOK,
			"message":    "Server is up and running",
		})
	})
}

func (r *Routes) CollaboratorRoutes() {
	collaboratorRepository := cr.NewCollaboratorMySQLRepository(r.db)
	collaboratorService := cs.NewCollaboratorService(collaboratorRepository)
	collaboratorHandler := ch.NewCollaboratorHandler(collaboratorService)
	collaboratorRouter := r.router.Group("/v1/collaborators")

	{
		collaboratorRouter.POST("", collaboratorHandler.CreateCollaborator)
		collaboratorRouter.GET("", collaboratorHandler.GetAllCollaborators)
		collaboratorRouter.GET(":id", collaboratorHandler.GetCollaboratorByID)
		collaboratorRouter.PATCH(":id", collaboratorHandler.UpdateCollaborator)
		collaboratorRouter.DELETE(":id", collaboratorHandler.DeleteCollaborator)
	}
}

func LoadRoutes(db *gorm.DB, appRouter *gin.Engine) {
	api := appRouter.Group("api")
	router := NewRouter(db, api)

	router.IndexRoutes(appRouter)
	router.CollaboratorRoutes()
}
