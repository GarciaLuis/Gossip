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

func refreshUserTable() error {
	err := testDB.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	err = testDB.AutoMigrate(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func seedOneUser() (User, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal("Error refreshingUserTable: ", err)
	}

	user := User{
		Nickname: "Aurie",
		Email:    "Aurie@email.com",
		Password: "dummypassword",
	}

	err = testDB.Model(&User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []User{
		User{
			Nickname: "Lucky",
			Email:    "lucky@email.com",
			Password: "dummypassword",
		},
		User{
			Nickname: "Sope",
			Email:    "sope@email.com",
			Password: "dummypassword",
		},
	}

	for i := range users {
		err := testDB.Model(&User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func refreshUserAndPostTable() error {

	err := testDB.DropTableIfExists(&User{}, &Post{}).Error
	if err != nil {
		return err
	}
	err = testDB.AutoMigrate(&User{}, &Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return Post{}, err
	}

	user := User{
		Nickname: "Peluchin",
		Email:    "peluchin@email.com",
		Password: "dummypassword",
	}
	err = testDB.Model(&User{}).Create(&user).Error
	if err != nil {
		return Post{}, err
	}

	post := Post{
		Title:    "Peluchin's post title",
		Content:  "This is a blog post by peluchin",
		AuthorID: user.ID,
	}

	err = testDB.Model(&Post{}).Create(&post).Error
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func seedUsersAndPosts() ([]User, []Post, error) {

	var err error

	if err != nil {
		return []User{}, []Post{}, err
	}

	var users = []User{
		User{
			Nickname: "Shakira",
			Email:    "shakira@email.com",
			Password: "dummypassword",
		},
		User{
			Nickname: "Zimba",
			Email:    "zimba@email.com",
			Password: "dummypassword",
		},
	}

	var posts = []Post{
		Post{
			Title:   "Post title Num 1",
			Content: "Content of Post Num 1",
		},
		Post{
			Title:   "Post title Num 2",
			Content: "Content of Post Num 2",
		},
	}

	for i := range users {
		err = testDB.Model(&User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = testDB.Model(&Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("Cannot seed posts table: %v", err)
		}
	}

	return users, posts, nil
}
