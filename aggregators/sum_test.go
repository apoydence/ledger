package aggregators_test

import (
	"github.com/poy/ledger/aggregators"
	"github.com/poy/ledger/database"
	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum", func() {

	var sum database.Aggregator

	BeforeEach(func() {
		sum = aggregators.NewSum()
	})

	It("Sums all the account values", func() {
		accs := []*transaction.Account{
			{
				Name:  "some-name-1",
				Value: 1234,
			},
			{
				Name:  "some-name-2",
				Value: 5678,
			},
		}

		Expect(sum.Aggregate(accs)).To(Equal("$69.12"))
	})

	It("registers itself with the aggregator store", func() {
		Expect(aggregators.Store()).To(HaveKey("sum"))
	})

})
