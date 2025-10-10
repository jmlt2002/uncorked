package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jmlt2002/uncorked/backend/internal/api/handlers"
)

func NewRouter(dbConn *sql.DB) *gin.Engine {
	r := gin.Default()

	userHandler := handlers.NewUserHandler(dbConn)
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)

	storageLocationHandler := handlers.NewStorageLocationHandler(dbConn)
	r.POST("/location", storageLocationHandler.CreateStorageLocation)
	r.DELETE("/location", storageLocationHandler.DeleteStorageLocation)

	tagsHandler := handlers.NewTagHandler(dbConn)
	r.POST("/tag", tagsHandler.CreateTag)
	r.DELETE("/tag", tagsHandler.DeleteTag)

	wineTagHandler := handlers.NewWineTagHandler(dbConn)
	r.POST("/wine_tag", wineTagHandler.CreateWineTag)
	r.DELETE("/wine_tag", wineTagHandler.DeleteWineTag)

	wineHandler := handlers.NewWineHandler(dbConn)
	r.POST("/wine", wineHandler.CreateWine)
	r.GET("/wine", wineHandler.GetWine)
	r.DELETE("/wine", wineHandler.DeleteWine)
	r.PUT("/wine", wineHandler.UpdateWine)
	r.GET("/wines", wineHandler.ListWinesByUser)

	return r
}
