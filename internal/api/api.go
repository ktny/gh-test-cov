package api

import "errors"

// Calculator provides simple integer calculations.
type Calculator struct {
    Base int
}

// Add adds v to the base.
func (c Calculator) Add(v int) int {
    return c.Base + v
}

// Divide divides numerator by denominator.
func Divide(numerator, denominator int) (int, error) {
    if denominator == 0 {
        return 0, errors.New("division by zero")
    }
    return numerator / denominator, nil
}
