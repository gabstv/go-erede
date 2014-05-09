package erede

func NewCreditCardTransaction(cardNumber, cvc2, expiresMonth, expiresYear, orderID string, moneyAmount float64) *Transaction {
	if len(expiresYear) > 2 {
		expiresYear = expiresYear[2:] // will become an issue after year 9999
	}
	t := &Transaction{
		CardNumber:       cardNumber,
		CardExpiryDate:   CardExpDate{expiresMonth, expiresYear},
		CVC2:             cvc2,
		OrderID:          orderID,
		OrderMoneyAmount: moneyAmount,
		OrderType:        "credit",
	}
	t.location = "Sao Paulo"
	return t
}
