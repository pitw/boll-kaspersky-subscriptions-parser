package main

import (
	"flag"
	"fmt"
	"github.com/pitw/boll-kaspersky-subscriptions-parser"
)

func main() {
	var filename string
	var username string
	var password string
	var clientid int
	flag.StringVar(&filename, "filename", "KasperksySubscriptions", "Name of exported CSV (without ending)")
	flag.StringVar(&username, "username", "", "Username for Boll.ch")
	flag.StringVar(&password, "password", "", "Password for Boll.ch")
	flag.IntVar(&clientid, "client", 0, "ID of Client")

	flag.Parse()

	if password == "" {
		fmt.Println("Password is mandatory")
	} else if username == "" {
		fmt.Println("Username is mandatory")
	} else {
		if clientid > 0 {
			WriteSubscriptionClientCSV(username, password, filename, clientid)
		} else {
			err := WriteSubscriptionsCSV(username, password, filename)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

// WriteSubscriptionsCSV parses a all Boll Kaspersky Subscription to CSV
func WriteSubscriptionsCSV(username string, password string, filename string) (err error) {
	subsc, err := subscriptionparser.ParseSubscriptions(username, password)
	if err != nil {
		return err
	}
	err = subscriptionparser.WriteCSV(subsc, filename)
	if err != nil {
		return err
	} else {
		fmt.Printf("CSV written to %v", filename)
	}
	return err
}

// WriteSubscriptionCSV parses a all Boll Kaspersky Subscription to CSV
func WriteSubscriptionClientCSV(username string, password string, filename string, clientid int) {
	subsc, err := subscriptionparser.ParseSubscriptionClient(username, password, clientid)
	if err != nil {
		fmt.Println(err)
	}
	err = subscriptionparser.WriteCSV(subsc, filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("CSV written to %v", filename)
	}
}
