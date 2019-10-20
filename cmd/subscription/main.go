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
	var version bool
	var format string

	var clientid int

	flag.StringVar(&format, "format", "csv", "Format of file (csv,json)")
	flag.StringVar(&filename, "filename", "KasperskySubscriptions", "Name of exported file")
	flag.StringVar(&username, "username", "", "Username for Boll.ch")
	flag.StringVar(&password, "password", "", "Password for Boll.ch")
	flag.BoolVar(&version, "version", false, "Shows version")
	flag.IntVar(&clientid, "client", 0, "ID of Subscription - Client")

	// Parse all flags
	flag.Parse()

	if password == "" {
		fmt.Println("Password is mandatory")
	} else if username == "" {
		fmt.Println("Username is mandatory")
	} else {
		if clientid > 0 {
			switch format {
			case "json":
				WriteSubscriptionClientJSON(username, password, filename, clientid)
			case "csv":
				WriteSubscriptionClientCSV(username, password, filename, clientid)
			default:
				WriteSubscriptionClientCSV(username, password, filename, clientid)
			}
		} else {

			switch format {
			case "json":
				WriteSubscriptionsJSON(username, password, filename)
			case "csv":
				WriteSubscriptionsCSV(username, password, filename)
			default:
				WriteSubscriptionsCSV(username, password, filename)
			}

		}
	}

}

// WriteSubscriptionsCSV parses a all Boll Kaspersky Subscription to CSV
func WriteSubscriptionsCSV(username string, password string, filename string) {
	subsc, err := subscriptionparser.ParseSubscriptions(username, password)
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

// WriteSubscriptionsJSON parses a all Boll Kaspersky Subscription to CSV
func WriteSubscriptionsJSON(username string, password string, filename string) {
	subsc, err := subscriptionparser.ParseSubscriptions(username, password)
	if err != nil {
		fmt.Println(err)
	}
	err = subscriptionparser.WriteJSON(subsc, filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("JSON written to %v", filename)
	}
}

// WriteSubscriptionClientJSON parses a all Boll Kaspersky Subscription to CSV
func WriteSubscriptionClientJSON(username string, password string, filename string, clientid int) {
	subsc, err := subscriptionparser.ParseSubscriptionClient(username, password, clientid)
	if err != nil {
		fmt.Println(err)
	}
	err = subscriptionparser.WriteJSON(subsc, filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("JSON written to %v", filename)
	}
}
