package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func mustInitMysql(addr string) gorm.DB {
	logrus.WithFields(logrus.Fields{
		"address": addr,
	}).Info("Connecting to MySQL.")

	var db, err = gorm.Open("mysql", addr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"conn":  addr,
		}).Error("Couldn't open database connection.")
		os.Exit(1)
	}

	if !db.HasTable(&Entry{}) {
		err = db.CreateTable(&Entry{}).Error
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err,
				"type":  "Entry",
			}).Error("Couldn't create database table connection.")
			os.Exit(1)
		}
	}

	return db
}
