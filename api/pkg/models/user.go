package models

import (
	"api/pkg/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique"`
	Password []byte    `json:"-"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u.Id = uuid.New()
	return
}

func GetUserByEmail(email string) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("email=?", email).Find(&getUser)
	return &getUser, db
}

func GetUserById(id uuid.UUID) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("id = ?", id).First(&getUser)
	return &getUser, db
}
