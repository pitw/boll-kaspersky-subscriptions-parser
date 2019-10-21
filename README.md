[![GitHub release](https://img.shields.io/github/tag/pitw/boll-kaspersky-subscriptions-parser.svg?label=latest&colorB=orange)](https://github.com/pitw/boll-kaspersky-subscriptions-parser/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/pitw/boll-kaspersky-subscriptions-parser)](https://goreportcard.com/report/github.com/pitw/boll-kaspersky-subscriptions-parser)
[![Build Status](https://travis-ci.org/pitwch/go-wrapper-proffix-restapi.svg?branch=master)](https://travis-ci.org/pitw/boll-kaspersky-subscriptions-parser)
[![GoDoc](https://godoc.org/github.com/pitw/boll-kaspersky-subscriptions-parser?status.svg)](https://godoc.org/github.com/pitw/boll-kaspersky-subscriptions-parser)
[![GitHub license](https://img.shields.io/github/license/pitw/boll-kaspersky-subscriptions-parser.svg)](https://github.com/pitw/boll-kaspersky-subscriptions-parser/blob/master/LICENSE)

# Boll Kaspersky Subscriptions Parser

Golang Library for parsing Kaspersky Subscriptions directly from [Boll.ch](https://www.boll.ch/kaspersky/subscriptions_faq.html)

- [Installation](#installation)
- [Example](#example)
- [CLI :computer: ](#cli)
- [Docker :whale2:](#docker)

## Installation

```
go get https://github.com/pitw/boll-kaspersky-subscriptions-parser
```

### Example

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
Detailed docs available in [GoDoc](https://godoc.org/github.com/pitw/boll-kaspersky-subscriptions-parser)

## CLI

Also included is a small :computer:  CLI for getting subscriptions
directly as:

- CSV
- JSON

Latest version of CLI is always available in [Releases](https://github.com/pitw/boll-kaspersky-subscriptions-parser/releases/latest)

Example:

```
subscription.exe -password 1234 -username test

// Parses Subscriptions and downloads them as CSV
```



### Help

Can be viewed with param -h

```
subscription.exe -h


  -client int
        ID of Subscription - Client
  -filename string
        Name of exported file (default "KasperskySubscriptions")
  -format string
        Format (csv,json) (default "csv")
  -password string
        Password for Boll.ch
  -username string
        Username for Boll.ch
  -version
        Shows version    
```

## Docker

CLI is also available as :whale2: [Docker Image](https://hub.docker.com/r/pitwch/boll-kaspersky-subscriptions).

```
docker pull pitwch/boll-kaspersky-subscriptions:latest
```

