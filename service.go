package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	app := cli.NewApp()
	app.Name = "example service"
	app.Usage = "example golang http service"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "http-listen",
			Value:  ":8080",
			Usage:  "HTTP server listen interface and port.",
			EnvVar: "HTTP_LISTEN",
		},
		cli.StringFlag{
			Name:   "mysql-conn",
			Value:  "localhost",
			Usage:  "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]",
			EnvVar: "MYSQL_CONN",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	logrus.SetLevel(logrus.DebugLevel)

	var con = mustDialMysql(c.String("mysql-conn"))
	defer con.Close()
	mustCreateTables(con)

	var err = listenHttp(c.String("http-listen"), con)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to server HTTP.")
		os.Exit(1)
	}
}

func mustDialMysql(addr string) gorm.DB {
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

	return db
}

func mustCreateTables(db gorm.DB) {
	db.CreateTable(&Entry{})
}

func listenHttp(address string, db gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, Entry{1, "foo", "bar"})
	})

	r.GET("/foo/:id", func(c *gin.Context) {
		var err error
		var entry = new(Entry)

		entry.ID, err = strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(500, err.Error())
		}

		err = db.First(entry).Error
		if err != nil {
			c.String(500, err.Error())
		}

		c.JSON(200, entry)
	})

	r.POST("/foo", func(c *gin.Context) {
		var err error
		var entry = new(Entry)

		err = c.Bind(entry)
		if err != nil {
			c.String(500, err.Error())
		}

		err = db.Create(entry).Error
		if err != nil {
			c.String(500, err.Error())
		}

		c.Header("Location", fmt.Sprintf("/foo/%d", entry.ID))
		c.String(201, "")
	})

	return r.Run(address)
}
