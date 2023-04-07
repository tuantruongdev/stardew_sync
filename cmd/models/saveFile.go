package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type WorldFile struct {
	Id           int        `json:"id" gorm:"id"`
	OwnerId      int        `json:"ownerId" gorm:"owner_id"`
	WorldOwnerId int        `json:"worldOwnerId" gorm:"world_owner_id"`
	Description  string     `json:"description" gorm:"description"`
	Path         string     `json:"path" gorm:"path"`
	Favorite     int        `json:"favorite" gorm:"favorite"`
	CreatedAt    *time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt    *time.Time `json:"updatedAt" gorm:"updated_at"`
	ImageUrl     string     `json:"imageUrl" gorm:"image_url"`
}

type WorldFileRequest struct {
	Description string `json:"description"`
	Favorite    int    `json:"favorite"`
	ImageUrl    string `json:"imageUrl"`
}

func (*WorldFile) TableName() string {
	return "save_file"
}

func InsertFile(db *gorm.DB, file *WorldFile) bool {
	if err := db.Create(&file).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
func QueryFile(db *gorm.DB, file *WorldFile) error {
	if err := db.Where("id=?", file.Id).Find(&file).Error; err != nil {
		return err
	}
	return nil
}
