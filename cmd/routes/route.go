package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {
	group := router.Group("/api")
	RegisterSync(group, 1, db)
	RegisterAuth(group, 1, db)
}
