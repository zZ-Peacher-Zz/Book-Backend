package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name     string  `gorm:"unique_index;not null"`
	Author   string  `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	Comments []Comment
	Tags     []Tag `gorm:"many2many:book_tags;association_autocreate:false"`
}

type Comment struct {
	gorm.Model
	BookId uint `gorm:"not null"`
	Book   Book `gorm:"foreignKey:BookId"`
	User   User
	UserId uint
	Body   string
}

type Tag struct {
	gorm.Model
	Tag   string `gorm:"unique_index"`
	Books []Book `gorm:"many2many:book_tags;"`
}
