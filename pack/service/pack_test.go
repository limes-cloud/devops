package service

import "testing"

func TestPack(t *testing.T) {
	pk := NewPack()
	pk.RegistryUrl = "registry.com"
	pk.WorkDir = "/Users/apple/docker"
	pk.GitUrl = "https://gitee.com/limeschool/hello.git"
	pk.RegistryUser = "root"
	pk.RegistryPass = "xrxy0852"
	pk.ServerName = "helloworld"
	pk.ServerBranch = "origin/master"
	pk.ServerVersion = "qwer"
	pk.Exec = "/bin/sh"
	pk.Dockerfile = `FROM golang:alpine AS build
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/build
ADD . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o entry main.go

FROM alpine
EXPOSE 9001
WORKDIR /go/build
COPY --from=build /go/build/entry /go/build/entry
CMD ["./entry"]`
	pk.Args = []string{}

	if err := pk.Start(); err != nil {
		t.Fatal(err)
	}
	t.Log("pack success")
}
