package nacos

import (
	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin(Name, caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	//p, err := goosConfigParser(c)
	//if err != nil {
	//	return plugin.Error(Name, err)
	//}
	//
	////c.Next() // 'goosconfig'
	////log.Println("goos config plugin start")
	////for c.NextArg() {
	////	log.Println(c.Val())
	////}
	//
	//dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
	//	p.Next = next
	//	return p
	//})

	go StartServer()

	return nil
}

//func goosConfigParser(c *caddy.Controller) (*GoosDiscovery, error) {
//	proxy := CreateProxy()
//
//	for c.Next() {
//		if c.NextBlock() {
//			for {
//				switch c.Val() {
//				case "schema":
//					if !c.NextArg() {
//						return &GoosConfigProxy{}, c.ArgErr()
//					}
//					proxy.schema = c.Val()
//				case "url":
//					if !c.NextArg() {
//						return &GoosConfigProxy{}, c.ArgErr()
//					}
//					proxy.url = c.Val()
//				case "port":
//					if !c.NextArg() {
//						return &GoosConfigProxy{}, c.ArgErr()
//					}
//					if port, err := strconv.Atoi(c.Val()); err != nil {
//						return &GoosConfigProxy{}, c.Err("port format error")
//					} else {
//						proxy.port = port
//					}
//				default:
//					if c.Val() != "}" {
//						return &GoosConfigProxy{}, c.Errf("unknown property '%s'", c.Val())
//					}
//				}
//				if !c.Next() {
//					break
//				}
//			}
//		}
//
//		return &proxy, nil
//	}
//	return &proxy, nil
//}