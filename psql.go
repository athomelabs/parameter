package parameter

import (
	"database/sql"
	"encoding/json"

	"github.com/athomelabs/dbcon"
	"github.com/lib/pq"
)

const (
	psqlInsert       = "INSERT INTO parameters (name, value) VALUES ($1, $2) RETURNING id, created_at"
	psqlUpdate       = "UPDATE parameters SET name = $1, value = $2, updated_at = now() WHERE id = $3"
	psqlDelete       = "DELETE FROM parameters WHERE id = $1"
	psqlGetAll       = "SELECT id, name, value, created_at, updated_at FROM parameters"
	psqlGetByID      = psqlGetAll + " WHERE id = $1"
	psqlGetByName    = psqlGetAll + " WHERE name = $1"
	psqlUpdateByName = "UPDATE parameters SET value = $1, updated_at = now() WHERE name = $2"
)

type Psql struct {
	db *sql.DB
}

func newPsql(db *sql.DB) *Psql {
	return &Psql{db: db}
}

func (p Psql) create(model *Parameter) error {
	stmt, err := p.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		model.Name,
		model.Value,
	).Scan(&model.ID, &model.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (p Psql) update(model *Parameter) error {
	stmt, err := p.db.Prepare(psqlUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = dbcon.ExecAffectingOneRow(stmt, model.Name, model.Value, model.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p Psql) updateByName(name string, value json.RawMessage) error {
	stmt, err := p.db.Prepare(psqlUpdateByName)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = dbcon.ExecAffectingOneRow(stmt, value, name)
	if err != nil {
		return err
	}

	return nil
}

func (p Psql) delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = dbcon.ExecAffectingOneRow(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (p Psql) getByID(id uint) (*Parameter, error) {
	var err error

	stmt, err := p.db.Prepare(psqlGetByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	return p.scanRow(row)
}

func (p Psql) getAll() (Parameters, error) {
	var err error
	var parameter Parameters

	stmt, err := p.db.Prepare(psqlGetAll)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		param, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}

		parameter = append(parameter, param)
	}

	return parameter, nil
}

func (p Psql) scanRow(s dbcon.RowScanner) (*Parameter, error) {
	r := &Parameter{}
	nt := pq.NullTime{}

	err := s.Scan(&r.ID,
		&r.Name,
		&r.Value,
		&r.CreatedAt,
		&nt)
	if err != nil {
		return r, err
	}

	r.UpdatedAt = nt.Time

	return r, nil
}

func (p Psql) getByName(name string) (*Parameter, error) {
	stmt, err := p.db.Prepare(psqlGetByName)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(name)

	return p.scanRow(row)
}
