package services

import (
	"client/internal/analysis/anomalies"
	"client/internal/storage"
	"client/internal/utils"
	"context"
	"fmt"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
)

var _ WriterAnomaliesService = (*writer)(nil)

type writer struct {
	storage      storage.Storage
	anomalies    *anomalies.Anomalies
	stat         *utils.Stat
	startAnomaly int
	endAnomaly   int
}

func NewService(storage storage.Storage, stat *utils.Stat, k int) *writer {
	return &writer{
		anomalies: anomalies.New(k),
		storage:   storage,
		stat:      stat,
	}
}

func (f *writer) WriteAnomalies(ctx context.Context, data *frv1.FrequencyResponse) error {
	f.stat.AddStream(data.Frequency)

	mean, stddev := f.stat.GetStats()
	count := f.stat.GetCount()

	if f.anomalies.IsEnable() && !f.inRange(count) {
		f.anomalies.DisableAnomalies()
	} else if !f.anomalies.IsEnable() && f.inRange(count) {
		f.anomalies.EnableAnomalies()
	}

	if f.anomalies.CalcAnomalies(data.Frequency, stddev, mean) {
		fmt.Println("[a[a]]")
		// return f.storage.Insert(ctx, data)
	}

	return nil
}

func (f *writer) SetRange(startAnomaly, endAnomaly int) {
	f.startAnomaly = startAnomaly
	f.endAnomaly = endAnomaly
}

func (f *writer) inRange(number int) bool {
	if f.endAnomaly == 0 {
		return true
	}
	return number >= f.startAnomaly && number <= f.endAnomaly
}
