package model

import (
	"time"
)

func (f *FileMeta) TableName() string {
	return "filemeta"
}

type FileMeta struct {
	Id        int       `gorm:"primaryKey;autoIncrement:true;column:id;" json:"id"`
	Length    int64     `gorm:"primaryKey;autoIncrement:true;column:length;" json:"length"`
	Type      string    `gorm:"column:type" json:"type"`
	Name      string    `gorm:"column:name" json:"name"`
	UserId    int       `gorm:"column:userId" json:"userId"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}
