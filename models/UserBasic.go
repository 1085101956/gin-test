package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name         string
	Password     string
	Phone        string
	Email        string
	ClentIp      string
	ClientPort   string
	LoginTime    uint64
	HearbeatTime uint64
	LoginOutTime uint64
	IsLogin      bool
	DeviceInfo   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
