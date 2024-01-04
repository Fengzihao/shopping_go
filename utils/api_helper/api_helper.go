package api_helper

import (
	"github.com/gin-gonic/gin"
	"shopping_go/utils/pagination"
)

var userIdText = "userId"

// GetUserId 从context中获取用户id
func GetUserId(g *gin.Context) uint {
	return uint(pagination.ParseInt(g.GetString(userIdText), -1))
}
