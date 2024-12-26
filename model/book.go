package model

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name   string `gorm:"unique_index;not null"`
	Author string `gorm:"not null"`
	Price  string `gorm:"not null"`
	Tags   []Tag  `gorm:"many2many:book_tags;association_autocreate:false"`
}

type Comment struct {
	gorm.Model
	Book   Book
	User   User
	UserId uint
}

type Tag struct {
	gorm.Model
	Tag   string `gorm:"unique_index"`
	Books []Book `gorm:"many2many:book_tags;"`
}
