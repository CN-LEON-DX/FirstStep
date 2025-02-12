package model

import (
	"awesomeProject/common"
	"database/sql/driver"
	"errors"
	"fmt"
)

const EntityName = "User"

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
	RoleShipper
	RoleMode
)

func (role UserRole) String() string {
	switch role {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	case RoleShipper:
		return "shipper"
	case RoleMode:
		return "mode"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB value:", value))

	}

	var r UserRole

	roleValue := string(bytes)

	if roleValue == "user" {
		r = RoleUser
	} else if roleValue == "admin" {
		r = RoleAdmin
	}
	*role = r

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}

	return role.String(), nil
}

func (role UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}

type User struct {
	common.SQLModel `json:",inline"`
	Email           string   `json:"email" gorm:"column:email"`
	Password        string   `json:"password" gorm:"column:password"`
	Salt            string   `json:"salt" gorm:"column:salt"`
	LastName        string   `json:"last_name" gorm:"column:last_name"`
	FirstName       string   `json:"first_name" gorm:"column:first_name"`
	Phone           string   `json:"phone" gorm:"column:phone"`
	Role            UserRole `json:"role" gorm:"column:role"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	Role            string `json:"role" gorm:"column:role"`
	Salt            string `json:"salt" gorm:"column:salt"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

// handel user login error
var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"UserLogin",
		"ErrEmailOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email existed"),
		"email existed",
		"UserCreate",
		"ErrEmailExisted",
	)
)
