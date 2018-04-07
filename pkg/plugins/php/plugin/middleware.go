package plugin

import (
	"strings"

	"github.com/tonyhhyip/vodka"
)

func stripJson(next vodka.Handler) vodka.Handler {
	return vodka.HandlerFunc(func(c *vodka.Context) {
		path := c.Request.URL.Path
		if strings.HasSuffix(path, ".json") {
			c.Request.URL.Path = path[:strings.LastIndex(path, ".json")]
		}
		next.Handle(c)
	})
}
