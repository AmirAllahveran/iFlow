package model

import "github.com/jinzhu/gorm"

type Package struct {
	gorm.Model
	UserID      uint
	Name        string
	User        User
	Modules     []Module     `gorm:"foreignKey:PackageID"`
	Experiments []Experiment `gorm:"foreignKey:PackageID"`
}
