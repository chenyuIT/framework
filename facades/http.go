package facades

import (
	"github.com/chenyuIT/framework/contracts/http"
)

func RateLimiter() http.RateLimiter {
	return App().MakeRateLimiter()
}
