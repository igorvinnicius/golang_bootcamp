package models

import (
	"strings"
)

const(
	ErrNotFound modelError = "models: resource not found"
	ErrIDInvalid privateError = "models: ID must me > 0"
	ErrPasswordIncorrect modelError = "models: incorrect password provided"
	ErrEmailRequired modelError = "models: email is required"
	ErrEmailInvalid  modelError = "models: email is required"
	ErrEmailTaken modelError = "models: email address is already taken"
	ErrPasswordTooShort modelError = "models: password must be at least 8 ccarachters long"
	ErrPasswordRequired modelError =  "models: password is required"
	ErrRememberTooShort privateError = "models: remember token must be at least 32 bytes"
	ErrRememberRequired privateError = "models: remember token is required"
	ErrUserIdRequired privateError = "models: user id is required"
	ErrTitleRequired modelError = "models: title is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}