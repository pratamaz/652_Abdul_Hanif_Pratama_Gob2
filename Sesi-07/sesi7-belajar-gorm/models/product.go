package models

import (
	"fmt"
	"time"
	"errors"
	"gorm.io/gorm"
)

type Product struct {
	ID 			string		`gorm:"primaryKey"`
	Name		string		`gorm:"not null; type:varchar(191)"`
	Brand		string		`gorm:"not null; type:varchar(191)"`
	UserID		uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product Name is Too Short")
	}	

	return
}