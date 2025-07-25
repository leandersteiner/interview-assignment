package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

var (
	ErrDivByZero = errors.New("division by zero")
	ErrOverflow  = errors.New("result is infinite (overflow)")
	ErrNaN       = errors.New("result is not a number (NaN)")
)

type storer interface {
	Store(Result)
}

type Service struct {
	precision int
	saver     storer
}

func NewService(precision int, saver storer) *Service {
	return &Service{
		precision,
		saver,
	}
}

func (s *Service) Add(a, b float64) (Result, error) {
	return s.calc(a, b, a+b, "+")
}

func (s *Service) Sub(a, b float64) (Result, error) {
	return s.calc(a, b, a-b, "-")
}

func (s *Service) Mul(a, b float64) (Result, error) {
	return s.calc(a, b, a*b, "*")
}

func (s *Service) Div(a, b float64) (Result, error) {
	if b == 0 {
		return Result{}, ErrDivByZero
	}
	return s.calc(a, b, a/b, "/")
}

func (s *Service) calc(a, b, result float64, op string) (Result, error) {
	if err := validateFloat(result); err != nil {
		return Result{}, err
	}

	scale := math.Pow10(s.precision)
	result = math.Round(result*scale) / scale

	format := "%." + strconv.Itoa(s.precision) + "f"
	expr := fmt.Sprintf(format+" %s "+format+" = "+format, a, op, b, result)

	res := Result{
		Value:      result,
		Expression: expr,
		Created:    time.Now(),
	}

	s.saver.Store(res)
	return res, nil
}

func validateFloat(val float64) error {
	if math.IsNaN(val) {
		return ErrNaN
	}
	if math.IsInf(val, 0) {
		return ErrOverflow
	}
	return nil
}
