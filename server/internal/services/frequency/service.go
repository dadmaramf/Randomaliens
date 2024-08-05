package frequency

import (
	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
)

type frequencyService struct {
	SessionId string
	Mean      float64
	Stddev    float64
}

func NewFrequencyService() *frequencyService {
	return &frequencyService{}
}

func (f *frequencyService) GenerateInit() {
	f.Mean = -10 + rand.Float64()*20
	f.Stddev = 0.3 + rand.Float64()*(1.5-0.3)
	f.SessionId = uuid.NewV4().String()
}

func (f *frequencyService) GetUUID() string {
	return f.SessionId
}

func (f *frequencyService) UpdateGenerateTime() *timestamppb.Timestamp {
	return timestamppb.Now()
}

func (f *frequencyService) UpdateGenerateFrequency() float64 {
	return rand.NormFloat64()*f.Stddev + f.Mean
}

func (f *frequencyService) GetMean() float64 {
	return f.Mean
}

func (f *frequencyService) GetStddev() float64 {
	return f.Stddev
}
