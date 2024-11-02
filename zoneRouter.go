package zonerouter

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type ZoneRouter struct {
	Next plugin.Handler
}

func (z *ZoneRouter) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return plugin.NextOrFailure(z.Name(), z.Next, ctx, w, r)
}

func (z *ZoneRouter) Name() string { return "zone_router" }
