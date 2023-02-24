package domains

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	password    string    `db:"password"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	if err != nil {
		return err
	}

	u.password = string(hashPassword)

	return nil
}

func (u *User) GetDBTableName() string {
	return "users_user"
}

func (u *User) GetDBValues() []any {
	return []any{uuid.New(), u.PhoneNumber, u.Email, u.password, time.Now().UTC(), time.Now().UTC()}
}

type CreateUser struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	password    string
}
