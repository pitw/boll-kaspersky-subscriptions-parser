# Boll Kaspersky Subscriptions Parser

Parsing Kaspersky Subscriptions directly from boll.ch


## Installation

```
go get https://github.com/pitw/boll-kaspersky-subscriptions-parser
```

## Example

```go
func main() {
	subscriptions, err := subscriptionparser.ParseSubscription("myuser", "mypsupersecretpassword")

	if (err != nil) {
		fmt.Print(err)
	}
  
  // Returns Subscription Struct
	fmt.Print(subscriptions)
}
```
