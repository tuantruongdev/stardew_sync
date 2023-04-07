package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stardew_sync/cmd/controllers"
)

func RegisterAuth(router *gin.RouterGroup, ver int, db *gorm.DB) {
	switch ver {
	case 1:
		group := router.Group(fmt.Sprintf("/v%d/account", ver))
		{
			group.POST("/login", controllers.Login(db)).
				GET("/auth", controllers.Authorize(db)).
				POST("/signup", controllers.SignUp(db)).
				POST("/logout", controllers.Authorize(db), controllers.Logout(db)).
				PATCH("/pass", controllers.Authorize(db), controllers.ChangePass(db)).
				POST("/forgot").
				PATCH("/name", controllers.Authorize(db), controllers.ChangeName(db))
		}
	}

}
