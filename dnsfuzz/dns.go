package dnsfuzz

import (
	"context"
	"fmt"
	"net"
	"time"
)

type dnsConfig struct {
	domain   string
	resolver *net.Resolver
	ctx      context.Context
}

func NewDnsConfig(ctx context.Context, domain string) dnsConfig {
	resolver := net.DefaultResolver

	return dnsConfig{
		domain:   domain,
		resolver: resolver,
		ctx:      ctx,
	}
}

func (dc dnsConfig) ProcessWord(word string) (string, error) {

	ctx, cancel := context.WithTimeout(dc.ctx, 2*time.Second)

	completeUrl := fmt.Sprintf("%s.%s", word, dc.domain)

	ips, err := dc.resolver.LookupNetIP(ctx, "ip", completeUrl)
	_, isDNSErr := err.(*net.DNSError)
	if isDNSErr {
		cancel()
		return "", nil
	}

	if err != nil {
		cancel()
		return "", err
	}

	if len(ips) == 0 {
		cancel()
		return "", err
	}

	cancel()
	return completeUrl, nil
}
