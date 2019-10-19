package subscriptionparser

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func WriteCSV(subscriptions []Subscriptions, filename string) (err error) {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, record := range subscriptions {
		line := []string{}
		for _, products := range record.Products {
			line = []string{
				record.Subscriber,
				strconv.Itoa(products.Menge),
				products.Artikel.ArticleNr,
				products.Artikel.Description,
				fmt.Sprintf("%f", products.EndkundenPreis)}
		}
		if err := w.Write(line); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	return err
}
