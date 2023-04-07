package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"stardew_sync/cmd/models"
	"stardew_sync/cmd/utils"
	"strconv"
	"strings"
)

func GetSyncFile(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		fileIdStr := context.Param("file")
		savedFile := models.WorldFile{}
		fileId, err := strconv.Atoi(fileIdStr)
		if err != nil {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "no valid id")
			context.JSON(res.Generate())
			return
		}
		user, ok := context.Keys["user"].(models.User)
		if !ok {
			res.SetStatus(http.StatusUnauthorized, models.StatusError, "no auth")
			context.JSON(res.Generate())
			return
		}
		savedFile.Id = fileId
		models.QueryFile(db, &savedFile)
		// check if who request file is owner or not
		if savedFile.OwnerId != user.Id {
			res.SetStatus(http.StatusForbidden, models.StatusError, "you don't have permission")
			context.JSON(res.Generate())
			return
		}
		//check if the file exits in static folder
		if utils.FileExists(savedFile.Path) {
			nameSlice := strings.Split(savedFile.Path, "/")
			//fmt.Println(nameSlice)
			context.FileAttachment(savedFile.Path, nameSlice[len(nameSlice)-1])
			//	context.File(savedFile.Path)
		} else {
			res.SetStatus(http.StatusNotFound, models.StatusError, "the save file no longer available")
			context.JSON(res.Generate())
			return
		}
	}
}
func GetSyncWorld(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		user, ok := context.Keys["user"].(models.User)
		if !ok {
			res.SetStatus(http.StatusUnauthorized, models.StatusError, "no auth")
			context.JSON(res.Generate())
			return
		}
		worlds, err := models.QueryWorldByOwnerId(user.Id, db)
		if err != nil {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "no auth")
			context.JSON(res.Generate())
			return
		}
		res.SetStatus(http.StatusOK, models.StatusOk, "ok")
		res.Set("data", worlds)
		context.JSON(res.Generate())

		//context.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func PostSyncFile(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		res := models.NetResponse{}.Build()
		done := make(chan bool)
		isSaved := true
		syncFile, err := context.FormFile("file")
		user, ok := context.Keys["user"].(models.User)
		if !ok {
			res.SetStatus(http.StatusUnauthorized, models.StatusError, "no auth")
			context.JSON(res.Generate())
			return
		}
		//in case no file uploaded
		if err != nil {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "File not found")
			context.JSON(res.Generate())
			return
		}
		go func() {
			//check if file is valid
			isValid := utils.ValidateFile(syncFile)
			if isValid {
				//if valid then save
				ok, path := utils.SaveFile(syncFile)
				if ok {
					exits, world := isWorldExist(syncFile.Filename, user.Id, db)
					if !exits {
						//insert new world to db
						world = &models.World{
							Id:        0,
							OwnerId:   user.Id,
							WorldName: syncFile.Filename,
							CreatedAt: nil,
							UpdatedAt: nil,
							ImageUrl:  "",
						}
						models.InsertWorld(db, world)
					}

					//save path to db
					worldFile := models.WorldFile{
						Id:           0,
						OwnerId:      user.Id,
						WorldOwnerId: world.Id,
						Description:  "",
						Path:         path,
						Favorite:     0,
						CreatedAt:    nil,
						UpdatedAt:    nil,
						ImageUrl:     "",
					}
					models.InsertFile(db, &worldFile)
					_ = path
					isSaved = true
				}
			} else {
				isSaved = false
			}

			done <- true
		}()

		//wait for validate file done
		<-done
		if isSaved {
			res.SetStatus(http.StatusAccepted, models.StatusOk, "File Uploaded ")
			context.JSON(res.Generate())
		} else {
			res.SetStatus(http.StatusBadRequest, models.StatusError, "File Upload failed")
			context.JSON(res.Generate())
		}

		//isValid = utils.ValidateFile(syncFile)

	}
}
func BodySizeLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		var w http.ResponseWriter = context.Writer
		context.Request.Body = http.MaxBytesReader(w, context.Request.Body, 10*1024*1024)
		context.Next()
	}
}
func isWorldExist(worldName string, id int, db *gorm.DB) (bool, *models.World) {
	worlds, err := models.QueryWorldByNameAndOwnerId(worldName, id, db)
	if err != nil {
		return false, nil
	}
	if len(worlds) > 0 {
		return true, &worlds[0]
	}
	return false, nil
}
