package worker

import (
	"context"
	"fmt"
	"time"
)

type Worker struct {
	delay time.Duration
	svc   Service
}

func New(svc Service, delay time.Duration) *Worker {
	return &Worker{
		svc:   svc,
		delay: delay,
	}
}

func (w *Worker) RunNotify(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nworker stopped by context")
			return
		case <-time.After(w.delay):
		}
		_ = w.svc.NotifyTodayWather(ctx)
	}
}
