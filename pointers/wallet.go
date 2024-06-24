package wallet

import "errors"

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.balance {
		return errors.New("Insufficient funds")
	}

	w.balance -= amount
	return nil
}
