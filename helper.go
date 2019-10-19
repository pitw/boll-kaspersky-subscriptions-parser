package subscriptionparser

import (
	"fmt"
	"strings"
)

// stringCleaner cleans line breaks from strings
func stringCleaner(s string) (r string) {
	r = strings.Replace(s, "\n", "", -1)
	r = strings.Replace(r, "Trial period  ", "", -1)
	return r
}

// parseArticle parses Article to Nr and Description
func parseArticle(s string) (article Article, err error) {
	length := len(s)
	if length < 10 {
		return article, fmt.Errorf("No article")
	}
	return Article{ArticleNr: s[0:11], Description: s[11:length]}, err
}
