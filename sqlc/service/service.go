package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/vrublevski/sqlc/query"
)

type Accountant struct {
	db *sql.DB
}

func NewAccountant(db *sql.DB) *Accountant {
	return &Accountant{db: db}
}

func (a *Accountant) DeductFromAccounts(amount int, openedTime time.Time) ([]query.BankAccount, error) {
	queries := query.New(a.db)

	tx, err := a.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("unable to start a transaction:%v", err)
	}
	defer tx.Rollback()

	var frozenAccounts []query.BankAccount
	queries = queries.WithTx(tx)
	stmts := []func() error{
		func() error {
			return queries.DeductFromAccount(context.Background(), query.DeductFromAccountParams{
				Balance:   int32(amount),
				Balance_2: int32(amount),
				Status:    query.AccountStatusRegular,
				Opened:    openedTime,
			})
		},
		func() error {
			return queries.FreezeAccounts(context.Background(), int32(amount))
		},
		func() error {
			frozenAccounts, err = queries.ListAccounts(context.Background(), query.AccountStatusFrozen)
			return err
		},
	}

	for _, stmt := range stmts {
		if err := stmt(); err != nil {
			return nil, fmt.Errorf("running deduct from accounts transaction:%v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("unable to commit accounts transaction:%v", err)
	}

	return frozenAccounts, nil
}
