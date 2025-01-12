package wrapper

import "context"

type RunInterface interface {
	Run(ctx context.Context) int
}

type Interface interface {
	RunInterface
}

type NewWrapper func() Interface
