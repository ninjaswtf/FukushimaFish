package main

import (
	"database/sql/driver"
	"strings"
)

type StringArray []string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (s *StringArray) Scan(value interface{}) error {
	str, ok := value.(string)

	if ok {
		sa := StringArray(strings.Split(str, ","))
		s = &sa
	}

	return nil
}

// Value return json value, implement driver.Valuer interface
func (s StringArray) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "", nil
	}
	return strings.Join(s, ","), nil
}
