package types

import (
	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"
)

type HashedString []byte

func (hashedString *HashedString) Scan(value interface{}) error {
	*hashedString = HashedString(value.([]byte))
	return nil
}

func (hashedString HashedString) Value() (driver.Value, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(hashedString, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

func (HashedString) GormDataType() string {
	return "bytea"
}
