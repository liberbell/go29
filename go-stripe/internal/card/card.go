package card

import "github.com/stripe/stripe-go"

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string
	BankReturnCode      string
}

func (c *Card) CreatePaymentIntent(amount int, currency string) (*stripe.PaymentIntent, error, string) {

}
