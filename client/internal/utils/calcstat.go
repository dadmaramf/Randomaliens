package utils

import "math"

type Stat struct {
	mean   float64
	stddev float64
	n      int
}

func NewStat() *Stat {
	return &Stat{}
}

func (s *Stat) GetCount() int {
	return s.n
}
func (s *Stat) AddStream(freq float64) {
	s.n++
	oldMean := s.mean
	s.mean += (freq - s.mean) / float64(s.n)
	s.stddev += (freq - oldMean) * (freq - s.mean)
}

func (s *Stat) GetStats() (float64, float64) {
	stdDev := math.Sqrt(s.stddev / float64(s.n-1))
	return s.mean, stdDev
}
