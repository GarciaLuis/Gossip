package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Post represents a contribution that a user has made to the blog.
//
// A Post can be created, edited or deleted by its author.
// swagger:model
type Post struct {
	// ID that belongs to the post, assigned by the db
	// required: false
	// min: 1
	ID uint64 `json:"id" gorm:"primary_key;auto_increment"`
	// The post's title
	// Required: true
	Title string `json:"title" gorm:"size:255;not null;unique"`
	// Content holds the main information that the use wants to share
	// Required: true
	Content string `json:"content" gorm:"size:255;not null"`
	// Author is the user that created the post
	// required: false
	Author User `json:"user"`
	// ID of the author user
	// required: true
	AuthorID uint32 `json:"author_id" gorm:"not null"`
	// createdAt is the time that the post was created in the db
	// required: false
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	// updatedAt is the time that the post is updated
	// required: false
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	// Populate the Author field with the record matching the author id
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) FindAllPosts(db *gorm.DB) (*[]Post, error) {
	var err error
	posts := []Post{}
	err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}

func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {

	var err error

	err = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Post) DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Post) VerifyExistence(db *gorm.DB, pid uint64) (*Post, error) {

	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if db.Error != nil {
		return &Post{}, errors.New("Post doesn't exist")
	}

	if gorm.IsRecordNotFoundError(err) {
		return &Post{}, errors.New("Post with specified ID not found")
	}

	return p, err
}
