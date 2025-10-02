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
	r.DELETE("location", storageLocationHandler.DeleteStorageLocation)

	return r
}
