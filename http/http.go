package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

	var router = httprouter.New()
	router.GET("/foo/:id", ctl.getFooById)
	router.POST("/foo", ctl.postFoo)

	log.Panic(http.ListenAndServe(address, router))
}

func (ctl *controller) getFooById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var foo = new(types.Foo)
	var id int

	id, err = strconv.Atoi(p.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err.Error())
	}

	foo, err = ctl.foo.FindById(id)
	if err != nil {
		w.WriteHeader(404)
		io.WriteString(w, err.Error())
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(foo)
}

func (ctl *controller) postFoo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var err error
	var foo = new(types.Foo)

	err = json.NewDecoder(r.Body).Decode(foo)
	if err != nil {
		w.WriteHeader(500)
		io.WriteString(w, err.Error())
	}

	err = ctl.foo.Create(foo)
	if err != nil {
		w.WriteHeader(400)
		io.WriteString(w, err.Error())
	}

	w.Header().Set("Location", fmt.Sprintf("/foo/%d", foo.ID))
	w.WriteHeader(200)
}
