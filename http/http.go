package http

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/rebuy-de/golang-service-playground/types"
)

type FooRepository interface {
	FindById(id int) (*types.Foo, error)
	Create(*types.Foo) error
}

type controller struct {
	foo FooRepository
}

func Listen(address string, foo FooRepository) {
	var ctl = controller{
		foo,
	}

	r := gin.Default()
	r.GET("/foo/:id", ctl.getFooById)
	r.POST("/foo", ctl.postFoo)

	var err = r.Run(address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panic("Failed to serve HTTP.")
	}
}

func (ctl *controller) getFooById(g *gin.Context) {
	var err error
	var foo = new(types.Foo)
	var id int

	id, err = strconv.Atoi(g.Param("id"))
	if err != nil {
		g.String(400, err.Error())
	}

	foo, err = ctl.foo.FindById(id)
	if err != nil {
		g.String(404, err.Error())
	}

	g.JSON(200, foo)
}

func (ctl *controller) postFoo(g *gin.Context) {
	var err error
	var foo = new(types.Foo)

	err = g.Bind(foo)
	if err != nil {
		g.String(500, err.Error())
	}

	err = ctl.foo.Create(foo)
	if err != nil {
		g.String(400, err.Error())
	}

	g.Header("Location", fmt.Sprintf("/foo/%d", foo.ID))
	g.String(201, "")
}
