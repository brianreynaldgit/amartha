// repo.go
// Simulates a simple in-memory loan repository (like a database).
// Supports loan creation and retrieval by loan ID with concurrency safety.

package loan

import (
	"errors"
	"sync"
)

type LoanRepo struct {
	mu    sync.RWMutex
	loans map[string]*Loan
}

func NewLoanRepo() *LoanRepo {
	return &LoanRepo{
		loans: make(map[string]*Loan),
	}
}

func (r *LoanRepo) CreateLoan(loanID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.loans[loanID]; exists {
		return errors.New("loan already exists")
	}
	r.loans[loanID] = NewLoan()
	return nil
}

func (r *LoanRepo) GetLoan(loanID string) (*Loan, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	loan, exists := r.loans[loanID]
	if !exists {
		return nil, errors.New("loan not found")
	}
	return loan, nil
}
