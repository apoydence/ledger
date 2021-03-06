package database

import "github.com/poy/ledger/transaction"

type FilterFunc func(*transaction.Transaction) []*transaction.Account

func (f FilterFunc) Filter(t *transaction.Transaction) []*transaction.Account {
	return f(t)
}
