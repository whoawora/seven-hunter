package handler

import (
	"net/http"
	"pie-fire-dire/service"

	"github.com/gin-gonic/gin"
)

func GetBeefSummary(c *gin.Context) {
	beefCounts, err := service.FetchAndCountBeef()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := gin.H{"beef": beefCounts}
	c.JSON(http.StatusOK, response)
}
