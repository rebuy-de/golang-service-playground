package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func listenHttp(address string, db gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
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
