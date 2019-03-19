package goos_discovery

import (
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

func (proxy *GoosConfigProxy) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}


	qname := state.Name()
	qtype := state.Type()

	location := proxy.findLocation(qname, nil)
	if len(location) == 0 { // empty, no results
		return proxy.errorResponse(state, "", dns.RcodeNameError, nil)
	}

	answers := make([]dns.RR, 0, 10)
	extras := make([]dns.RR, 0, 10)

	record := proxy.get(location, nil)

	switch qtype {
	case "A":
		answers, extras = proxy.A(qname, nil, record)
	case "AAAA":
		answers, extras = proxy.AAAA(qname, nil, record)
	default:
		return proxy.errorResponse(state, "", dns.RcodeNotImplemented, nil)
	}

	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative, m.RecursionAvailable, m.Compress = true, false, true

	m.Answer = append(m.Answer, answers...)
	m.Extra = append(m.Extra, extras...)

	state.SizeAndDo(m)
	m = state.Scrub(m)
	_ = w.WriteMsg(m)
	return dns.RcodeSuccess, nil

}


// Name implements the Handler interface.
func (proxy *GoosConfigProxy) Name() string { return Name }

func (proxy *GoosConfigProxy) errorResponse(state request.Request, zone string, rcode int, err error) (int, error) {
	m := new(dns.Msg)
	m.SetRcode(state.Req, rcode)
	m.Authoritative, m.RecursionAvailable, m.Compress = true, false, true

	// m.Ns, _ = redis.SOA(state.Name(), zone, nil)

	state.SizeAndDo(m)
	state.W.WriteMsg(m)
	// Return success as the rcode to signal we have written to the client.
	return dns.RcodeSuccess, err
}