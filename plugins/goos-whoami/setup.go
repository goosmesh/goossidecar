package goos_whoami

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
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
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Whoami{}
	})


	return nil
}
