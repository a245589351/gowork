package model

type Hello struct {
	ID int64 `gorm:"primary_key;not null;auto_increment"`
}
