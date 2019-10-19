package subscriptionparser

type Subscriptions struct {
	Subscriber string
	Products   []Product
}

type Product struct {
	Menge          int
	Artikel        Article
	EndkundenPreis float64
}

type Article struct {
	ArticleNr   string
	Description string
}
