package services

import (
	"context"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
)

type WriterAnomaliesService interface {
	WriteAnomalies(ctx context.Context, data *frv1.FrequencyResponse) (err error)
	SetRange(startAnomaly int, endAnomaly int)
}
