package database

import (
	"github.com/apoydence/ledger/transaction"
)

type Aggregator interface {
	Aggregate(acc []*transaction.Account) int64
}

type Filter interface {
	Filter(*transaction.Transaction) []*transaction.Account
}

type Database struct {
	transactionList []*transaction.Transaction
}

func New() *Database {
	return new(Database)
}

func (db *Database) Add(ts ...*transaction.Transaction) {
	db.transactionList = append(db.transactionList, ts...)
}

func (db *Database) Aggregate(start, end *transaction.Date, f Filter, aggs ...Aggregator) ([]*transaction.Transaction, []int64) {
	results, accs := db.subQuery(start, end, f)

	var aggResults []int64
	for _, agg := range aggs {
		aggResults = append(aggResults, agg.Aggregate(accs))
	}

	return results, aggResults
}

func (db *Database) Query(start, end *transaction.Date, f Filter) []*transaction.Transaction {
	results, _ := db.subQuery(start, end, f)
	return results
}

func (db *Database) subQuery(start, end *transaction.Date, f Filter) ([]*transaction.Transaction, []*transaction.Account) {
	var ts []*transaction.Transaction
	var accsResults []*transaction.Account

	for _, t := range db.transactionList {
		if !inTimeRange(t.Date, start, end) {
			continue
		}
		accs := filter(t, f)
		if len(accs) > 0 {
			ts = append(ts, t)
			accsResults = append(accsResults, accs...)
		}
	}
	return ts, accsResults
}

func inTimeRange(date, start, end *transaction.Date) bool {
	return date.GreaterThanEqualTo(start) && end.GreaterThanEqualTo(date)
}

func filter(t *transaction.Transaction, f Filter) []*transaction.Account {
	if t.Accounts == nil {
		return nil
	}

	if f == nil {
		return t.Accounts.Accounts
	}

	return f.Filter(t)
}
