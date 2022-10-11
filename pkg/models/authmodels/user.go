package authmodels

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id            int    `json:"id,omitempty"`
	Login         string `json:"login,omitempty"`
	Password      string `json:"password,omitempty"`
	Comment       string `json:"comment,omitempty"`
	IsAdmin       bool   `json:"isAdmin,omitempty"`
	IsBlocked     bool   `json:"isBlocked,omitempty"`
	CreatedUserId int64  `json:"createdUserId,omitempty"`
	DateCreated   string `json:"dateCreated,omitempty"`
	DateChanged   string `json:"dateChanged,omitempty"`
	Token         string `json:"token,omitempty"`
}

type Response struct {
	Id int64 `json:"id"`
}

func (u *User) ComparePassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}

func (u *User) CreateToken() string {
	return uuid.New().String()
}
