package service

import (
	client "code.stakefish.test/service/ip_validator/pkg/client/mocks"
	repository "code.stakefish.test/service/ip_validator/pkg/repository/mocks"
	"context"
	"github.com/stretchr/testify/mock"
	"net"
	"testing"
)

func TestLookup(t *testing.T) {
	dnsClientMock := client.NewMockDNSResolver(t)
	repoMock := repository.NewMockQueries(t)

	srvLookup := NewLookupService(dnsClientMock, repoMock)

	dnsClientMock.On("ResolveDomainToIPV4", mock.Anything, "google.com").Return([]net.IP{[]byte("192.168.0.1")}, nil).Once()
	repoMock.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

	srvLookup.LookupDomainIPV4(context.Background(), "192.168.0.1", "google.com")
}
