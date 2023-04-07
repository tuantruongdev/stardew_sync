package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type World struct {
	Id        int        `json:"id" gorm:"id"`
	OwnerId   int        `json:"ownerId" gorm:"owner_id"`
	WorldName string     `json:"worldName" gorm:"world_name"`
	CreatedAt *time.Time `json:"createdAt" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"updated_at"`
	ImageUrl  string     `json:"imageUrl" gorm:"image_url"`
}

func (*World) TableName() string {
	return "save_world"
}

func InsertWorld(db *gorm.DB, world *World) bool {
	if err := db.Create(&world).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func QueryWorldByNameAndOwnerId(worldName string, ownerId int, db *gorm.DB) ([]World, error) {
	var worlds []World
	//should not use * here tbh
	selectQuery := "*"
	if err := db.Select(selectQuery).Where("world_name=? and owner_id=?", worldName, ownerId).Find(&worlds).Error; err != nil {
		return []World{}, err
	}
	return worlds, nil
}
func QueryWorldByOwnerId(ownerId int, db *gorm.DB) ([]World, error) {
	var worlds []World
	//should not use * here tbh
	selectQuery := "*"
	if err := db.Select(selectQuery).Where("owner_id=?", ownerId).Find(&worlds).Error; err != nil {
		return []World{}, err
	}
	return worlds, nil
}
