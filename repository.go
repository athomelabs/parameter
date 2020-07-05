package parameter

import (
	"database/sql"
)

type Repository interface {
	create(model *Parameter) error
	update(model *Parameter) error
	updateByName(name string, value string) error
	delete(id uint) error
	getByID(id uint) (*Parameter, error)
	getAll() (Parameters, error)
	getByName(name string) (*Parameter, error)
}

func NewRepository(engine string, db *sql.DB) Repository {
	switch engine {
	case "postgres":
		return newPsql(db)
	default:
		return nil
	}
}
