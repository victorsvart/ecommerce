package utils

import (
	"errors"

	"gorm.io/gorm"
)

func CheckNotFoundErr(err error, entityNotFound error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entityNotFound
	}

	return err
}
