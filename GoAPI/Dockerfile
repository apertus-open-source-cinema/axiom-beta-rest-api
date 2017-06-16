FROM golang:1.8

RUN mkdir -p /go/src/client-app

WORKDIR /go/src/client-app

COPY . /go/src/client-app

RUN go-wrapper download
RUN go-wrapper install

ENV PORT 8080

EXPOSE 8080

CMD ["client", "run"]
