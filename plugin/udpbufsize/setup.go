package udpbufsize

import (
	"strconv"

	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("udpbufsize", setup) }

func setup(c *caddy.Controller) error {
	bufsize, err := parse(c)
	if err != nil {
		return plugin.Error("bufsize", err)
	}

	cfg := dnsserver.GetConfig(c)
	cfg.UdpReadBufferBytes = bufsize

	// not really a plugin, just modifies the config
	return nil
}

func parse(c *caddy.Controller) (int, error) {
	const defaultBufSize = 0
	for c.Next() {
		args := c.RemainingArgs()
		switch len(args) {
		case 1:
			// Specified value is needed to verify
			bufsize, err := strconv.Atoi(args[0])
			if err != nil {
				return -1, plugin.Error("udpbufsize", c.ArgErr())
			}
			if bufsize <= 0 {
				return -1, plugin.Error("udpbufsize", c.ArgErr())
			}
			return bufsize, nil
		default:
			// Only 1 argument is acceptable
			return -1, plugin.Error("udpbufsize", c.ArgErr())
		}
	}
	return -1, plugin.Error("udpbufsize", c.ArgErr())
}
