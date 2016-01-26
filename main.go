package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
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

	var con = mustInitMysql(c.String("mysql-conn"))
	defer con.Close()

	var err = listenHttp(c.String("http-listen"), con)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to server HTTP.")
		os.Exit(1)
	}
}
