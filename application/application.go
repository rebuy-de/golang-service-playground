package application

import (
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"github.com/rebuy-de/golang-service-playground/http"
	"github.com/rebuy-de/golang-service-playground/types"
)

type Context struct {
	MysqlDsn   string
	HttpListen string

	gorm gorm.DB
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

	c.gorm, err = gorm.Open("mysql", c.MysqlDsn)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"conn":  c.MysqlDsn,
		}).Panic("Couldn't open database connection.")
	}
}

func (c *Context) createTables() {
	if !c.gorm.HasTable(&types.Entry{}) {
		var err = c.gorm.CreateTable(&types.Entry{}).Error
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
				"type":  "Entry",
			}).Panic("Couldn't create database table connection.")
		}
	}
}

func (c *Context) listenHttp() {
	http.Listen(c.HttpListen, c.gorm)
}
