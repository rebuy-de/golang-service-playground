package main

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Entry struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

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

func mustDialMysql(addr string) *sql.DB {
	logrus.WithFields(logrus.Fields{
		"address": addr,
	}).Info("Connection to MySQL.")

	var con, err = sql.Open("mysql", addr)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
			"conn":  addr,
		}).Error("Couldn't open database connection.")
		os.Exit(1)
	}

	return con
}

func mustCreateTables(con *sql.DB) {
	var err error

	_, err = con.Exec(`CREATE TABLE IF NOT EXISTS mytable (
		id INT(11) PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		value VARCHAR(255) NOT NULL);`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Couldn't create database table.")
		os.Exit(1)
	}

}

func listenHttp(address string, db *sql.DB) error {
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
		var id string

		id = c.Param("id")

		err = db.QueryRow("SELECT id, name, value FROM mytable WHERE id=?", id).
			Scan(&entry.ID, &entry.Name, &entry.Value)
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

		var result sql.Result
		result, err = db.Exec(`INSERT INTO mytable (name, value) VALUES (?, ?)`, entry.Name, entry.Value)
		if err != nil {
			c.String(500, err.Error())
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			c.String(500, err.Error())
		}

		c.Header("Location", fmt.Sprintf("/foo/%d", id))
		c.String(201, "")
	})

	return r.Run(address)
}
