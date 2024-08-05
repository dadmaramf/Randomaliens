package storage

import (
	"context"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
)

type FrequencyStorage interface {
	CreateTable(ctx context.Context) error
	Insert(ctx context.Context, data *frv1.FrequencyResponse) error
}

type Storage interface {
	FrequencyStorage
}
