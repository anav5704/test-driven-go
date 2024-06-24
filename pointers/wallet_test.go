package wallet

import "testing"

func TestWallet(t *testing.T){
    t.Run("Deposit bitcoin", func (t * testing.T){
        wallet := Wallet{}

        wallet.Deposit(BitCoin(10))

        got := wallet.Balance()
        want := BitCoin(10)

        if(got != want){
            t.Errorf("got %d but wanted %d", got, want)
        }
    })
    
    t.Run("Withdraw bitcoin", func (t *testing.T){
        wallet := Wallet{balance: BitCoin(10)}

        wallet.Withdraw(BitCoin(10))

        got := wallet.Balance()
        want := BitCoin(0)

        if(got != want){
            t.Errorf("got %d but wanted %d", got, want)
        }

    })
}


