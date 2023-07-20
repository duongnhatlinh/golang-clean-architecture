package model

import (
	"errors"
	"food_delivery/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email"`
	Password        string        `json:"password" gorm:"column:password"`
	LastName        string        `json:"last_name" gorm:"column:last_name"`
	FirstName       string        `json:"first_name" gorm:"column:first_name"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"role" gorm:"column:role"`
	Salt            string        `json:"salt" gorm:"column:salt"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar"`
}

type UserUpdate struct {
	LastName  *string `json:"last_name" gorm:"column:last_name"`
	FirstName *string `json:"first_name" gorm:"column:first_name"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (u *User) Mask(isAdminOrUser bool) {
	u.GenUID(common.DbTypeUser)
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
