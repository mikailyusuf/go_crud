package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	err = db.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id =?", id).Delete(User).Error
	if err != nil {
		return err
	}
	return nil
}
