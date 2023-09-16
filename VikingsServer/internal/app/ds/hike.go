package ds

import (
	"gorm.io/gorm"
	"time"
)

type Hike struct {
	gorm.Model
	ID          uint       `json:"id" gorm:"primary_key"`
	HikeName    string     `json:"hikeName" gorm:"type:varchar(50);not null"`
	DateStart   time.Time  `json:"dateStart" gorm:"type:date;not null;default:current_date"`
	DateEnd     time.Time  `json:"dateEnd" gorm:"type:date"`
	AuthorID    uint       `json:"authorID"`
	Author      Author     `json:"author" gorm:"foreignkey:AuthorID"`
	StatusID    uint       `json:"statusID"`
	Status      HikeStatus `json:"status" gorm:"foreignkey:StatusID"`
	Description string     `json:"description" gorm:"type:text"`
}