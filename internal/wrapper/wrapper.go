package wrapper

import (
	"context"
)

type wrapperInterface interface {
	Run(ctx context.Context) int
	Exit(code int)
	ExitWithPrint(code int, msg string)
	ExitWithPrintln(code int, msg string)
}

type Interface interface {
	wrapperInterface
}

type NewWrapper func(name string) Interface
