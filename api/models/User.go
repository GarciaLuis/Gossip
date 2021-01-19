package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User represents the user for this application
//
// A user can create, delete, and edit their own posts
// swagger:model
type User struct {
	// the id for this user, created by the database
	// required: false
	// min: 1
	ID uint32 `json:"id" gorm:"primary_key;auto_increment"`
	// the nickname/username for the user
	// required: true
	Nickname string `json:"nickname" gorm:"size:255;not null;unique"`
	// the email address for the user
	// required: true
	// example: user@email.com
	Email string `json:"email" gorm:"size:100;not null;unique"`
	// the user's login password
	// required: true
	Password string `json:"password" gorm:"size:100;not null;"`
	// the user's age
	// required: true
	Age uint8 `json:"age" gorm:"size:15"`
	// the user's gender
	// required: true
	Gender uint8 `json:"gender" gorm:"size:2"`
	// the user's height in inches
	// required: true
	Height float64 `json:"height" gorm:"size:15"`
	// the user's weight in lbs
	// required: true
	Weight float64 `json:"weight" gorm:"size:100"`
	// The time that the user record was created in db
	// read only: true
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	// The time the user record is updated in the db
	// read only: true
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Hash questions: what are bcrypt costs?
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares the hashed password with the given password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave stores the stringyfied hashed password into the password field of struct
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate ensures that when one of these actions (update, login) want to be performed
// that there must be a nickname, password, and valid email
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Nickname is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	}
}

//SaveUser is used to save the user into the database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

//FindAllUsers returns a max of 100 records of users
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User with specified ID not found")
	}

	return u, err
}

func (u *User) UpdateUserAccount(db *gorm.DB, uid uint32) (*User, error) {

	// Hash password:
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"nickname":   u.Nickname,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// Take updated user record
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateUserAttributes(db *gorm.DB, uid uint32) (*User, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"age":        u.Age,
			"weight":     u.Weight,
			"height":     u.Height,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	// Take updated user record
	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (u *User) FindUserByEmail(db *gorm.DB, email string) (*User, error) {

	var err error
	err = db.Debug().Model(&User{}).Where("email = ?", email).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User with specified email not found")
	}

	return u, err
}

func (u *User) UpdateWeight(db *gorm.DB, uid uint32, weight float64) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumn("weight", weight)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateHeight(db *gorm.DB, uid uint32, height float64) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumn("height", height)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateAge(db *gorm.DB, uid uint32, age uint8) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumn("age", age)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) GetWeightAndHeight(db *gorm.DB, uid uint32) (weight, height float64, err error) {

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return weight, height, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return weight, height, errors.New("User with the specified ID not found")
	}

	return u.Weight, u.Height, err
}
