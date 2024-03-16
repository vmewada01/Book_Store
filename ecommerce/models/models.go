package models

import "time"

type User struct {
	ID              int    `gorm:"primary_key"`
	First_Name      string `gorm:"unique"`
	Last_Name       string `gorm:"unique"`
	Email           string `gorm:"unique"`
	Password        string `gorm:"unique"`
	Phone           string `gorm:"unique"`
	Token           string `gorm:"unique"`
	Refresh_Token   string `gorm:"unique"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User_ID         int    `gorm:"primary_key"`
	UserCart        int    `gorm:"primary_key"`
	Address_Details string `gorm:"primary_key"`
	Order_Status    string `gorm:"primary_key"`
}
