package service

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPack(t *testing.T) {
	pk := NewPack()
	pk.RegistryUrl = "registry.com"
	pk.WorkDir = "/Users/apple/docker"
	pk.GitUrl = "https://gitee.com/ptl-f/ps-go.git"
	pk.GitUser = "ptl-f"
	pk.GitPass = "xrxy0852"
	pk.RegistryUser = "root"
	pk.RegistryPass = "xrxy0852"
	pk.ServerName = "ps-go"
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
COPY --from=build /go/build/entry /go/build/{entry}
CMD ["./entry"]`
	pk.Args = map[string]string{"entry": "entry"}

	pk.SetWatch(func(s string) {
		s = strings.ReplaceAll(s, "\n", "")
		fmt.Println(s + "  -" + time.Now().Format("2006-01-02 15:04:05"))
	})

	if err := pk.Start(); err != nil {
		t.Fatal(err)
	}

	t.Log("pack success")
}
