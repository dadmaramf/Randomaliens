package anomalies

import "math"

type Anomalies struct {
	k      int
	enable bool
}

func New(k int) *Anomalies {
	return &Anomalies{k: k, enable: true}
}

func (a *Anomalies) DisableAnomalies() {
	a.enable = false
}

func (a *Anomalies) EnableAnomalies() {
	a.enable = true
}

func (a *Anomalies) CalcAnomalies(freq float64, stddev float64, mean float64) bool {
	if !a.enable {
		return false
	}
	return math.Abs(freq-mean) > float64(a.k)*stddev
}

func (a *Anomalies) IsEnable() bool {
	return a.enable
}
