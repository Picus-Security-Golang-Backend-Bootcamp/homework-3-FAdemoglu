package infrastructure

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewMySQLDB(conString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "book_",
			SingularTable: true,
			NoLowerCase:   true,
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database : %s", err.Error()))
	}

	return db
}
