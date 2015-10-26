package erede

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type Webservice struct {
	user             string
	password         string
	environment      int
	verbose          bool
	XMLRequestLogger io.Writer
	XMLResultLogger  io.Writer
}

func NewWebservice(user, pw string, env int) *Webservice {
	return &Webservice{user, pw, env, false, nil, nil}
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

func (ws *Webservice) ConsultBoleto(gatewayRef string, reqccbuf, respccbuf io.Writer) (*QueryResponse, error) {
	tpl := new(bytes.Buffer)
	tpl.WriteString(`<Request version="2">`)
	tpl.WriteString("<Authentication><AcquirerCode><rdcd_pv>")
	tpl.WriteString(ws.user)
	tpl.WriteString("</rdcd_pv></AcquirerCode><password>")
	tpl.WriteString(ws.password)
	tpl.WriteString("</password></Authentication>")
	tpl.WriteString("<Transaction><HistoricTxn><reference>")
	tpl.WriteString(gatewayRef)
	tpl.WriteString("</reference>")
	tpl.WriteString("<method>query</method></HistoricTxn></Transaction></Request>")
	if Verbose {
		log.Println("Sending ConsultBoleto: ", tpl.String())
	}
	if reqccbuf != nil {
		reqccbuf.Write(tpl.Bytes())
	}
	resp, err := http.Post(URL(), "application/xml", tpl)
	if err != nil {
		return nil, err
	}
	tpl.Reset()
	defer resp.Body.Close()
	io.Copy(tpl, resp.Body)
	if Verbose {
		log.Println("XML RESPONSE", tpl.String())
	}
	if respccbuf != nil {
		respccbuf.Write(tpl.Bytes())
	}
	tr := &QueryResponse{}
	err = xml.Unmarshal(tpl.Bytes(), tr)
	return tr, err
}

func (ws *Webservice) NewQueryTransaction(gatewayref string) *QueryTx {
	qtx := &QueryTx{
		GatewayReference: gatewayref,
	}
	qtx.ws = ws
	return qtx
}

//TODO: proper fulfill and pre methods
