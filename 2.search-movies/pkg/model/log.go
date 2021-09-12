package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Log struct {
	URL            string    `gorm:"url"`
	ResponseStatus int       `gorm:"response_status"`
	CreatedAt      time.Time `gorm:"created_at"`
}

//BeforeCreate what to do before creating data to databases
func (log *Log) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

//TableName return the name of table wants to create in database
func (log *Log) TableName() string {
	return "logs"
}
