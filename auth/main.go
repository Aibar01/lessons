package main

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Domain interface {
	GetDBTableName() string
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetDomainFields(domain Domain, onlyFields ...string) []string {
	tableName := domain.GetDBTableName()
	length := len(onlyFields)
	fields := make([]string, length)
	if len(onlyFields) > 0 {
		for i, field := range onlyFields {
			fields[i] = fmt.Sprintf("%s.%s", tableName, field)
		}

		return fields
	}

	value := reflect.Indirect(reflect.ValueOf(domain))
	length = value.Type().NumField()
	fields = make([]string, length)
	for i := 0; i < length; i++ {
		fields[i] = fmt.Sprintf("%s.%s", tableName, value.Type().Field(i).Tag.Get("json"))
	}

	return fields
}

func (u *User) GetDBTableName() string {
	return "users_user"
}

func main() {
	user := &User{}
	result := GetDomainFields(user, "id", "email")
	for _, r := range result {
		fmt.Println(r)
	}
}
