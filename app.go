package main

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/go-mysql-org/go-mysql/driver"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Connect(connectionType string, dsnRaw string) error {
	d, err := sqlx.Connect(connectionType, dsnRaw)
	if err != nil {
		return err
	}

	db = d
	return nil
}

type resultRun struct {
	Columns []string
	Rows    []map[string]any
}

func (a *App) Run(query string) (*resultRun, error) {
	if query == "" {
		return nil, errors.New("can't execute empty query")
	}

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	for rows.Next() {
		r := make(map[string]any)
		err := rows.MapScan(r)
		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}

	return &resultRun{columns, result}, nil
}
