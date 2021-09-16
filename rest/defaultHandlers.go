package rest

import (
	"deni1688/myHealthTrack/domain"

	"github.com/gin-gonic/gin"
)

type defaultHandlers struct {
	service domain.Service
}

func (h *defaultHandlers) HandleCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleCreate not implmented",
	})
}

func (h *defaultHandlers) HandleGetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetAll not implmented",
	})
}

func (h *defaultHandlers) HandleGetOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HandleGetOne not implmented",
	})
}