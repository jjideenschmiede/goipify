# goipify

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goipify.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/goipify/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goipify/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goipify)](https://goreportcard.com/report/github.com/jjideenschmiede/goipify) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/goipify?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goipify) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/goipify) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Since we liked to access ipify's API for a small project, we developed the following library.

## Install

```console
go get github.com/jjideenschmiede/goipify
```

## How to use?

## Get your IP

With this function you get your current public IP address. You get the data back in JSON format. And you can access it directly via the Struct.

```go
// Get my ip address
ip, err := goipify.Ip()
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(ip)
}
```