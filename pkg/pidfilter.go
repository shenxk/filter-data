package filter

//"log"
//. "math"

type pidData struct {
	setPoint                         float64
	Proportion, Integral, Derivative float64
	LastError, PrevError, SumError   float64
}

func InitPidData(p, i, d float64) pidData {
	pid := pidData{Proportion: p, Integral: i, Derivative: d}
	return pid
}

func (pp *pidData) Update(NextPoint float64) float64 {
	var dError, Error float64
	Error = pp.setPoint - NextPoint
	pp.SumError += Error
	dError = pp.LastError - pp.PrevError
	pp.PrevError = pp.LastError
	pp.LastError = Error

	return (pp.Proportion*Error +
		pp.Integral*pp.SumError + pp.Derivative*dError)
}
