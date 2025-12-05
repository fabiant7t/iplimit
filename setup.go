package iplimit

import (
	"fmt"
	"strconv"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/log"
)

const name = "iplimit"

func init() {
	plugin.Register(name, setup)
}

func setup(c *caddy.Controller) error {
	logger := log.NewWithPlugin(name)

	for c.Next() {
		plug := &Plugin{name: name}

		args := c.RemainingArgs()
		switch len(args) {
		case 1:
			ipLimit, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Errorf("cannot parse %q as IP limit integer value: %v", args[0], err)
				return err
			}
			if ipLimit <= 0 {
				err := fmt.Errorf("config error: %s IP limit must be positive, got %d", name, ipLimit)
				logger.Error(err)
				return err
			}
			plug.IPLimit = ipLimit
		default:
			err := fmt.Errorf("config error: %s: expected exactly one argument: <ip_limit>", name)
			logger.Error(err)
			return err
		}

		dnsConfig := dnsserver.GetConfig(c)
		dnsConfig.AddPlugin(func(next plugin.Handler) plugin.Handler {
			plug.Next = next
			return plug
		})
	}

	return nil
}
