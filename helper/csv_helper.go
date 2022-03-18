package helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FAdemoglu/homeworkthree/domain/book"
)

func ReadCsv(fileName string) ([]book.Book, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close() //Başına defer koyulan yer fonksiyon bitince en son satırdır yani burada bütün işlemlerimiz bitince okuma kısmını kapatmış olacağız.

	reader := csv.NewReader(f)

	lines, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var result []book.Book

	for _, line := range lines[1:] {
		stockCode, _ := strconv.Atoi(line[1])
		isbnNumber, _ := strconv.Atoi(line[2])
		pageNumber, _ := strconv.Atoi(line[3])
		price, _ := strconv.Atoi(line[4])
		stockCount, _ := strconv.Atoi(line[5])
		isDeleted, _ := strconv.ParseBool(line[7])
		data := book.Book{
			BookName:   line[0],
			StockCode:  stockCode,
			ISBNNumber: isbnNumber,
			PageNumber: pageNumber,
			Price:      price,
			StockCount: stockCount,
			Author:     line[6],
			IsDeleted:  isDeleted,
		}

		result = append(result, data)
	}

	return result, nil

}

func LowerCaseString(s string) string {
	return strings.ToLower(s)
}

func Contains(list []book.Book, searches string) {
	var searchedItems []string
	for _, v := range list {
		if strings.Contains(v.Author, searches) {
			searchedItems = append(searchedItems, v.Author)
		}
		if strings.Contains(v.BookName, searches) {
			searchedItems = append(searchedItems, v.BookName)
		}
	}
	fmt.Printf("Aranan kelimedeki yazarlar ve kitaplar %v", searchedItems)
}
