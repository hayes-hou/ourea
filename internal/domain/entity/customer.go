// Package entity Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package entity

import (
	"os"
	"time"
)

type UploadImg struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;" json:"name"`
	Path      string     `gorm:"size:100;" json:"path"`
	Url       string     `json:"url"`
	Content   os.File    `gorm:"-" json:"-"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
