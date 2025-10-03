package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmlt2002/uncorked/backend/internal/db"
)

type StorageLocationHandler struct {
	dbConn *sql.DB
}

func NewStorageLocationHandler(dbConn *sql.DB) *StorageLocationHandler {
	return &StorageLocationHandler{dbConn: dbConn}
}

func (h *StorageLocationHandler) CreateStorageLocation(c *gin.Context) {
	var req struct {
		UserID int64  `json:"user_id" binding:"required"`
		Name   string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	location, err := db.CreateStorageLocation(h.dbConn, context.Background(), req.UserID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create storage location"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Storage location created", "location": location})
}

func (h *StorageLocationHandler) DeleteStorageLocation(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id" binding:"required"`
		UserID int64 `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := db.DeleteStorageLocation(h.dbConn, context.Background(), req.ID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete storage location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Storage location deleted successful"})
}
