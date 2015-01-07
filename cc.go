package erede

// Deprecated: Use Webservice.NewCreditCardTransaction
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

// Deprecated: Use Webservice.NewDebitCardTransaction
func NewDebitCardTransaction(cardNumber, cvc2, expiresMonth, expiresYear, orderID string, moneyAmount float64, debitDescription, debitStoreURL string) *Transaction {
	if len(expiresYear) > 2 {
		expiresYear = expiresYear[2:] // will become an issue after year 9999
	}
	t := &Transaction{
		CardNumber:       cardNumber,
		CardExpiryDate:   CardExpDate{expiresMonth, expiresYear},
		CVC2:             cvc2,
		OrderID:          orderID,
		OrderMoneyAmount: moneyAmount,
		OrderType:        "debit",
	}
	t.location = "Sao Paulo"
	t.DebitInfo.DateTime = DTNow()
	t.DebitInfo.Description = debitDescription
	t.DebitInfo.URL = debitStoreURL
	t.skipRisk = true
	return t
}

// Deprecated: Use Webservice.NewBoletoTransaction
func NewBoletoTransaction(orderID string, moneyAmount float64, details TransactionBoletoInfo) *Transaction {
	t := &Transaction{
		OrderID:          orderID,
		OrderMoneyAmount: moneyAmount,
		OrderType:        "boleto",
		BoletoInfo:       details,
	}
	return t
}
