package worker

import "context"

type Service interface {
	NotifyTodayWather(context.Context) error
}
