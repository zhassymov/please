# please - generic data validation library for Go

[![GoDoc](https://godoc.org/github.com/zhassymov/please?status.svg)](https://godoc.org/github.com/zhassymov/please)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhassymov/please)](https://goreportcard.com/report/github.com/zhassymov/please)
![License](https://img.shields.io/dub/l/vibe-d.svg)


## Installation
```bash
go get -u github.com/zhassymov/please
```

`please` is a library for data validation in the Go programming language, designed to simplify checking various conditions and constraints imposed on data.

## Goal
The goal of library is to provide convenient and flexible tools for validating data against specified criteria.
This can be useful for validating user input, working with APIs, data processing, and other scenarios that require checking and confirming the format or content of data.

## Features
- [x] compile-time type checking
- [x] no reflection
- [x] multiple, complex and configurable validation rules
- [x] custom validation rules
- [x] custom error messages
- [x] nested data validation


## Usage

### Simple Example
```go
package main

import (
    "fmt"
    "github.com/zhassymov/please"
)

func main() {
    password := "userPassword"
    err := please.Join(password,
        please.StringUTF8(),                            // validate the string is valid UTF-8
        please.StringLenBetween(16, 64),                // validate the string length is between 16 and 64 characters
        please.StringAlphaNumeric(),                    // validate the string contains only alphanumeric characters
        please.StringMinUniqueRuneCount(8),             // validate the string contains at least 8 unique characters
        please.StringContainsAny("123456789"),          // validate the string contains at least one number
        please.StringContainsAny(`!"#$%&'()*+,-./`),    // validate the string contains at least one of the special characters
    )
    if err != nil {
        fmt.Println(err)
    }
}
```