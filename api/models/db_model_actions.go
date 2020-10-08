package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

func RefreshUserTable(testDB *gorm.DB) error {
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

func SeedOneUser(testDB *gorm.DB) (User, error) {

	err := RefreshUserTable(testDB)
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

func SeedUsers(testDB *gorm.DB) error {

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

func RefreshUserAndPostTable(testDB *gorm.DB) error {

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

func SeedOneUserAndOnePost(testDB *gorm.DB) (Post, error) {

	err := RefreshUserAndPostTable(testDB)
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

func SeedUsersAndPosts(testDB *gorm.DB) ([]User, []Post, error) {

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
