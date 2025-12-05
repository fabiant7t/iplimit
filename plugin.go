package iplimit

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type Plugin struct {
	name    string
	IPLimit int
	Next    plugin.Handler
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return plugin.NextOrFailure(
		p.Name(),
		p.Next,
		ctx,
		ModResponseWriter(
			w,
			NewLimitResponseModifier(p.IPLimit),
		),
		r,
	)
}
