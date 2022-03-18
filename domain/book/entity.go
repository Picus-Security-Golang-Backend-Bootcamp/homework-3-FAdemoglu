package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName   string
	StockCode  int
	ISBNNumber int
	PageNumber int
	Price      int
	StockCount int
	Author     string
	IsDeleted  bool
}

func (Book) TableName() string {
	return "Book"
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Code : %s, CountryCode : %s,CreatedAt : %s", b.ISBNNumber, b.BookName, b.StockCode, b.ISBNNumber, b.PageNumber)
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("City (%s) deleting... \n", b.BookName)
	return nil
}
