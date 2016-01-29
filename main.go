package main

import (
	"flag"

	"github.com/rebuy-de/golang-service-playground/application"
)

type Config struct {
	HttpListen string
	MysqlDsn   string
}

func main() {
	var config Config

	flag.StringVar(&config.HttpListen, "http-listen", ":8080",
		"HTTP server listen interface and port.")
	flag.StringVar(&config.MysqlDsn, "mysql-dsn", "localhost",
		"[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]")
	flag.Parse()

	var app = application.Context{
		HttpListen: config.HttpListen,
		MysqlDsn:   config.MysqlDsn,
	}

	app.Run()
}
