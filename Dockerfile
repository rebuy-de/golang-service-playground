FROM golang:1.5

RUN go get \
	github.com/Sirupsen/logrus \
	github.com/codegangsta/cli \
	github.com/gin-gonic/gin \
	gopkg.in/bluesuncorp/validator.v5 \
	github.com/manucorporat/sse \
	github.com/mattn/go-colorable \
	golang.org/x/net/context \
	golang.org/x/net/context \
	github.com/go-sql-driver/mysql \
	github.com/jinzhu/gorm \
	github.com/jinzhu/inflection \
	github.com/lib/pq

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
