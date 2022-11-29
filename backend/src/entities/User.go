package entities

import (
	"regexp"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type User struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()

	return nil
}

func (user User) validate() error {
	return validation.ValidateStruct(
		&user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&user.Email, validation.Required, validation.Match(regexp.MustCompile("[^@ \\t\\r\\n]+@[^@ \\t\\r\\n]+\\.[^@ \\t\\r\\n]+"))),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 50)),
	)
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}
