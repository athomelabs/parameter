package parameter

import (
	"database/sql"
	"encoding/json"
)

type Repository interface {
	create(model *Parameter) error
	update(model *Parameter) error
	updateByName(name string, value json.RawMessage) error
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
