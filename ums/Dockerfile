FROM golang:alpine AS build
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
EXPOSE 8080
WORKDIR /go/build
COPY ./config /go/build/config
COPY ./cert /go/build/cert
COPY --from=build /go/build/entry /go/build/entry
CMD ["./entry","-c","./config/conf.json"]
