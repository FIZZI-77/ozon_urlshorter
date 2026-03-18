package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"ozon/models"
)

func (h *Handler) CreateShort(c *gin.Context) {
	ctx := c.Request.Context()

	var req models.CreateRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	short, err := h.service.CreateShortUrl(ctx, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"short": short})
}

func (h *Handler) GetOriginal(c *gin.Context) {
	ctx := c.Request.Context()

	short := c.Param("short")

	url, err := h.service.GetOriginalUrl(ctx, short)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url)
}
