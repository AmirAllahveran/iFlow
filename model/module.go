package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
)

type Module struct {
	gorm.Model
	UserID        uint
	PackageID     uint
	Name          string
	ImageName     string
	Count         int64
	User          User
	Package       Package
	ModuleConfigs []ModuleConfig `gorm:"foreignKey:ModuleID"`
}

type ModuleConfig struct {
	gorm.Model
	ExperimentID uint
	ModuleID     uint
	Cmd          string
	Config       JSON // Custom type to handle map[string]string
	Experiment   Experiment
	Module       Module
}

// JSON is a custom type to handle map[string]string, making it easier to work with in GORM
type JSON map[string]string

// Value makes the JSON type implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan makes the JSON type implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (j *JSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &j)
}
