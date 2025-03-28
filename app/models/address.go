package models

import "time"

type Address struct {
	ID 			string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User 		User
	UserID 		string `gorm:"size:36;index"`
	Name 		string `gorm:"size:255"`
	IsPrimary 	bool
	CityID 		string `gorm:"size:100"`
	Address1 	string `gorm:"size:255"`
	Address2 	string `gorm:"size:255"`
	Phone 		string `gorm:"size:255"`
	Email 		string `gorm:"size:255"`
	PostCode 	string `gorm:"size:255"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	}