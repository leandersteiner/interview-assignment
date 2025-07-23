package calculator

import (
	"math"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int
		wantValue float64
		wantExpr  string
		wantErr   error
	}{
		{
			name:      "positive numbers",
			a:         1.5,
			b:         2.3,
			precision: 2,
			wantValue: 3.8,
			wantExpr:  "1.50 + 2.30 = 3.80",
			wantErr:   nil,
		},
		{
			name:      "zero and positive number",
			a:         0,
			b:         5.2,
			precision: 1,
			wantValue: 5.2,
			wantExpr:  "0.0 + 5.2 = 5.2",
			wantErr:   nil,
		},
		{
			name:      "zero and zero",
			a:         0,
			b:         0,
			precision: 0,
			wantValue: 0,
			wantExpr:  "0 + 0 = 0",
			wantErr:   nil,
		},
		{
			name:      "negative numbers",
			a:         -1.1,
			b:         -2.2,
			precision: 2,
			wantValue: -3.3,
			wantExpr:  "-1.10 + -2.20 = -3.30",
			wantErr:   nil,
		},
		{
			name:      "large numbers",
			a:         1e10,
			b:         2e10,
			precision: 0,
			wantValue: 3e10,
			wantExpr:  "10000000000 + 20000000000 = 30000000000",
			wantErr:   nil,
		},
		{
			name:      "floating-point precision limit",
			a:         1.0000000001,
			b:         2.0000000002,
			precision: 10,
			wantValue: 3.0000000003,
			wantExpr:  "1.0000000001 + 2.0000000002 = 3.0000000003",
			wantErr:   nil,
		},
		{
			name:      "invalid float result (NaN)",
			a:         math.Inf(1),
			b:         math.Inf(-1),
			precision: 2,
			wantErr:   ErrNaN,
		},
		{
			name:      "invalid float result (overflow)",
			a:         math.MaxFloat64,
			b:         math.MaxFloat64,
			precision: 2,
			wantErr:   ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.precision)
			got, err := c.Add(tt.a, tt.b)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got.Value != tt.wantValue {
				t.Errorf("Add() got value = %v, want %v", got.Value, tt.wantValue)
			}

			if got.Expression != tt.wantExpr {
				t.Errorf("Add() got expression = %q, want %q", got.Expression, tt.wantExpr)
			}
		})
	}
}

func TestCalculator_Subtract(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int
		wantValue float64
		wantExpr  string
		wantErr   error
	}{
		{
			name:      "positive numbers",
			a:         5.4,
			b:         3.2,
			precision: 2,
			wantValue: 2.2,
			wantExpr:  "5.40 - 3.20 = 2.20",
			wantErr:   nil,
		},
		{
			name:      "zero and positive number",
			a:         0,
			b:         4.7,
			precision: 1,
			wantValue: -4.7,
			wantExpr:  "0.0 - 4.7 = -4.7",
			wantErr:   nil,
		},
		{
			name:      "zero and zero",
			a:         0,
			b:         0,
			precision: 0,
			wantValue: 0,
			wantExpr:  "0 - 0 = 0",
			wantErr:   nil,
		},
		{
			name:      "negative numbers",
			a:         -5.5,
			b:         -2.2,
			precision: 2,
			wantValue: -3.3,
			wantExpr:  "-5.50 - -2.20 = -3.30",
			wantErr:   nil,
		},
		{
			name:      "large numbers",
			a:         1e10,
			b:         5e9,
			precision: 0,
			wantValue: 5e9,
			wantExpr:  "10000000000 - 5000000000 = 5000000000",
			wantErr:   nil,
		},
		{
			name:      "floating-point precision limit",
			a:         3.0000000003,
			b:         2.0000000002,
			precision: 10,
			wantValue: 1.0000000001,
			wantExpr:  "3.0000000003 - 2.0000000002 = 1.0000000001",
			wantErr:   nil,
		},
		{
			name:      "invalid float result (NaN)",
			a:         math.Inf(1),
			b:         math.Inf(1),
			precision: 2,
			wantErr:   ErrNaN,
		},
		{
			name:      "invalid float result (overflow)",
			a:         math.MaxFloat64,
			b:         -math.MaxFloat64,
			precision: 2,
			wantErr:   ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.precision)
			got, err := c.Subtract(tt.a, tt.b)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Subtract() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got.Value != tt.wantValue {
				t.Errorf("Subtract() got value = %v, want %v", got.Value, tt.wantValue)
			}

			if got.Expression != tt.wantExpr {
				t.Errorf("Subtract() got expression = %q, want %q", got.Expression, tt.wantExpr)
			}
		})
	}
}

func TestCalculator_Multiply(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int
		wantValue float64
		wantExpr  string
		wantErr   error
	}{
		{
			name:      "positive numbers",
			a:         2.5,
			b:         3.2,
			precision: 2,
			wantValue: 8.0,
			wantExpr:  "2.50 * 3.20 = 8.00",
			wantErr:   nil,
		},
		{
			name:      "zero multiplied by positive number",
			a:         0,
			b:         4.7,
			precision: 1,
			wantValue: 0,
			wantExpr:  "0.0 * 4.7 = 0.0",
			wantErr:   nil,
		},
		{
			name:      "zero multiplied by zero",
			a:         0,
			b:         0,
			precision: 0,
			wantValue: 0,
			wantExpr:  "0 * 0 = 0",
			wantErr:   nil,
		},
		{
			name:      "negative numbers",
			a:         -4.5,
			b:         -2,
			precision: 1,
			wantValue: 9,
			wantExpr:  "-4.5 * -2.0 = 9.0",
			wantErr:   nil,
		},
		{
			name:      "positive and negative numbers",
			a:         5,
			b:         -1.5,
			precision: 2,
			wantValue: -7.5,
			wantExpr:  "5.00 * -1.50 = -7.50",
			wantErr:   nil,
		},
		{
			name:      "large numbers",
			a:         1e6,
			b:         1e3,
			precision: 0,
			wantValue: 1e9,
			wantExpr:  "1000000 * 1000 = 1000000000",
			wantErr:   nil,
		},
		{
			name:      "floating-point precision limit",
			a:         1.0000000001,
			b:         2.0000000002,
			precision: 10,
			wantValue: 2.0000000004,
			wantExpr:  "1.0000000001 * 2.0000000002 = 2.0000000004",
			wantErr:   nil,
		},
		{
			name:      "invalid float result (NaN)",
			a:         math.Inf(1),
			b:         0,
			precision: 2,
			wantErr:   ErrNaN,
		},
		{
			name:      "invalid float result (overflow)",
			a:         math.MaxFloat64,
			b:         2,
			precision: 2,
			wantErr:   ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.precision)
			got, err := c.Multiply(tt.a, tt.b)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got.Value != tt.wantValue {
				t.Errorf("Multiply() got value = %v, want %v", got.Value, tt.wantValue)
			}

			if got.Expression != tt.wantExpr {
				t.Errorf("Multiply() got expression = %q, want %q", got.Expression, tt.wantExpr)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	tests := []struct {
		name      string
		a         float64
		b         float64
		precision int
		wantValue float64
		wantExpr  string
		wantErr   error
	}{
		{
			name:      "positive numbers",
			a:         8.4,
			b:         2.0,
			precision: 2,
			wantValue: 4.2,
			wantExpr:  "8.40 / 2.00 = 4.20",
			wantErr:   nil,
		},
		{
			name:      "division by zero",
			a:         5,
			b:         0,
			precision: 2,
			wantErr:   ErrDivByZero,
		},
		{
			name:      "zero divided by positive number",
			a:         0,
			b:         5,
			precision: 2,
			wantValue: 0,
			wantExpr:  "0.00 / 5.00 = 0.00",
			wantErr:   nil,
		},
		{
			name:      "zero divided by zero",
			a:         0,
			b:         0,
			precision: 2,
			wantErr:   ErrDivByZero,
		},
		{
			name:      "negative numbers",
			a:         -8.4,
			b:         -2.0,
			precision: 2,
			wantValue: 4.2,
			wantExpr:  "-8.40 / -2.00 = 4.20",
			wantErr:   nil,
		},
		{
			name:      "large numbers",
			a:         1e10,
			b:         2e5,
			precision: 0,
			wantValue: 50000,
			wantExpr:  "10000000000 / 200000 = 50000",
			wantErr:   nil,
		},
		{
			name:      "floating-point precision limit",
			a:         1.0000000001,
			b:         2.0000000002,
			precision: 10,
			wantValue: 0.5000000000,
			wantExpr:  "1.0000000001 / 2.0000000002 = 0.5000000000",
			wantErr:   nil,
		},
		{
			name:      "invalid float result (NaN)",
			a:         math.Inf(1),
			b:         math.Inf(1),
			precision: 2,
			wantErr:   ErrNaN,
		},
		{
			name:      "invalid float result (overflow)",
			a:         math.MaxFloat64,
			b:         math.SmallestNonzeroFloat64,
			precision: 2,
			wantErr:   ErrOverflow,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.precision)
			got, err := c.Divide(tt.a, tt.b)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if got.Value != tt.wantValue {
				t.Errorf("Divide() got value = %v, want %v", got.Value, tt.wantValue)
			}

			if got.Expression != tt.wantExpr {
				t.Errorf("Divide() got expression = %q, want %q", got.Expression, tt.wantExpr)
			}
		})
	}
}
