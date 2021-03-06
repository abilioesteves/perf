# BUILD
FROM golang:1.11-alpine as builder

RUN apk add --no-cache git mercurial 

ENV BUILD_PATH=$GOPATH/src/github.com/abilioesteves/perf/src

RUN mkdir -p ${BUILD_PATH}
WORKDIR ${BUILD_PATH}

ADD ./src ./
RUN go get -v ./...

WORKDIR ${BUILD_PATH}/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /perf .

# PKG
FROM ubuntu:latest

ENV WRITE_PATH ""

COPY --from=builder /perf /

ADD startup.sh /

EXPOSE 17333

USER root

CMD ["/startup.sh"]
