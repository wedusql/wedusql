package queries

import (
	"errors"
	"wedusql/backend/connections"

	"github.com/jmoiron/sqlx"
)

type resultRun struct {
	Columns []string
	Rows    []map[string]any
}

func (q *Query) runSqlxQuery(query string) (*resultRun, error) {

	if query == "" {
		return nil, errors.New("can't execute empty query")
	}

	db := connections.C.GetConnection().(*sqlx.DB)

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

		for k, v := range r {
			switch v.(type) {
			case []uint8:
				r[k] = string(v.([]byte))
			}
		}

		result = append(result, r)
	}

	return &resultRun{columns, result}, nil

}
