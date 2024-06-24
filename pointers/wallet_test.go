package wallet

import "testing"

func TestWallet(t *testing.T){
    assetBalance := func(t testing.TB, wallet Wallet, want BitCoin){
        t.Helper()

        got := wallet.Balance()

        if got != want {
            t.Errorf("got %d but wanted %d", got, want)
        }
    }

    assertError := func(t testing.TB, err error){
        t.Helper()

        if err == nil {
            t.Error("wanted an error but didn't get one")
        }
    }

    t.Run("Deposit bitcoin", func (t *testing.T){
        wallet := Wallet{}

        wallet.Deposit(BitCoin(10))

        assetBalance(t, wallet, BitCoin(10))
    })

    t.Run("Withdraw bitcoin", func (t *testing.T){
        wallet := Wallet{balance: BitCoin(10)}

        wallet.Withdraw(BitCoin(10))

        assetBalance(t, wallet, BitCoin(0)) 

    })

    t.Run("Withdraw insufficient funds", func(t *testing.T){
        wallet := Wallet{balance: BitCoin(10)}

        err := wallet.Withdraw(BitCoin(200))

        assetBalance(t, wallet, BitCoin(10))
        assertError(t, err)
    })
}


