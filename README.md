[![GitHub release](https://img.shields.io/github/tag/pitw/boll-kaspersky-subscriptions-parser.svg?label=latest&colorB=orange)](https://github.com/pitw/boll-kaspersky-subscriptions-parser/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/pitw/boll-kaspersky-subscriptions-parser)](https://goreportcard.com/report/github.com/pitw/boll-kaspersky-subscriptions-parser)
[![Build Status](https://travis-ci.org/pitwch/go-wrapper-proffix-restapi.svg?branch=master)](https://travis-ci.org/pitw/boll-kaspersky-subscriptions-parser)
[![GoDoc](https://godoc.org/github.com/pitw/boll-kaspersky-subscriptions-parser?status.svg)](https://godoc.org/github.com/pitw/boll-kaspersky-subscriptions-parser)
[![codecov](https://codecov.io/gh/pitw/boll-kaspersky-subscriptions-parser/branch/master/graph/badge.svg)](https://codecov.io/gh/pitw/boll-kaspersky-subscriptions-parser)
[![GitHub license](https://img.shields.io/github/license/pitw/boll-kaspersky-subscriptions-parser.svg)](https://github.com/pitw/boll-kaspersky-subscriptions-parser/blob/master/LICENSE)

# Boll Kaspersky Subscriptions Parser

Parsing Kaspersky Subscriptions directly from boll.ch


## Installation

```
go get https://github.com/pitw/boll-kaspersky-subscriptions-parser
```

## Example

```go
func main() {
	subscriptions, err := subscriptionparser.ParseSubscriptions("myuser", "mypsupersecretpassword")

	if (err != nil) {
		fmt.Print(err)
	}
  
  // Returns Subscription Struct
	fmt.Print(subscriptions)
}
```


## CLI

Also included is a small CLI for getting subscriptions
directly as:

- CSV
- JSON
- XLS


### Help

Can be showed with param -h

```
  -client int
        ID of Subscription - Client
  -filename string
        Name of exported file (default "KasperskySubscriptions")
  -format string
        Format (csv,json,xls) (default "csv")
  -password string
        Password for Boll.ch
  -username string
        Username for Boll.ch
  -version
        Shows version    
```
       