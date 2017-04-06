package repos

import (
	"database/sql"
)

type Base struct {
	DB *sql.DB
}

func NewBaseRepo(db *sql.DB) *Base {
	base := new(Base)
	base.DB = db

	return base
}
