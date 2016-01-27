package application

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rebuy-de/golang-service-playground/database"
	"github.com/rebuy-de/golang-service-playground/http"
)

type Context struct {
	MysqlDsn   string
	HttpListen string

	db            *sql.DB
	fooRepository *database.FooRepository
}

func (c *Context) Run() {
	c.initMysql()
	c.createTables()
	c.listenHttp()
}

func (c *Context) initMysql() {
	var err error

	logrus.WithFields(logrus.Fields{
		"address": c.MysqlDsn,
	}).Info("Connecting to MySQL.")

	c.db, err = sql.Open("mysql", c.MysqlDsn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"conn":  c.MysqlDsn,
		}).Panic("Couldn't open database connection.")
	}

	c.fooRepository = database.NewFooRepository(c.db)
}

func (c *Context) createTables() {
	var err error

	_, err = c.db.Exec(`CREATE TABLE IF NOT EXISTS foo (
		id INT(11) PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		value VARCHAR(255) NOT NULL);`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Couldn't create database table.")
	}
}

func (c *Context) listenHttp() {
	http.Listen(c.HttpListen, c.fooRepository)
}
