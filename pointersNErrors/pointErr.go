package pointersnerrors

type Wallet struct {
	mon int
}

func (w *Wallet) Deposit(amount int) {

	w.mon = amount
}
func (w Wallet) Balance() int {
	return w.mon
}
