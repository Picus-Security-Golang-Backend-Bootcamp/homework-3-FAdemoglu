package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/FAdemoglu/homeworkthree/domain/book"
	"github.com/FAdemoglu/homeworkthree/helper"
	"github.com/FAdemoglu/homeworkthree/infrastructure"
	valid "github.com/asaskevich/govalidator"
)

var (
	bookRepository *book.BookRepository
)

func init() {
	db := infrastructure.NewMySQLDB("root:Furkan7937.@tcp(127.0.0.1:3306)/books?parseTime=true&loc=Local")
	bookRepository = book.NewBookRepository(db)
	bookRepository.Migration()
	bookRepository.InsertSampleData()
	books, _ := helper.ReadCsv("book.csv")
	bookRepository.InsertCsvDatas(books)
}
func main() {
	args := os.Args
	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search=> arama işlemi için \n list=> listeleme işlemi için \n delete=> silme işlemi için \n buy=> kitabı satın almak için\n", projectName)
		return
	}

	argument := helper.LowerCaseString(args[1])

	if argument == "search" && len(args) < 3 {
		fmt.Printf("Aramak istediğiniz kitabı veya yazarı girmelisiniz!")
	} else if (argument == "buy" || argument == "delete") && len(args) >= 2 {
		for i := 2; i < len(args); i++ {
			if !valid.IsInt(args[i]) {
				fmt.Printf("Numerik bir değer girmelisiniz\n")
				return
			}

		}
	}
	switch argument {
	case "search":
		bookRepository.SearchByAuthorAndBookName(strings.Join(args[2:], " "))
	case "list":
		bookList := bookRepository.FindAll()
		fmt.Println(bookList)
	case "buy":
		if len(args) < 4 {
			fmt.Printf("Diğer komutları da girmelisiniz")
		} else {
			firstArgument, _ := strconv.Atoi(args[2])
			secondArgument, _ := strconv.Atoi(args[3])
			bookRepository.Update(firstArgument, secondArgument)
		}
	case "delete":
		if len(args) < 3 {
			fmt.Printf("Diğer komutları da girmelisiniz")
		} else {
			firstArgument, _ := strconv.Atoi(args[2])
			bookRepository.DeleteById(firstArgument)
		}
	default:
		fmt.Printf("Yanlış bir komut girdiniz.")
	}
	//books, _ := helper.ReadCsv("book.csv")
	//fmt.Println(books)

}
