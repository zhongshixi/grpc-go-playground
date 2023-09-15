package utils

import (
	"sync"
	"sync/atomic"
	"time"

	"log/slog"
)

func SpikeWithFunc(concurrentSpike int64, reqRespfunc func(id int64) (int64, error)) {
	startTime := time.Now()
	var wg sync.WaitGroup

	totalCount := int64(0)
	successCount := atomic.Int64{}
	failCount := atomic.Int64{}
	successSize := atomic.Int64{}

	defer func() {
		slog.Info("test finished:",
			slog.Duration("time_elaspsed", time.Since(startTime)),
			slog.Any("success", successCount.Load()),
			slog.Any("fail", failCount.Load()),
			slog.Any("success_rate", float64(successCount.Load())/float64(totalCount)),
			slog.Any("success_sizes(MB)", float64(successSize.Load())/1024/1024),
			slog.Any("throughput_(MB/s)", float64(successSize.Load())/1024/1024/time.Since(startTime).Seconds()),
		)
	}()

	for {
		if totalCount == concurrentSpike {
			break
		}
		totalCount = totalCount + 1
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()

			size, err := reqRespfunc(id)
			if err != nil {
				failCount.Add(1)
				slog.Error("resp failed", slog.Any("error", err))
				return
			} else {
				successCount.Add(1)
				successSize.Add(size)
			}
		}(totalCount)
	}

	wg.Wait()
}
