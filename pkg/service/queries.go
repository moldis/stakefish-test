package service

import (
	"context"
	"fmt"
	"time"

	"code.stakefish.test/service/ip_validator/pkg/repository"

	"code.stakefish.test/service/ip_validator/pkg/client"
	"code.stakefish.test/service/ip_validator/pkg/model"
)

type QueriesService interface {
	LookupDomainIPV4(ctx context.Context, clientIP, domain string) (*model.Query, error)
	History(ctx context.Context, skip, limit int64) ([]model.Query, error)
}

type queriesService struct {
	resolver client.DNSResolver
	repo     repository.Queries
}

func NewLookupService(resolver client.DNSResolver, repo repository.Queries) *queriesService {
	return &queriesService{resolver: resolver, repo: repo}
}

func (s *queriesService) LookupDomainIPV4(ctx context.Context, clientIP, domain string) (*model.Query, error) {
	res, err := s.resolver.ResolveDomainToIPV4(ctx, domain)
	if err != nil {
		return nil, err
	}

	addresses := make([]model.Address, 0)
	for _, ip := range res {
		addresses = append(addresses, model.Address{IP: ip.To4().String()})
	}

	query := &model.Query{
		Addresses: addresses,
		ClientIP:  clientIP,
		Domain:    domain,
		CreatedAt: time.Now().UTC().Unix(),
	}

	err = s.repo.Create(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("internal error")
	}

	return query, nil
}

func (s *queriesService) History(ctx context.Context, skip, limit int64) ([]model.Query, error) {
	return s.repo.List(ctx, skip, limit)
}
