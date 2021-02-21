package models

import (
	"gorm.io/gorm"
	"time"
)


type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
  	CreatedAt time.Time		 `json:"created_at"`
  	UpdatedAt time.Time		 `json:"updated_at"`
  	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	BookName string `json:"book_name" gorm:"size:191;index"`
	BookAuthor string `json:"book_author" gorm:"size:191;index"`
	BookYear string `json:"book_year" gorm:"size:191;index"`
}