package initializers

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func SessionStorage() {
	Store = session.New(
		session.Config{
			CookieHTTPOnly: true,
			Expiration:     time.Hour * 2,
		},
	)
}
