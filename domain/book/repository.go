package book

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) FindAll() []Book {
	var books []Book
	r.db.Find(&books)
	return books
}

func (r *BookRepository) InsertCsvDatas(books []Book) {
	for _, c := range books {
		r.db.Where(Book{ISBNNumber: c.ISBNNumber}).Attrs(Book{BookName: c.BookName, StockCode: c.StockCode, ISBNNumber: c.ISBNNumber, PageNumber: c.PageNumber, Price: c.Price, StockCount: c.StockCount, Author: c.Author, IsDeleted: c.IsDeleted}).FirstOrCreate(&c)
	}
}

func (r *BookRepository) Create(b Book) error {
	result := r.db.Create(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *BookRepository) SearchByAuthorAndBookName(searched string) []Book {
	var books []Book
	r.db.Where("Author LIKE ?", "%"+searched+"%").Or("BookName LIKE ?", "%"+searched+"%").Find(&books)
	fmt.Println(books)
	return books
}

func (r *BookRepository) DeleteById(id int) error {
	var exists bool
	result := r.db.Delete(&Book{}, id)

	if err := result.Scan(&exists); err != nil {
		fmt.Printf("Bu id ile bir kitap bulunamadı")
		return result.Error
	} else if !exists {
		fmt.Printf("Bu id ile bir kitap bulunamadı")
		return nil
	}
	if result.Error != nil {
		fmt.Printf("Bu id ile kayıtlı bir kitap bulunmamakta")
		return result.Error
	}
	fmt.Printf("Silme işlemi başarılı")
	return nil
}

func (r *BookRepository) Update(id int, count int) error {
	var book Book
	result := r.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Bu id ile bir kitap bulunmuyor id : %d", id)
		return result.Error
	}
	if result.Error != nil {
		return result.Error
	}
	book.StockCount -= count
	resultSave := r.db.Save(book)
	if resultSave.Error != nil {
		return resultSave.Error
	}
	fmt.Printf("Kitap satma işlemi başarılı")
	return nil

}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) InsertSampleData() {
	books := []Book{
		{BookName: "Savas Sanati", StockCode: 123123123, ISBNNumber: 123123123, PageNumber: 237, Price: 20, StockCount: 15, Author: "Machiavelli", IsDeleted: false},
	}

	for _, c := range books {
		r.db.Where(Book{ISBNNumber: c.ISBNNumber}).Attrs(Book{BookName: c.BookName, StockCode: c.StockCode, ISBNNumber: c.ISBNNumber, PageNumber: c.PageNumber, Price: c.Price, StockCount: c.StockCount, Author: c.Author, IsDeleted: c.IsDeleted}).FirstOrCreate(&c)
	}
	fmt.Printf("Başarılı bir şekilde veriler eklendi\n")
}
