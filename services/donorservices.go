package services

import (
	"gos/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDonors(c *gin.Context) {
    c.JSON(http.StatusOK, repositories.GetDonors())
}
