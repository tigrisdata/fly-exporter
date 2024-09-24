package provider

import (
	"log/slog"
	"time"

	"github.com/tigrisdata/fly-exporter/reporter"
)

type MetricProvider struct {
	interval time.Duration
	reporter *reporter.Reporter
}

func NewMetricProvider() MetricProvider {
	mp := MetricProvider{
		interval: 10 * time.Second,
		reporter: reporter.NewReporter(),
	}
	return mp
}

func (mp *MetricProvider) Collect() {
	ticker := time.NewTicker(mp.interval)

	if err := mp.reporter.CollectOnce(); err != nil {
		slog.Error("failed to collect metrics", "error", err)
	}

	for range ticker.C {
		reporter := reporter.NewReporter()
		if err := reporter.CollectOnce(); err != nil {
			slog.Error("failed to collect metric data:", "error", err)
		}
		time.Sleep(1 * time.Second)
		oldReporter := mp.reporter
		mp.reporter = reporter
		oldReporter.Close()
	}
	defer ticker.Stop()
}

func (mp *MetricProvider) ServeHttp() {
	mp.reporter.ServeHttp()
}

func (mp *MetricProvider) Close() {
	mp.reporter.Close()
}
