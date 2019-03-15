package goos_whoiam

import (
	"log"
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("gooswhoami", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	c.Next() // 'gooswhoami'
	for c.NextArg() {
		log.Println(c.Val())
	}
	return nil
}
