package model

import uuid "github.com/satori/go.uuid"

type Person struct {
	ID   uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	Name string    `gorm:"column:name" json:"name"`
}

func (p *Person) TableName() string {
	return "u"
}
