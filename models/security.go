package models

import (
	"database/sql"
)

type Security struct {
	Id   int
	Name string
}

type SecurityActivityEnum struct {
	SecurityActivityId       int
	Name                     string
	Description              sql.NullString
	FilterSecurityActivityId int
}

type SecurityActivitySection struct {
	Id          int
	Desctiption string
}
