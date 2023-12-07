package model

import "github.com/jinzhu/gorm"

type Experiment struct {
	gorm.Model
	UserID        uint
	PackageID     uint
	Title         string
	Description   string
	User          User
	Package       Package
	ModuleConfigs []ModuleConfig `gorm:"foreignKey:ExperimentID"`
}
