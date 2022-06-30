package util

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartProcessAtBackground(ps ...func() error) {
	for _, p := range ps {
		if p != nil {
			go func(_p func() error) {
				_ = _p()
			}(p)
		}
	}
}

func GracefullStopProcessAtBackground(duration time.Duration, ps ...func(ctx context.Context) error) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	for _, p := range ps {
		if p == nil {
			continue
		}
		ctx, stop := context.WithTimeout(context.Background(), duration)
		defer stop()
		_ = p(ctx)

	}
}
