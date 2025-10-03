package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmlt2002/uncorked/backend/internal/db"
)

type TagHandler struct {
	dbConn *sql.DB
}

func NewTagHandler(dbConn *sql.DB) *TagHandler {
	return &TagHandler{dbConn: dbConn}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var req struct {
		UserID  int64  `json:"user_id" binding:"required"`
		TagName string `json:"tag_name" binding:"required"`
		Color   string `json:"color" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create tag"})
		return
	}

	location, err := db.CreateTag(h.dbConn, context.Background(), req.UserID, req.TagName, req.Color)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create tag"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tag created", "location": location})
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id" binding:"required"`
		UserID int64 `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := db.DeleteTag(h.dbConn, context.Background(), req.ID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successful"})
}
