package models

import (
	"github.com/FaizBShah/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	// Automatically migrate your schema, to keep your schema up to date
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	// It will change existing column’s type if its size, precision, nullable changed.
	// It WON’T delete unused columns to protect your data
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	if db.NewRecord(book) {
		db.Create(&book)
	}

	return book
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return book
}
