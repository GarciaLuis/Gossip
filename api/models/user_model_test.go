package models

import (
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var userInstance = User{}

func TestFindAllUsers(t *testing.T) {

	err := RefreshUserTable(testDB)
	if err != nil {
		log.Fatal(err)
	}

	err = SeedUsers(testDB)
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(testDB)
	if err != nil {
		t.Errorf("Error returned when retrieving all users: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := RefreshUserTable(testDB)
	if err != nil {
		log.Fatal(err)
	}

	newUser := User{
		ID:       1,
		Email:    "skye@email.com",
		Nickname: "Skye",
		Password: "dummypassword",
	}

	savedUser, err := newUser.SaveUser(testDB)
	if err != nil {
		t.Errorf("Error returned while saving new user: %v\n", err)
		return
	}

	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestGetUserByID(t *testing.T) {

	err := RefreshUserTable(testDB)
	if err != nil {
		log.Fatal(err)
	}

	user, err := SeedOneUser(testDB)
	if err != nil {
		log.Fatalf("Error returned while seeding 1 user: %v\n", err)
	}
	foundUser, err := userInstance.FindUserByID(testDB, user.ID)
	if err != nil {
		t.Errorf("Error returned while retrieving 1 user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateUser(t *testing.T) {

	err := RefreshUserTable(testDB)
	if err != nil {
		log.Fatal(err)
	}

	user, err := SeedOneUser(testDB)
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	modifiedUser := User{
		ID:       1,
		Nickname: "Darcy",
		Email:    "Darcy@email.com",
		Password: "dummypassword",
	}

	updatedUser, err := modifiedUser.UpdateUserAccount(testDB, user.ID)
	if err != nil {
		assert.Equal(t, updatedUser.ID, modifiedUser.ID)
		assert.Equal(t, updatedUser.Email, modifiedUser.Email)
		assert.Equal(t, updatedUser.Nickname, modifiedUser.Nickname)
	}
}

func TestDeleteUser(t *testing.T) {

	err := RefreshUserTable(testDB)
	if err != nil {
		log.Fatal(err)
	}

	user, err := SeedOneUser(testDB)
	if err != nil {
		log.Fatalf("Error seeding 1 user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteUser(testDB, user.ID)
	if err != nil {
		t.Errorf("Error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
