package count

import (
	"context"
)

type CountSource func(context.Context) (uint64, error)

type CountOutput func(context.Context, uint64) error

func (s CountSource) OutputCount(ctx context.Context, out CountOutput) error {
	cnt, e := s(ctx)
	if nil != e {
		return e
	}
	return out(ctx, cnt)
}
