package cbor2iter

import (
	"context"
	"iter"

	cc "github.com/takanoriyanagitani/go-cbor-count"
)

type IterSourceEmpty func() iter.Seq[struct{}]

type Counter func(context.Context, iter.Seq[struct{}]) (uint64, error)

func (i IterSourceEmpty) ToCountSource(c Counter) cc.CountSource {
	return func(ctx context.Context) (uint64, error) {
		return c(ctx, i())
	}
}

func CounterDefault(ctx context.Context, i iter.Seq[struct{}]) (uint64, error) {
	var cnt uint64 = 0
	for range i {
		select {
		case <-ctx.Done():
			return cnt, ctx.Err()
		default:
		}
		cnt += 1
	}
	return cnt, nil
}
