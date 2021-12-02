package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

const (
	DnsTemplate = "%s:%s@tcp(%s:%d)/%s?multiStatements=true&parseTime=true"
	DefaultHost = "localhost"
	DefaultPort = 3306
)

type Client struct {
	*gorm.DB
}

type Config struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

func New(db *gorm.DB) *Client {
	return &Client{db}
}

func (cf *Config) Dns() string {
	var (
		host               = DefaultHost
		port               = DefaultPort
		username, password string
	)
	if len(cf.Host) > 0 {
		host = cf.Host
	}
	if cf.Port > 0 {
		port = cf.Port
	}
	if len(cf.User) > 0 {
		username = cf.User
	}
	if len(cf.Password) > 0 {
		password = cf.Password
	}
	return fmt.Sprintf(DnsTemplate, username, password, host, port, cf.DatabaseName)
}

func Open(conf Config) (*Client, error) {
	session, err := gorm.Open(mysql.Open(conf.Dns()), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "db-gorm : failed initialize db session.")
	}
	db, err := session.DB()
	if err != nil {
		return nil, errors.Wrap(err, "db-gorm : failed open connection.")
	}
	if conf.MaxOpenConns > 0 {
		db.SetMaxIdleConns(conf.MaxOpenConns)
	}
	if conf.MaxIdleConns > 0 {
		db.SetMaxIdleConns(conf.MaxIdleConns)
	}
	return &Client{session}, nil
}

func (c *Client) Close() error {
	db, err := c.DB.DB()
	if err != nil {
		return errors.Wrap(err, "db : failed to get db")
	}
	return db.Close()
}
