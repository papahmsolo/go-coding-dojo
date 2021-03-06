// Code generated by sqlc. DO NOT EDIT.

package query

import (
	"fmt"
	"time"
)

type AccountStatus string

const (
	AccountStatusFrozen  AccountStatus = "frozen"
	AccountStatusAudited AccountStatus = "audited"
	AccountStatusRegular AccountStatus = "regular"
)

func (e *AccountStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AccountStatus(s)
	case string:
		*e = AccountStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for AccountStatus: %T", src)
	}
	return nil
}

type BankAccount struct {
	ID      int32
	Balance int32
	Status  AccountStatus
	Opened  time.Time
}
