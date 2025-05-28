# go-kebabcase

[![tag](https://img.shields.io/github/tag/cyg-pd/go-kebabcase.svg)](https://github.com/cyg-pd/go-kebabcase/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
[![GoDoc](https://godoc.org/github.com/cyg-pd/go-kebabcase?status.svg)](https://pkg.go.dev/github.com/cyg-pd/go-kebabcase)
![Build Status](https://github.com/cyg-pd/go-kebabcase/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/cyg-pd/go-kebabcase)](https://goreportcard.com/report/github.com/cyg-pd/go-kebabcase)
[![Coverage](https://img.shields.io/codecov/c/github/cyg-pd/go-kebabcase)](https://codecov.io/gh/cyg-pd/go-kebabcase)
[![Contributors](https://img.shields.io/github/contributors/cyg-pd/go-kebabcase)](https://github.com/cyg-pd/go-kebabcase/graphs/contributors)
[![License](https://img.shields.io/github/license/cyg-pd/go-kebabcase)](./LICENSE)

Fast snakecase implementation, believe it or not this was a large bottleneck in our application, Go's regexps are very slow.

## 🚀 Install

```sh
go get github.com/cyg-pd/go-kebabcase@v1
```

This library is v1 and follows SemVer strictly.

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

You can import `kebabcase` using:

```go
package main

import (
	"log/slog"

	"github.com/cyg-pd/go-kebabcase"
)

func main() {
	slog.Info(kebabcase.Kebabcase("HelloWorld")) // hello-world
}
```
