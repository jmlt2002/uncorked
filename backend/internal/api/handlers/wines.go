package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmlt2002/uncorked/backend/internal/db"
)

type WineHandler struct {
	dbConn *sql.DB
}

func NewWineHandler(dbConn *sql.DB) *WineHandler {
	return &WineHandler{dbConn: dbConn}
}

func (h *WineHandler) CreateWine(c *gin.Context) {
	var wine db.Wine
	if err := c.ShouldBindJSON(&wine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	created, err := db.CreateWine(h.dbConn, context.Background(), &wine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create wine"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"wine": created})
}

func (h *WineHandler) GetWine(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	wine, err := db.GetWine(h.dbConn, context.Background(), req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wine not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wine": wine})
}

func (h *WineHandler) DeleteWine(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := db.DeleteWine(h.dbConn, context.Background(), req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete wine"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wine deleted"})
}

func (h *WineHandler) UpdateWine(c *gin.Context) {
	var wine db.Wine
	if err := c.ShouldBindJSON(&wine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := db.UpdateWine(h.dbConn, context.Background(), &wine); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update wine"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Wine updated"})
}

func (h *WineHandler) ListWinesByUser(c *gin.Context) {
	var req struct {
		UserID int64 `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	wines, err := db.ListWinesByUser(h.dbConn, context.Background(), req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not list wines"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"wines": wines})
}
