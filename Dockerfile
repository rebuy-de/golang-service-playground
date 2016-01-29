FROM golang:1.5

RUN go get \
	github.com/Sirupsen/logrus \
	github.com/go-sql-driver/mysql \
	github.com/julienschmidt/httprouter

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
