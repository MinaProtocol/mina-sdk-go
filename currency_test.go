package mina

import (
	"errors"
	"testing"
)

func TestNewCurrency(t *testing.T) {
	c := NewCurrency(5)
	if c.Nanomina() != 5_000_000_000 {
		t.Errorf("expected 5000000000, got %d", c.Nanomina())
	}
	if c.Mina() != "5.000000000" {
		t.Errorf("expected 5.000000000, got %s", c.Mina())
	}
}

func TestCurrencyFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"2.5", 2_500_000_000},
		{"100", 100_000_000_000},
		{"0.5", 500_000_000},
		{"1.23", 1_230_000_000},
	}
	for _, tt := range tests {
		c, err := CurrencyFromString(tt.input)
		if err != nil {
			t.Errorf("CurrencyFromString(%q) error: %v", tt.input, err)
			continue
		}
		if c.Nanomina() != tt.expected {
			t.Errorf("CurrencyFromString(%q) = %d, want %d", tt.input, c.Nanomina(), tt.expected)
		}
	}
}

func TestCurrencyFromNanomina(t *testing.T) {
	c := CurrencyFromNanomina(500_000_000)
	if c.Mina() != "0.500000000" {
		t.Errorf("expected 0.500000000, got %s", c.Mina())
	}
}

func TestCurrencyFromGraphQL(t *testing.T) {
	c, err := CurrencyFromGraphQL("1500000000")
	if err != nil {
		t.Fatal(err)
	}
	if c.Nanomina() != 1_500_000_000 {
		t.Errorf("expected 1500000000, got %d", c.Nanomina())
	}
}

func TestNanominaString(t *testing.T) {
	c := NewCurrency(3)
	if c.NanominaString() != "3000000000" {
		t.Errorf("expected 3000000000, got %s", c.NanominaString())
	}
}

func TestCurrencyNanoRejectsFloat(t *testing.T) {
	// In Go, this is handled by type system - CurrencyFromNanomina only takes uint64
	// but we can test that string parsing rejects invalid formats
	_, err := CurrencyFromString("1.2345678901") // >9 decimal places
	if err == nil {
		t.Error("expected error for >9 decimal places")
	}
}

func TestAddition(t *testing.T) {
	a := NewCurrency(1)
	b := NewCurrency(2)
	result := a.Add(b)
	if result.Nanomina() != 3_000_000_000 {
		t.Errorf("expected 3000000000, got %d", result.Nanomina())
	}
}

func TestSubtraction(t *testing.T) {
	a := NewCurrency(3)
	b := NewCurrency(1)
	result, err := a.Sub(b)
	if err != nil {
		t.Fatal(err)
	}
	if result.Nanomina() != 2_000_000_000 {
		t.Errorf("expected 2000000000, got %d", result.Nanomina())
	}
}

func TestSubtractionUnderflow(t *testing.T) {
	a := NewCurrency(1)
	b := NewCurrency(2)
	_, err := a.Sub(b)
	if err == nil {
		t.Error("expected underflow error")
	}
	var uf *CurrencyUnderflowError
	if !errors.As(err, &uf) {
		t.Errorf("expected CurrencyUnderflowError, got %T", err)
	}
}

func TestMultiplication(t *testing.T) {
	c := NewCurrency(2)
	result := c.Mul(3)
	if result.Nanomina() != 6_000_000_000 {
		t.Errorf("expected 6000000000, got %d", result.Nanomina())
	}
}

func TestEquality(t *testing.T) {
	a := NewCurrency(1)
	b := CurrencyFromNanomina(1_000_000_000)
	if !a.Equal(b) {
		t.Error("expected equal")
	}
}

func TestComparison(t *testing.T) {
	a := NewCurrency(1)
	b := NewCurrency(2)
	if !a.Less(b) {
		t.Error("expected a < b")
	}
	if !a.LessOrEqual(b) {
		t.Error("expected a <= b")
	}
	if !b.Greater(a) {
		t.Error("expected b > a")
	}
	if !b.GreaterOrEqual(a) {
		t.Error("expected b >= a")
	}
}

func TestString(t *testing.T) {
	c := CurrencyFromNanomina(0)
	if c.String() != "0.000000000" {
		t.Errorf("expected 0.000000000, got %s", c.String())
	}
}

func TestRepr(t *testing.T) {
	c := MustCurrencyFromString("1.23")
	if c.Mina() != "1.230000000" {
		t.Errorf("expected 1.230000000, got %s", c.Mina())
	}
}

func TestSmallNanominaDisplay(t *testing.T) {
	c := CurrencyFromNanomina(1)
	if c.Mina() != "0.000000001" {
		t.Errorf("expected 0.000000001, got %s", c.Mina())
	}
}

func TestRandomCurrency(t *testing.T) {
	lower := NewCurrency(1)
	upper := NewCurrency(10)
	for i := 0; i < 50; i++ {
		r, err := RandomCurrency(lower, upper)
		if err != nil {
			t.Fatal(err)
		}
		if r.Less(lower) || r.Greater(upper) {
			t.Errorf("random %s out of range [%s, %s]", r, lower, upper)
		}
	}
}

func TestRandomCurrencyEqualBounds(t *testing.T) {
	c := NewCurrency(5)
	r, err := RandomCurrency(c, c)
	if err != nil {
		t.Fatal(err)
	}
	if !r.Equal(c) {
		t.Errorf("expected %s, got %s", c, r)
	}
}

func TestIsZero(t *testing.T) {
	if !CurrencyFromNanomina(0).IsZero() {
		t.Error("expected zero")
	}
	if NewCurrency(1).IsZero() {
		t.Error("expected not zero")
	}
}
