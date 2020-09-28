package models

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {

	InitializeTestDB()
	os.Exit(m.Run())
}

func InitializeTestDB() {
	var err error

	TestDbDriver := os.Getenv("DB_POSTGRES_DRIVER")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
		testDB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
}
