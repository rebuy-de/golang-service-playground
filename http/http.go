package http

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rebuy-de/golang-service-playground/types"
)

func Listen(address string, db gorm.DB) {
	var ctx = context{
		&db,
	}

	r := gin.Default()
	r.GET("/foo/:id", ctx.getFooById)
	r.POST("/foo", ctx.postFoo)

	var err = r.Run(address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Failed to serve HTTP.")
	}
}

type context struct {
	db *gorm.DB
}

func (ctx *context) getFooById(g *gin.Context) {
	var err error
	var entry = new(types.Entry)

	entry.ID, err = strconv.Atoi(g.Param("id"))
	if err != nil {
		g.String(500, err.Error())
	}

	err = ctx.db.First(entry).Error
	if err != nil {
		g.String(500, err.Error())
	}

	g.JSON(200, entry)
}

func (ctx *context) postFoo(g *gin.Context) {
	var err error
	var entry = new(types.Entry)

	err = g.Bind(entry)
	if err != nil {
		g.String(500, err.Error())
	}

	err = ctx.db.Create(entry).Error
	if err != nil {
		g.String(500, err.Error())
	}

	g.Header("Location", fmt.Sprintf("/foo/%d", entry.ID))
	g.String(201, "")
}
