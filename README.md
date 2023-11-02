# Telegram Login Widget

:star: Star us on GitHub — it motivates us a lot!

Telegram Login Widget library for Go/Golang.

## Table of Content

- [Usage/Examples](#usageexamples)
  - [Unmarshal JSON](#unmarshal-json)
  - [NewFromQuery](#newfromquery)
  - [NewFromURI](#newfromuri)
- [FAQ](#faq)
- [License](#license)
- [Links](#links)

## Usage/Examples

#### Unmarshal JSON

```go
package main

import (
	"encoding/json"
	
	"github.com/LipsarHQ/go-telegram-login-widget"
)

// Telegram bot token.
const token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	// 1. Get "AuthorizationData" from JSON.
	data := []byte(`{"first_name":"Klim","hash":"b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95","last_name":"Sidorov","photo_url":"https://t.me/klimsidorov","username":"klimsidorov","auth_date":976255200,"id":1}`)
	
	var modelAuthorizationData telegramloginwidget.AuthorizationData
	
	if err := json.Unmarshal(data, &modelAuthorizationData); err != nil {
		return
	}

	// 2. Call "Check" method to validate hash.
	if err := modelAuthorizationData.Check(token); err != nil {
		// Invalid hash.
		return
	}
	
	// Hash is valid.
}
```

#### NewFromQuery

```go
package main

import (
	"net/url"

	"github.com/LipsarHQ/go-telegram-login-widget"
)

// Telegram bot token.
const token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	// 1. Get "AuthorizationData" from url.Values.
	values := url.Values{
		"auth_date":  []string{"976255200"},
		"first_name": []string{"Klim"},
		"hash":       []string{"b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95"},
		"id":         []string{"1"},
		"last_name":  []string{"Sidorov"},
		"photo_url":  []string{"https://t.me/klimsidorov"},
		"username":   []string{"klimsidorov"},
	}
	
	modelAuthorizationData, err := telegramloginwidget.NewFromQuery(values)
	if err != nil {
		return
	}

	// 2. Call "Check" method to validate hash.
	if err = modelAuthorizationData.Check(token); err != nil {
		// Invalid hash.
		return
	}

	// Hash is valid.
}
```

#### NewFromURI

```go
package main

import (
	"github.com/LipsarHQ/go-telegram-login-widget"
)

// Telegram bot token.
const token = "XXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	// 1. Get "AuthorizationData" from uri.
	const uri = "https://example.com/?auth_date=976255200&first_name=Klim&hash=b7a7fc776729077786e4190aec2c5dcecd2ec66ae0faf1b44316d541b955da95&id=1&last_name=Sidorov&photo_url=https%3A%2F%2Ft.me%2Fklimsidorov&username=klimsidorov"
	
	modelAuthorizationData, err := telegramloginwidget.NewFromURI(uri)
	if err != nil {
		return
	}

	// 2. Call "Check" method to validate hash.
	if err = modelAuthorizationData.Check(token); err != nil {
		// Invalid hash.
		return
	}

	// Hash is valid.
}
```

## FAQ

#### What is Telegram Login Widget?

The Telegram login widget is a simple way to authorize users on your website.
Check out [this page](https://core.telegram.org/widgets/login) for a general overview of the widget.

#### How to validate hash?

Call `Check` method on `AuthorizationData` struct to validate hash.

#### How to calculate hash (for debug purpose)?

Call `Sum` method on `AuthorizationData` struct to calculate hash.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Links

* [Telegram Login Widget](https://core.telegram.org/widgets/login)
