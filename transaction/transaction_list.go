package transaction

type TransactionList []*Transaction

func (t TransactionList) Len() int {
	return len(t)
}

func (t TransactionList) Less(i, j int) bool {
	return t[i].Date.Before(t[j].Date)
}

func (t TransactionList) Swap(i, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
}
