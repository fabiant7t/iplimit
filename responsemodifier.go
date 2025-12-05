package iplimit

import "github.com/miekg/dns"

type ResponseModifier func(*dns.Msg) error

func NewIPLimitResponseModifier(ipLimit int) ResponseModifier {
	return func(res *dns.Msg) error {
		if res == nil {
			return nil
		}

		if len(res.Answer) <= ipLimit {
			return nil
		}

		filtered := make([]dns.RR, 0, len(res.Answer))
		ipCount := 0
		for _, rr := range res.Answer {
			switch rr.(type) {
			case *dns.A, *dns.AAAA:
				if ipCount < ipLimit {
					filtered = append(filtered, rr)
					ipCount++
				}
			default:
				filtered = append(filtered, rr)
			}
		}
		res.Answer = filtered
		return nil
	}
}
