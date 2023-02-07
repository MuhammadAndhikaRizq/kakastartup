package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByemail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error //crate pada gorm berfungsi untuk menyimpan data

	if err != nil {
		return user, err

	}

	return user, nil
}

func (r *repository) FindByemail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("ID = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error //Save pada gorm berfungsi mengupdate data yang sudah ada di database

	if err != nil {
		return user, err
	}

	return user, nil
}
