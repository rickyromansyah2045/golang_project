package models

import "time"

type People struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Jeniskelamin string
	Email        string
	NoHp         uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
