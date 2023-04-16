package types

import (
	"database/sql/driver"
)

type AuthRole string

const (
	Admin AuthRole = "Admin"
	User  AuthRole = "User"
)

func (authRole *AuthRole) Scan(value interface{}) error {
	*authRole = AuthRole(value.(string))
	return nil
}

func (authRole AuthRole) Value() (driver.Value, error) {
	return string(authRole), nil
}

func (AuthRole) GormDataType() string {
	return "auth_role"
}
