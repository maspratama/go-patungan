package entity

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id             int       `json:"id" gorm:"AUTO_INCREMENT; primary_key;"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Occupation     string    `json:"occupation"`
	PasswordHash   string    `json:"password"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
