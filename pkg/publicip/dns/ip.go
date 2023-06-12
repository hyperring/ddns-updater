package dns

import (
	"context"
	"net/netip"
	"sync/atomic"
)

func (f *Fetcher) IP(ctx context.Context) (publicIP netip.Addr, err error) {
	return f.ip(ctx, f.client)
}

func (f *Fetcher) IP4(ctx context.Context) (publicIP netip.Addr, err error) {
	return f.ip(ctx, f.client4)
}

func (f *Fetcher) IP6(ctx context.Context) (publicIP netip.Addr, err error) {
	return f.ip(ctx, f.client6)
}

func (f *Fetcher) ip(ctx context.Context, client Client) (
	publicIP netip.Addr, err error) {
	index := int(atomic.AddUint32(f.ring.counter, 1)) % len(f.ring.providers)
	provider := f.ring.providers[index]
	return fetch(ctx, client, provider.data())
}
