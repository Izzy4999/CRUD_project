package initializers

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func SessionStorage() {
	Store = session.New(
		session.Config{
			Expiration:     time.Hour * 2,
			CookieHTTPOnly: false,
			CookieSecure:   true,
		},
	)
}
