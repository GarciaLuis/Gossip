package models

import (
	"log"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

var postInstance = Post{}

func TestingFindAllPosts(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and posts tables %v\n", err)
	}
	_, _, err = seedUsersAndPosts()
	if err != nil {
		log.Fatalf("Error seeding users and posts into table: %v\n", err)
	}
	posts, err := postInstance.FindAllPosts(testDB)
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 2)
}

func TestSavePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and posts table: %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newPost := Post{
		ID:       1,
		Title:    "Post to test save",
		Content:  "This post is to test saving a post",
		AuthorID: user.ID,
	}

	savedPost, err := newPost.SavePost(testDB)
	if err != nil {
		log.Fatalf("Error returned when saving post: %v\n", err)
		return
	}

	assert.Equal(t, newPost.ID, savedPost.ID)
	assert.Equal(t, newPost.Title, savedPost.Title)
	assert.Equal(t, newPost.Content, savedPost.Content)
	assert.Equal(t, newPost.AuthorID, savedPost.AuthorID)

}

func TestGetPostByID(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}

	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatalf("Error seeding the table")
	}

	foundPost, err := postInstance.FindPostByID(testDB, post.ID)
	if err != nil {
		t.Errorf("Error retrieving post by id: %v\n", err)
		return
	}

	assert.Equal(t, foundPost.ID, post.ID)
	assert.Equal(t, foundPost.Title, post.Title)
	assert.Equal(t, foundPost.Content, post.Content)
}

func TestUpdatePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatal("Error Seeding table")
	}
	modifiedPost := Post{
		ID:       1,
		Title:    "Modified title",
		Content:  "content of modified post",
		AuthorID: post.AuthorID,
	}
	updatedPost, err := modifiedPost.UpdatePost(testDB)
	if err != nil {
		t.Errorf("Error returned when updating the post: %v\n", err)
		return
	}
	assert.Equal(t, updatedPost.ID, modifiedPost.ID)
	assert.Equal(t, updatedPost.Title, modifiedPost.Title)
	assert.Equal(t, updatedPost.Content, modifiedPost.Content)
	assert.Equal(t, updatedPost.AuthorID, modifiedPost.AuthorID)
}

func TestDeletePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table: %v\n", err)
	}
	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatalf("Error Seeding tables")
	}
	isDeleted, err := postInstance.DeletePost(testDB, post.ID, post.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
