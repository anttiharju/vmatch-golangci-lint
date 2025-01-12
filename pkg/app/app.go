package app

import "context"

type Interface interface {
	Run(ctx context.Context) int
}

type NewAppInterface func(versionFileName string) Interface
