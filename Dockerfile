# syntax=docker/dockerfile:1

FROM golang:1.20

RUN go version
ENV GOPATH=/go
ENV GO111MODULE=on

COPY ./ /go/src/GoProjects/goservice-library
WORKDIR /go/src/GoProjects/goservice-library
RUN go mod download
#ENV GO111MODULE=off
RUN CGO_ENABLED=0 GOOS=linux go build -o goservice-library ./cmd/main.go

#RUN go mod download
#RUN go build -o goservice-library ./cmd/main/main.go

EXPOSE 8080

CMD ["./goservice-library"]
