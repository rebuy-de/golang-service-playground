package main

import (
	"os"

	"database/sql"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	var err = listenHttp(c.String("http-listen"))
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

/*func createRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/foo/:name", fooGet)
	router.PUT("/foo/:name", fooPut)
	router.DELETE("/foo/:name", fooDelete)
	router.POST("/foo", fooPost)

	return router
}*/

func listenHttp(address string) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r.Run(address)
}
