package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	name       *string
	categoryID *int
)

// VARIABLE => untuk menyimpan suatu data (e.g. name)
// string, float64, int, bool

// type data bentukan (struct)
type Category struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FUNCTION => untuk menyimpan kumpulan process (e.g. main())
// PACKAGE => untuk membungkus semua func / variables

func Greeting(n string) {
	// mengeluarkan output sapaan
	fmt.Println("Hello from First Golang CLI " + n)
}

func main() {

	// menerima input name
	// package flag untuk menerima input dari CLI
	name = flag.String("name", "john", "name input to be greet")
	// SECTION 1: menerima input name
	Greeting(*name)

	// SECTION 2: menerima input category id
	categoryID = flag.Int("category_id", 1, "product id input")
	flag.Parse()

	// SECTION 3: koneksi ke database menggunakan ORM (GORM)
	// ORM => layer untuk berkomunikasi dengan database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Jakarta application_name=%v",
		"127.0.0.1",
		"docker",
		"mysecretpassword",
		"product",
		"5432",
		"workshop")
	// db => koneksi ke db
	// err => menyimpan value failure
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// jika terjadi error
		// aplikasi akan keluar
		log.Fatal(err)
	}
	// kalau tidak ada error
	// koneksi db bisa kita gunakan

	// SECTION 4: query untuk mendapatkan category by id
	cat1 := Category{}
	db.
		Table("categories").
		Where("id = ?", categoryID).
		Find(&cat1)

	// SECTION 5: API untuk mendapatkan category by id
	// ngambil data
	gr := gin.Default()
	gr.GET("/product", func(ctx *gin.Context) {
		query := ctx.Query("id")

		cat1 := Category{}
		db.
			Table("categories").
			Where("id = ?", query).
			Find(&cat1)
		ctx.JSON(http.StatusOK, cat1)
	})
	gr.Run(":9091")
}
