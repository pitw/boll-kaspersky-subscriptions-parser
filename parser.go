package subscriptionparser

import (
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

// ParseSubscriptions parses all Boll Kaspersky Subscriptions
func ParseSubscriptions(username string, password string) (subscriptions []Subscriptions, err error) {
	return parseSubscriptionPage(username, password, previewURL)
}

// ParseSubscriptionClient parses a specific Boll Kaspersky Client Subscription
func ParseSubscriptionClient(username string, password string, id int) (subscriptions []Subscriptions, err error) {
	return parseSubscriptionPage(username, password, previewURLClient+strconv.Itoa(id))
}

// parseSubscriptionPage parses Boll Kaspersky Subscription Tables
func parseSubscriptionPage(username string, password string, url string) (subscriptions []Subscriptions, err error) {
	c := colly.NewCollector()

	// authenticate
	err = c.Post(loginURL, map[string]string{"_profile_username": username, "_noasst_account_password": password, "_do_login": "Anmelden"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {

	})
	var subscriber string
	newBill := &Subscriptions{}
	newProduct := &Product{}
	count := 0
	c.OnHTML("body > table > tbody", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(p int, el *colly.HTMLElement) {

			if (el.Attr("bgcolor")) != "" {
				subscriber = el.ChildText("td:nth-child(2)")
				newBill = &Subscriptions{Subscriber: subscriber}
				subscriptions = append(subscriptions, *newBill)
				count++
			} else {
				menge, noLine := strconv.Atoi(stringCleaner(el.ChildText("td:nth-child(1)")))
				artikelname, noLine := parseArticle(el.ChildText("td:nth-child(2)"))
				preisEndkunde, noLine := strconv.ParseFloat(stringCleaner(el.ChildText("td:nth-child(3)")), 64)

				if noLine == nil {
					newProduct = &Product{Artikel: artikelname, Menge: menge, EndkundenPreis: preisEndkunde}
					subscriptions[count-1].Products = append(newBill.Products, *newProduct)

				}

			}

		})

	})

	// start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	// logout
	_ = c.Visit(logoutURL)

	return subscriptions, err
}
