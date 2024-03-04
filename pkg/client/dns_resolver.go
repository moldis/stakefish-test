package client

import (
	"context"
	"net"
)

//go:generate mockery --name=DNSResolver --structname=MockDNSResolver --outpkg=client --output ./mocks --filename dns_resolver_mock.go
type DNSResolver interface {
	ResolveDomainToIPV4(ctx context.Context, host string) ([]net.IP, error)
}

type dnsResolver struct {
}

func NewDNSResolver() *dnsResolver {
	return &dnsResolver{}
}

func (s *dnsResolver) ResolveDomainToIPV4(ctx context.Context, host string) ([]net.IP, error) {
	return net.DefaultResolver.LookupIP(ctx, "ip4", host)
}
