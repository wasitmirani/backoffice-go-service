package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"-"`
	ProfilePic string `json:"profile_pic"`
}