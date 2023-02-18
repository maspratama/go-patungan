package user

import "gorm.io/gorm"

type Repository interface {
	Save(user Users) (Users, error)
}

type repository struct {
	db *gorm.DB
}

// membuat var/obj baru dari repository struct
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
