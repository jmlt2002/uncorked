package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmlt2002/uncorked/backend/internal/db"
)

type WineTagHandler struct {
	dbConn *sql.DB
}

func NewWineTagHandler(dbConn *sql.DB) *WineTagHandler {
	return &WineTagHandler{dbConn: dbConn}
}

func (h *WineTagHandler) CreateWineTag(c *gin.Context) {
	var req struct {
		WineID int64 `json:"wine_id" binding:"required"`
		TagID  int64 `json:"tag_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	wineTag, err := db.CreateWineTag(h.dbConn, context.Background(), req.WineID, req.TagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not associate tag with wine"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tag associated with wine", "wine_tag": wineTag})
}

func (h *WineTagHandler) DeleteWineTag(c *gin.Context) {
	var req struct {
		WineID int64 `json:"wine_id" binding:"required"`
		TagID  int64 `json:"tag_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := db.DeleteWineTag(h.dbConn, context.Background(), req.WineID, req.TagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not remove tag from wine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag removed from wine"})
}
