package zonerouter

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() {
	plugin.Register("zone_router", setup)
}

func setup(c *caddy.Controller) error {
	zr := ZoneRouter{}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		zr.Next = next
		return &zr
	})

	return nil
}
