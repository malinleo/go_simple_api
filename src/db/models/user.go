package models

import "gorm.io/gorm"

//Represent User data in db.
type User struct{
	gorm.Model
	Email string `gorm:"size:256;not null"`
}