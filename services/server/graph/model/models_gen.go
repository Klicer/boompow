// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	Type      UserType `json:"type"`
}

type UserInput struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Type     UserType `json:"type"`
}

type VerifyEmailInput struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type WorkGenerateInput struct {
	UserID               string `json:"userID"`
	APIKey               string `json:"apiKey"`
	Hash                 string `json:"hash"`
	DifficultyMultiplier int    `json:"difficultyMultiplier"`
}

type UserType string

const (
	UserTypeProvider  UserType = "PROVIDER"
	UserTypeRequester UserType = "REQUESTER"
)

var AllUserType = []UserType{
	UserTypeProvider,
	UserTypeRequester,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeProvider, UserTypeRequester:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
