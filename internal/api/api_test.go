package api

import "testing"

func TestCalculatorAdd(t *testing.T) {
    c := Calculator{Base: 10}
    if got := c.Add(5); got != 15 {
        t.Fatalf("want 15, got %d", got)
    }
}

func TestDivide(t *testing.T) {
    t.Run("normal", func(t *testing.T) {
        result, err := Divide(10, 2)
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if result != 5 {
            t.Fatalf("want 5, got %d", result)
        }
    })

    t.Run("error", func(t *testing.T) {
        if _, err := Divide(10, 0); err == nil {
            t.Fatal("expected error")
        }
    })
}
