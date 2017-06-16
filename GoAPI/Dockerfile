FROM iron/go:dev

WORKDIR /app


ENV SRC_DIR=/go/src/github.com/apertus-open-source-cinema/axiom-beta-rest-interface/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o client; cp client /app/
ENTRYPOINT ["./client"]
