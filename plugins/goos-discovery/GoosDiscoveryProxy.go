package goos_discovery

import (
	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	"net"
	"strings"
)

const (
	defaultTtl = 360
)

type GoosDiscoveryProxy struct {
	Next plugin.Handler
	schema string
	url string
	port int
	Ttl            uint32
}

func CreateProxy() GoosDiscoveryProxy {
	return GoosDiscoveryProxy{
		schema: "http",
		url: "localhost",
		port: 4321,
	}
}

type Zone struct {
	Name      string
	Locations map[string]struct{}
}

type Record struct {
	A     []A_Record `json:"a,omitempty"`
	AAAA  []AAAA_Record `json:"aaaa,omitempty"`
}

type A_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

type AAAA_Record struct {
	Ttl uint32 `json:"ttl,omitempty"`
	Ip  net.IP `json:"ip"`
}

func (proxy *GoosDiscoveryProxy) A(name string, z *Zone, record *Record) (answers, extras []dns.RR) {
	for _, a := range record.A {
		if a.Ip == nil {
			continue
		}
		r := new(dns.A)
		r.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeA,
			Class: dns.ClassINET, Ttl: proxy.minTtl(a.Ttl)}
		r.A = a.Ip
		answers = append(answers, r)
	}
	return
}

func (proxy *GoosDiscoveryProxy) AAAA(name string, z *Zone, record *Record) (answers, extras []dns.RR) {
	for _, aaaa := range record.AAAA {
		if aaaa.Ip == nil {
			continue
		}
		r := new(dns.AAAA)
		r.Hdr = dns.RR_Header{Name: name, Rrtype: dns.TypeAAAA,
			Class: dns.ClassINET, Ttl: proxy.minTtl(aaaa.Ttl)}
		r.AAAA = aaaa.Ip
		answers = append(answers, r)
	}
	return
}

func (proxy *GoosDiscoveryProxy) minTtl(ttl uint32) uint32 {
	if proxy.Ttl == 0 && ttl == 0 {
		return defaultTtl
	}
	if proxy.Ttl == 0 {
		return ttl
	}
	if ttl == 0 {
		return proxy.Ttl
	}
	if proxy.Ttl < ttl {
		return proxy.Ttl
	}
	return  ttl
}

func (redis *GoosDiscoveryProxy) findLocation(query string, z *Zone) string {
	var (
		ok bool
		closestEncloser, sourceOfSynthesis string
	)

	// request for zone records
	if query == z.Name {
		return query
	}

	query = strings.TrimSuffix(query, "." + z.Name)

	if _, ok = z.Locations[query]; ok {
		return query
	}

	closestEncloser, sourceOfSynthesis, ok = splitQuery(query)
	for ok {
		ceExists := keyMatches(closestEncloser, z) || keyExists(closestEncloser, z)
		ssExists := keyExists(sourceOfSynthesis, z)
		if ceExists {
			if ssExists {
				return sourceOfSynthesis
			} else {
				return ""
			}
		} else {
			closestEncloser, sourceOfSynthesis, ok = splitQuery(closestEncloser)
		}
	}
	return ""
}

func (proxy *GoosDiscoveryProxy) get(key string, z *Zone) *Record {
	return nil
}

func keyExists(key string, z *Zone) bool {
	_, ok := z.Locations[key]
	return ok
}

func keyMatches(key string, z *Zone) bool {
	for value := range z.Locations {
		if strings.HasSuffix(value, key) {
			return true
		}
	}
	return false
}

func splitQuery(query string) (string, string, bool) {
	if query == "" {
		return "", "", false
	}
	var (
		splits []string
		closestEncloser string
		sourceOfSynthesis string
	)
	splits = strings.SplitAfterN(query, ".", 2)
	if len(splits) == 2 {
		closestEncloser = splits[1]
		sourceOfSynthesis = "*." + closestEncloser
	} else {
		closestEncloser = ""
		sourceOfSynthesis = "*"
	}
	return closestEncloser, sourceOfSynthesis, true
}