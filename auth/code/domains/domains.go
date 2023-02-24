package domains

import (
	"fmt"
	"reflect"
)

type Domain interface {
	GetDBTableName() string
	GetDBValues() []any
}

func GetDomainFields(domain Domain) []string {
	tableName := domain.GetDBTableName()
	value := reflect.Indirect(reflect.ValueOf(domain))

	length := value.Type().NumField()
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = fmt.Sprintf("%s.%s", tableName, value.Type().Field(i).Tag.Get("db"))
	}

	return result
}
