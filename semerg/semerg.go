package semerg

import (
	"context"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type ErrorGroup struct {
	ctx context.Context
	eg  *errgroup.Group
	sem *semaphore.Weighted
}

func NewMaxSharedCtx(ctx context.Context, max int) *ErrorGroup {
	eg, lctx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(int64(max))
	return &ErrorGroup{
		ctx: lctx,
		eg:  eg,
		sem: sem,
	}
}

func (g *ErrorGroup) Go(f func() error) {
	g.eg.Go(func() error {
		if err := g.sem.Acquire(g.ctx, 1); err != nil {
			return err
		}
		defer g.sem.Release(1)
		return f()
	})
}

func (g *ErrorGroup) Wait() error {
	return g.eg.Wait()
}
