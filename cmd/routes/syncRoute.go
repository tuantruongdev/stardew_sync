package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"stardew_sync/cmd/controllers"
)

func RegisterSync(router *gin.RouterGroup, ver int, db *gorm.DB) {

	switch ver {
	case 1:
		group := router.Group(fmt.Sprintf("/v%d/sync", ver))
		{
			group.GET("world/:file", controllers.Authorize(db), controllers.GetSyncFile(db)).
				GET("/world", controllers.Authorize(db), controllers.GetSyncWorld(db)).
				POST("/", controllers.Authorize(db), controllers.BodySizeLimit(), controllers.PostSyncFile(db))
		}
	}

}
