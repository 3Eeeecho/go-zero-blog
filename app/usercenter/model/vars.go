package model

import (
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrInvalidParams = "invalid params"
