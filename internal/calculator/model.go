package calculator

import "time"

type Result struct {
	Value      float64
	Expression string
	Created    time.Time
}

func (r Result) String() string {
	return r.Expression
}
