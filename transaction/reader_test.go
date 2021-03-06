package transaction_test

import (
	"strings"

	"github.com/poy/ledger/transaction"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reader", func() {

	var reader *transaction.Reader

	BeforeEach(func() {
		strReader := strings.NewReader(`
2015/10/12 Exxon
		Expenses:Auto:Gas         $10.00
		Liabilities:MasterCard   $-10.00
    

2015/10/12 Qdoba
		Expenses:Food:FastFood $21.50
		Liabilities:AmericanExpress $-21.50
`)
		reader = transaction.NewReader(strReader)
	})

	It("reads all the transactions", func() {
		t1, err := reader.Next()
		Expect(err).ToNot(HaveOccurred())
		t2, err := reader.Next()
		Expect(err).ToNot(HaveOccurred())

		Expect(t1).To(Equal(
			&transaction.Transaction{
				Date: transaction.NewDate(2015, 10, 12),
				Title: &transaction.Title{
					Value: "Exxon",
				},
				Accounts: &transaction.AccountList{
					Accounts: []*transaction.Account{
						{
							Name:  "Expenses:Auto:Gas",
							Value: 1000,
						},
						{
							Name:  "Liabilities:MasterCard",
							Value: -1000,
						},
					},
				},
			}))

		Expect(t2).To(Equal(
			&transaction.Transaction{
				Date: transaction.NewDate(2015, 10, 12),
				Title: &transaction.Title{
					Value: "Qdoba",
				},
				Accounts: &transaction.AccountList{
					Accounts: []*transaction.Account{
						{
							Name:  "Expenses:Food:FastFood",
							Value: 2150,
						},
						{
							Name:  "Liabilities:AmericanExpress",
							Value: -2150,
						},
					},
				},
			}))
	})

	Context("invalid transaction", func() {
		BeforeEach(func() {
			strReader := strings.NewReader(`
2015/10/12 Qdoba
		Expenses:Food:FastFood $21.50
		Liabilities:AmericanExpress $-21.50

invalid
`)
			reader = transaction.NewReader(strReader)
		})

		It("reports the line number with an error", func() {
			_, err := reader.Next()
			Expect(err).ToNot(HaveOccurred())
			_, err = reader.Next()
			Expect(err).To(HaveOccurred())
			Expect(err.Line).To(BeEquivalentTo(5))
		})
	})

})
