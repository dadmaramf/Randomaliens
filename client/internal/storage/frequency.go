package storage

import (
	"context"

	frv1 "github.com/dadmaramf/protos/gen/go/frequency_service"
	"github.com/uptrace/bun"
)

var _ FrequencyStorage = (*frequency)(nil)

type frequency struct {
	db *bun.DB
}

func NewFrequency(db *bun.DB) *frequency {
	return &frequency{db: db}
}

func (f *frequency) CreateTable(ctx context.Context) error {
	_, err := f.db.NewCreateTable().Model((*frv1.FrequencyResponse)(nil)).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (f *frequency) Insert(ctx context.Context, data *frv1.FrequencyResponse) error {
	_, err := f.db.NewInsert().Model(data).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
