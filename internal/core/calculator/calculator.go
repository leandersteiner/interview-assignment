package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	ErrDivByZero = errors.New("division by zero")
	ErrOverflow  = errors.New("result is infinite (overflow)")
	ErrNaN       = errors.New("result is not a number (NaN)")
)

type Result struct {
	Value      float64
	Expression string
}

type Calculator struct {
	precision int
}

func New(precision int) *Calculator {
	return &Calculator{precision: precision}
}

func (c *Calculator) Add(a, b float64) (Result, error) {
	return c.compute(a, b, a+b, "+")
}

func (c *Calculator) Subtract(a, b float64) (Result, error) {
	return c.compute(a, b, a-b, "-")
}

func (c *Calculator) Multiply(a, b float64) (Result, error) {
	return c.compute(a, b, a*b, "*")
}

func (c *Calculator) Divide(a, b float64) (Result, error) {
	if b == 0 {
		return Result{}, ErrDivByZero
	}
	return c.compute(a, b, a/b, "/")
}

func (c *Calculator) compute(a, b, result float64, op string) (Result, error) {
	if err := validateFloat(result); err != nil {
		return Result{}, err
	}

	scale := math.Pow10(c.precision)
	result = math.Round(result*scale) / scale

	format := "%." + strconv.Itoa(c.precision) + "f"
	expr := fmt.Sprintf(format+" %s "+format+" = "+format, a, op, b, result)

	return Result{
		Value:      result,
		Expression: expr,
	}, nil
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
