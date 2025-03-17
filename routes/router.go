package routes

import (
	"gos/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()


    donors := r.Group("/donor")
    {
        donors.GET("/", services.GetDonors)
    }

    return r
}
