package campaign

import (
	"gorm.io/gorm"
	"time"
)

type CampaignsImages struct {
	ID         int       `json:"id" gorm:"AUTO_INCREMENT; primary_key;"`
	CampaignID int       `json:"campaign_id"`
	FileName   string    `json:"file_name"`
	IsPrimary  int       `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *CampaignsImages) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *CampaignsImages) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
