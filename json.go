package subscriptionparser

import (
	"encoding/json"
	"log"
	"os"
)

func WriteJSON(subscriptions []Subscriptions, filename string) (err error) {
	file, err := os.Create(filename + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	w := json.NewEncoder(file)

	if err := w.Encode(subscriptions); err != nil {
		log.Fatalln("error writing record to json:", err)
	}

	return err
}
