package repositories

import "github.com/gin-gonic/gin"

func GetDonors() (any) {
    return gin.H{"message": "success"}
}
