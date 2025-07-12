// loan.go
// Defines the Loan struct and core billing engine logic.
// Provides thread-safe methods for making weekly payments,
// calculating outstanding balance, and detecting delinquency
// based on two or more consecutive missed payments.

package loan

import (
	"errors"
	"sync"
)

type Loan struct {
	TotalAmount   int
	Weeks         int
	WeeklyPayment int
	Payments      []bool
	mu            sync.Mutex
}

func NewLoan() *Loan {
	weeks := 50
	total := 5500000
	return &Loan{
		TotalAmount:   total,
		Weeks:         weeks,
		WeeklyPayment: total / weeks,
		Payments:      make([]bool, weeks),
	}
}

func (l *Loan) MakePayment(week int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if week < 1 || week > l.Weeks {
		return errors.New("invalid week number")
	}
	if l.Payments[week-1] {
		return errors.New("already paid")
	}
	l.Payments[week-1] = true
	return nil
}

func (l *Loan) GetOutstanding() int {
	l.mu.Lock()
	defer l.mu.Unlock()

	paid := 0
	for _, p := range l.Payments {
		if p {
			paid++
		}
	}
	return l.TotalAmount - paid*l.WeeklyPayment
}

func (l *Loan) IsDelinquent() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	missed := 0
	for _, p := range l.Payments {
		if !p {
			missed++
			if missed >= 2 {
				return true
			}
		} else {
			missed = 0
		}
	}
	return false
}
