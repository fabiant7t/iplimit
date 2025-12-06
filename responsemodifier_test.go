package iplimit_test

import (
	"reflect"
	"testing"

	"github.com/coredns/coredns/plugin/iplimit"
	"github.com/miekg/dns"
	"github.com/pkg/errors"
)

func TestIPLimitResponseModifier(t *testing.T) {
	ipv4Answer, err := NewIPv4TestAnswer()
	if err != nil {
		t.Error(err)
	}

	ipv6Answer, err := NewIPv6TestAnswer()
	if err != nil {
		t.Error(err)
	}

	for _, tc := range []struct {
		name    string
		ipLimit int
		msg     *dns.Msg
		want    *dns.Msg
	}{
		{"IPv4: CNAME and three As", 3, &dns.Msg{Answer: ipv4Answer}, &dns.Msg{Answer: ipv4Answer[:1+3]}},
		{"IPv4: Must return at least 1 IP", 0, &dns.Msg{Answer: ipv4Answer}, &dns.Msg{Answer: ipv4Answer[:1+1]}},
		{"IPv4: Limit equals answers", len(ipv4Answer), &dns.Msg{Answer: ipv4Answer}, &dns.Msg{Answer: ipv4Answer}},
		{"IPv4: Limit is higher than answers", len(ipv4Answer) + 1, &dns.Msg{Answer: ipv4Answer}, &dns.Msg{Answer: ipv4Answer}},
		{"IPv6: CNAME and three AAAAs", 3, &dns.Msg{Answer: ipv6Answer}, &dns.Msg{Answer: ipv6Answer[:1+3]}},
		{"IPv6: Must return at least 1 IP", 0, &dns.Msg{Answer: ipv6Answer}, &dns.Msg{Answer: ipv6Answer[:1+1]}},
		{"IPv6: Limit equals answers", len(ipv6Answer), &dns.Msg{Answer: ipv6Answer}, &dns.Msg{Answer: ipv6Answer}},
		{"IPv6: Limit is higher than answers", len(ipv6Answer) + 1, &dns.Msg{Answer: ipv6Answer}, &dns.Msg{Answer: ipv6Answer}},
		{"nil returns nil", 3, nil, nil},
	} {
		mod := iplimit.NewIPLimitResponseModifier(tc.ipLimit)
		mod(tc.msg)
		if got, want := tc.msg, tc.want; !reflect.DeepEqual(got, want) {
			t.Errorf("%s: Got %+v, want %+v", tc.name, got, want)
		}
	}
}

func NewIPv4TestAnswer() ([]dns.RR, error) {
	lines := []string{
		"k8s.topfstedt.com. 60 IN CNAME some.pizzabox.es.",
		"some.pizzabox.es. 60 IN A 1.3.3.7",
		"some.pizzabox.es. 60 IN A 1.3.5.5",
		"some.pizzabox.es. 60 IN A 1.45.3.7",
		"some.pizzabox.es. 60 IN A 5.3.4.5",
		"some.pizzabox.es. 60 IN A 7.3.5.7",
		"some.pizzabox.es. 60 IN A 8.0.4.7",
		"some.pizzabox.es. 60 IN A 83.4.5.7",
	}
	rrs := make([]dns.RR, 0, len(lines))
	for _, line := range lines {
		rr, err := dns.NewRR(line)
		if err != nil {
			return nil, errors.Errorf("error parsing line '%s' to resource record: %s", line, err)
		}
		rrs = append(rrs, rr)
	}
	return rrs, nil
}

func NewIPv6TestAnswer() ([]dns.RR, error) {
	lines := []string{
		"k8s.topfstedt.com. 60 IN CNAME some.pizzabox.es.",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:1111",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:2222",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:3333",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:4444",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:5555",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:6666",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:7777",
		"some.pizzabox.es. 60 IN AAAA 2001:0db8:85a3:1234:5678:9abc:def0:8888",
	}
	rrs := make([]dns.RR, 0, len(lines))
	for _, line := range lines {
		rr, err := dns.NewRR(line)
		if err != nil {
			return nil, errors.Errorf("error parsing line '%s' to resource record: %s", line, err)
		}
		rrs = append(rrs, rr)
	}
	return rrs, nil
}
