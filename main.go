package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/rebuy-de/golang-service-playground/application"
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

	var app = application.Context{
		MysqlDsn:   c.String("mysql-conn"),
		HttpListen: c.String("http-listen"),
	}

	app.Run()
}
