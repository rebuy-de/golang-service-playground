package database

import (
	"database/sql"

	"github.com/rebuy-de/golang-service-playground/types"
)

type FooRepository struct {
	db *sql.DB
}

func NewFooRepository(db *sql.DB) *FooRepository {
	return &FooRepository{db}
}

func (r *FooRepository) FindById(id int) (*types.Foo, error) {
	var query = `SELECT id, name, value FROM foo WHERE id = ?;`
	var row = r.db.QueryRow(query, id)
	var foo = new(types.Foo)
	var err = row.Scan(&foo.ID, &foo.Name, &foo.Value)
	if err != nil {
		return nil, err
	}

	return foo, nil
}

func (r *FooRepository) Create(foo *types.Foo) error {
	var err error
	var result sql.Result

	var query = `INSERT INTO foo (name, value) VALUES (?, ?);`
	result, err = r.db.Exec(query, foo.Name, foo.Value)
	if err != nil {
		return err
	}

	foo.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
