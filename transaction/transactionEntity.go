package transaction

import (
	"github.com/maspratama/go-patungan/campaign"
	"github.com/maspratama/go-patungan/user"
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	ID         int    `json:"id" gorm:"AUTO_INCREMENT; primary_key;"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
	User       user.Users
	Campaign   campaign.Campaigns
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *Transactions) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Transactions) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
