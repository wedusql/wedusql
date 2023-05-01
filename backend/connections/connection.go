package connections

import (
	"fmt"
	"net/url"
	"strings"

	_ "github.com/go-mysql-org/go-mysql/driver"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var C *Connection

type Connection struct {
	// _type is a connection type, e.g. mysql, pgx, etc.
	_type string
	dsn   string
	db    *sqlx.DB
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(connectionType string, dsnRaw string) error {

	if connectionType == "" {
		return nil
	}

	c.dsn = dsnRaw

	if connectionType == "mysql" || connectionType == "pgx" {
		d, err := sqlx.Connect(connectionType, dsnRaw)
		if err != nil {
			return err
		}
		c.db = d
		c._type = connectionType
		C = c
	}

	return nil
}

func (c *Connection) Disconnect() {
	C = nil
	c.db = nil
	c._type = ""
}

func (c *Connection) TestConnection(connectionType string, dsnRaw string) (bool, error) {
	if connectionType == "" {
		return false, nil
	} else if connectionType == "mysql" || connectionType == "pgx" {
		d, err := sqlx.Connect(connectionType, dsnRaw)
		if err != nil {
			return false, err
		}
		defer d.Close()

		err = d.Ping()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func (c *Connection) Close() error {
	if c._type == "mysql" || c._type == "pgx" {
		return c.db.Close()
	}
	return nil
}

func (c *Connection) GetConnectionType() string {
	return c._type
}

func (c *Connection) GetConnection() any {
	if c._type == "mysql" || c._type == "pgx" {
		return c.db
	}

	return nil
}

type connectionInfo struct {
	Driver   string
	User     string
	Host     string
	Database string
}

func (c *Connection) GetConnectionInfo() (*connectionInfo, error) {
	rawDsn := c.dsn
	if c._type == "mysql" {
		rawDsn = "mysql://" + rawDsn
	}

	dsn, err := url.Parse(rawDsn)
	if err != nil {
		return nil, err
	}

	fmt.Println("mantap")

	ci := &connectionInfo{
		Driver:   c._type,
		User:     dsn.User.Username(),
		Host:     dsn.Host,
		Database: strings.ReplaceAll(dsn.Path, "/", ""),
	}

	return ci, nil
}
