package erede

type Webservice struct {
	user        string
	password    string
	environment int
	verbose     bool
}

func NewWebservice(user, pw string, env int) *Webservice {
	return &Webservice{user, pw, env, false}
}

func (ws *Webservice) URL() string {
	if ws.environment == ProductionEnv {
		return URLPROD
	}
	return URLDEV
}

func (ws *Webservice) Verbose(tf bool) {
	ws.verbose = tf
}

func (ws *Webservice) NewCreditCardTransaction(cardNumber, cvc2, expiresMonth, expiresYear, orderID string, moneyAmount float64) *Transaction {
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
	t.ws = ws
	t.location = "Sao Paulo"
	return t
}

func (ws *Webservice) NewDebitCardTransaction(cardNumber, cvc2, expiresMonth, expiresYear, orderID string, moneyAmount float64, debitDescription, debitStoreURL string) *Transaction {
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
	t.ws = ws
	return t
}

func (ws *Webservice) NewBoletoTransaction(orderID string, moneyAmount float64, details TransactionBoletoInfo) *Transaction {
	t := &Transaction{
		OrderID:          orderID,
		OrderMoneyAmount: moneyAmount,
		OrderType:        "boleto",
		BoletoInfo:       details,
	}
	t.ws = ws
	return t
}

//TODO: proper fulfill and pre methods
