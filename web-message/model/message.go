package model

import (
	"time"
)

type Message struct {
	Id        uint32  `gorm:"primaryKey;autoIncrement"`
	Name      string  `gorm:"type:varchar(50);not null;default:'';comment:注释"`
	Content   *string `gorm:"type:text"`
	CreatedAt uint32  `gorm:"UNSIGNED;not NULL;default:0"`
	UpdatedAt time.Time
}
