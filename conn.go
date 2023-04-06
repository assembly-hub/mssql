package mssql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

type Config struct {
	Username        string
	Password        string
	DBName          string
	Server          string
	Port            int
	Params          string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
	ConnMaxIdleTime int
}

type Client struct {
	cfg *Config
}

func NewClient(cfg *Config) *Client {
	c := new(Client)
	c.cfg = cfg
	return c
}

func (c *Client) Connect() (*sql.DB, error) {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		c.cfg.Username, c.cfg.Password, c.cfg.Server, c.cfg.Port, c.cfg.DBName)
	if c.cfg.Params != "" {
		connString += "&" + c.cfg.Params
	}
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Duration(c.cfg.ConnMaxLifeTime) * time.Millisecond)
	db.SetConnMaxIdleTime(time.Duration(c.cfg.ConnMaxIdleTime) * time.Millisecond)
	db.SetMaxOpenConns(c.cfg.MaxOpenConn)
	db.SetMaxIdleConns(c.cfg.MaxIdleConn)
	return db, err
}
