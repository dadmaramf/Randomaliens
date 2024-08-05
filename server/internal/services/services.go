package services

import (
	"frequency_service/internal/services/frequency"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Services struct {
	ServiceFrequency
}

func NewServices() *Services {
	return &Services{
		frequency.NewFrequencyService(),
	}
}

type ServiceFrequency interface {
	GenerateInit()
	UpdateGenerateTime() *timestamppb.Timestamp
	UpdateGenerateFrequency() float64
	GetUUID() string
	GetMean() float64
	GetStddev() float64
}
