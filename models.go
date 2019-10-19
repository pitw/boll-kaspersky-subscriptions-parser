package subscriptionparser

// Subscriptions provides meta data for subscriptions
type Subscriptions struct {
	Subscriber string
	Products   []Product
}

// Product provides meta data for products
type Product struct {
	Menge          int
	Artikel        Article
	EndkundenPreis float64
}

// Article provides meta data for each article
type Article struct {
	ArticleNr   string
	Description string
}
