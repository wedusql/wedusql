package queries

import (
	"errors"
	"wedusql/backend/connections"
)

type Query struct {
	connection connections.Connection
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) SaveConnection(connType, dsn string) error {
	return q.connection.Connect(connType, dsn)
}

func (q *Query) TestConnection(connType, dsn string) (bool, error) {
	return q.connection.TestConnection(connType, dsn)
}

func (q *Query) Run(query string) (*resultRun, error) {

	connType := connections.C.GetConnectionType()
	if connType == "mysql" || connType == "pgx" {
		return q.runSqlxQuery(query)
	}

	return nil, errors.New("can't execute empty connection")
}
