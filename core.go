package erede

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

var (
	User        string
	Password    string
	Environment = 1
	Verbose     = false
)

const (
	OrderTypeCredit = "credit"
	URLDEV          = "https://scommerce.redecard.com.br/Beta/wsTransaction"
	URLPROD         = "https://ecommerce.userede.com.br/Transaction/wsTransaction"
)

const (
	DevelopmentEnv = 0
	ProductionEnv  = 1
)

func SetProductionEnv() {
	Environment = ProductionEnv
}

func SetDevelopmentEnv() {
	Environment = DevelopmentEnv
}

func URL() string {
	if Environment == ProductionEnv {
		return URLPROD
	}
	return URLDEV
}

type Transaction struct {
	CardNumber          string
	CardExpiryDate      CardExpDate
	CVC2                string
	OrderID             string
	OrderMoneyAmount    float64
	OrderType           string
	Installments        int
	DeliveryDetails     TransactionDelivery
	BillingAddress      TransactionAddress
	Products            []TransactionProduct
	Buyer               TransactionBuyer
	DebitInfo           TransactionDebitInfo
	BoletoInfo          TransactionBoletoInfo
	ThreeDSecureData    ThreeDSecure
	skipRisk            bool
	threedSecure        bool
	skipAddressDetails  bool
	skipShippingDetails bool
	skipBillingDetails  bool
	skipOrderDetails    bool
	skipLineItems       bool
	location            string
}

type TransactionDebitInfo struct {
	DateTime    TransactionDateTime
	Description string
	URL         string
}

type TransactionBoletoInfo struct {
	Email       string
	Nome        string
	Sobrenome   string
	Endereco    string
	Cidade      string
	CEP         string
	Telefone    string
	JurosDiaPct float64
	MultaPct    float64
	Vencimento  TransactionDate
	BancoID     string
	Instrucoes  string
}

type ThreeDSecure struct {
	MobileNumber     string
	MerchantURL      string
	PurchaseDesc     string
	PurchaseDatetime string
	Browser          ThreeDSBrowser
}

func (t *ThreeDSecure) SetDatetime(t0 time.Time) {
	t.PurchaseDatetime = fmt.Sprintf("%04d%02d%02d %02d:%02d:%02d", t0.Year(), int(t0.Month()), t0.Day(), t0.Hour(), t0.Minute(), t0.Second())
}

type ThreeDSBrowser struct {
	DeviceCategory int
	UserAgent      string
}

type CardExpDate struct {
	Month string
	Year  string
}

type TransactionBuyer struct {
	BuyerID string
	Email   string
	IPv4    string
	Phone1  string
	Phone2  string
	DOB     TransactionDate
	RG      string
	Name    string
	Surname string
	Address TransactionAddress
}

func DTNow() TransactionDateTime {
	now := time.Now()
	v := TransactionDateTime{}
	v.Date.Year = fmt.Sprintf("%04d", now.Year())
	v.Date.Month = fmt.Sprintf("%02d", int(now.Month()))
	v.Date.Day = fmt.Sprintf("%02d", now.Day())
	v.Time.Hour = fmt.Sprintf("%02d", now.Hour())
	v.Time.Minute = fmt.Sprintf("%02d", now.Minute())
	v.Time.Second = fmt.Sprintf("%02d", now.Hour())
	return v
}

func (dt *TransactionDateTime) Sprint() string {
	return fmt.Sprintf("%s%s%s %s:%s:%s", dt.Date.Year, dt.Date.Month, dt.Date.Day, dt.Time.Hour, dt.Time.Minute, dt.Time.Second)
}

type TransactionDateTime struct {
	Date TransactionDate
	Time TransactionTime
}

type TransactionTime struct {
	Hour   string
	Minute string
	Second string
}

type TransactionDate struct {
	Year  string
	Month string
	Day   string
}

type TransactionAddress struct {
	Line1   string
	Line2   string
	City    string
	State   string
	Country string
	ZIP     string
}

type TransactionDelivery struct {
	Date         TransactionDate
	DeliveryKind string // SEDEX 10
	Address      TransactionAddress
}

type TransactionProduct struct {
	ProductId   string
	Description string
	Category    string
	Quantity    int
	MoneyAmount float64
	Risk        string // High | Low
}

/* ~~~~~ Transaction ~~~~~ */

func (t *Transaction) AddProduct(id, description, category string, quantity int, moneyAmount float64) *TransactionProduct {
	t.skipLineItems = false
	i0 := TransactionProduct{id, description, category, quantity, moneyAmount, "low"}
	t.Products = append(t.Products, i0)
	return &t.Products[len(t.Products)-1]
}

func (t *Transaction) SetVendorLocation(loc string) *Transaction {
	t.location = loc
	return t
}

func (t *Transaction) SetSkipAddressDetails(val bool) *Transaction {
	t.skipAddressDetails = val
	return t
}

func (t *Transaction) SetSkipRisk(val bool) *Transaction {
	t.skipRisk = val
	return t
}

func (t *Transaction) SetSkipShippingDetails(val bool) *Transaction {
	t.skipShippingDetails = val
	return t
}

func (t *Transaction) SetSkipBillingDetails(val bool) *Transaction {
	t.skipBillingDetails = val
	return t
}

func (t *Transaction) SetUseThreeDSecure(val bool) *Transaction {
	t.threedSecure = val
	return t
}

func FulFillTxn(gatewayReference, authcode string) string {
	vals := make(map[string]interface{}, 0)
	vals["User"] = User
	vals["Password"] = Password
	vals["GatewayReference"] = gatewayReference
	vals["AuthCode"] = authcode

	tpl := template.New("xml")
	tpl = template.Must(tpl.Parse(tpl_fulfill))
	var buffer bytes.Buffer
	tpl.Execute(&buffer, vals)
	log.Println(URL())
	log.Println(buffer.String())
	resp, err := http.Post(URL(), "application/xml", &buffer)
	if err != nil {
		return err.Error()
		//return nil, err
	}
	//tr := &TransactionResponse{}
	buffer.Truncate(0)
	defer resp.Body.Close()
	io.Copy(&buffer, resp.Body)
	if Verbose {
		log.Println(buffer.String())
	}

	return buffer.String()
	//err = xml.Unmarshal(buffer.Bytes(), tr)
	//return tr, err
}

func ConfirmDebitTxn(gatewayRef, debitPaRes string, reqccbuf, respccbuf io.Writer) (*TransactionResponse, error) {
	tpl := new(bytes.Buffer)
	tpl.WriteString(`<Request version="2">`)
	tpl.WriteString("<Authentication><AcquirerCode><rdcd_pv>")
	tpl.WriteString(User)
	tpl.WriteString("</rdcd_pv></AcquirerCode><password>")
	tpl.WriteString(Password)
	tpl.WriteString("</password></Authentication>")
	tpl.WriteString("<Transaction><HistoricTxn><reference>")
	tpl.WriteString(gatewayRef)
	tpl.WriteString("</reference>")
	tpl.WriteString("<method tx_status_u=\"accept\">threedsecure_authorization_request</method>")
	tpl.WriteString("<pares_message>")
	tpl.WriteString(debitPaRes)
	tpl.WriteString("</pares_message>")
	tpl.WriteString("</HistoricTxn></Transaction></Request>")
	if Verbose {
		log.Println("Sending ConfirmDebitTxn: ", tpl.String())
	}
	if reqccbuf != nil {
		reqccbuf.Write(tpl.Bytes())
	}
	resp, err := http.Post(URL(), "application/xml", tpl)
	if err != nil {
		return nil, err
	}
	tpl.Truncate(0)
	defer resp.Body.Close()
	io.Copy(tpl, resp.Body)
	if Verbose {
		log.Println("XML RESPONSE", tpl.String())
	}
	if respccbuf != nil {
		respccbuf.Write(tpl.Bytes())
	}
	tr := &TransactionResponse{}
	err = xml.Unmarshal(tpl.Bytes(), tr)
	return tr, err
}

func ConsultBoleto(gatewayRef string, reqccbuf, respccbuf io.Writer) (*TransactionResponse, error) {
	tpl := new(bytes.Buffer)
	tpl.WriteString(`<Request version="2">`)
	tpl.WriteString("<Authentication><AcquirerCode><rdcd_pv>")
	tpl.WriteString(User)
	tpl.WriteString("</rdcd_pv></AcquirerCode><password>")
	tpl.WriteString(Password)
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
	tpl.Truncate(0)
	defer resp.Body.Close()
	io.Copy(tpl, resp.Body)
	if Verbose {
		log.Println("XML RESPONSE", tpl.String())
	}
	if respccbuf != nil {
		respccbuf.Write(tpl.Bytes())
	}
	tr := &TransactionResponse{}
	err = xml.Unmarshal(tpl.Bytes(), tr)
	return tr, err
}

func (t *Transaction) getXMLBoleto() *bytes.Buffer {
	vals := make(map[string]interface{}, 0)
	vals["User"] = User
	vals["Password"] = Password
	fillif := func(key, val string) {
		if len(val) > 0 {
			vals[key] = val
		}
	}
	fillif("BoletoNome", t.BoletoInfo.Nome)
	fillif("BoletoSobrenome", t.BoletoInfo.Sobrenome)
	fillif("BoletoEndereco", t.BoletoInfo.Endereco)
	fillif("BoletoCidade", t.BoletoInfo.Cidade)
	fillif("BoletoCEP", t.BoletoInfo.CEP)
	fillif("BoletoTel", t.BoletoInfo.Telefone)
	vals["BoletoJurosDia"] = fmt.Sprintf("%v", t.BoletoInfo.JurosDiaPct)
	vals["BoletoMulta"] = fmt.Sprintf("%v", t.BoletoInfo.MultaPct)
	vals["BoletoVencimento"] = t.BoletoInfo.Vencimento.String()
	vals["BoletoBanco"] = t.BoletoInfo.BancoID
	vals["BoletoInstrucoes"] = t.BoletoInfo.Instrucoes
	vals["MReference"] = t.OrderID
	vals["Amount"] = fmt.Sprintf("%v", t.OrderMoneyAmount)

	tpl := template.New("xml")
	tpl = template.Must(tpl.Parse(tpl_request_boleto))
	buffer := new(bytes.Buffer)
	tpl.Execute(buffer, vals)
	return buffer
}

func (t *Transaction) GetXML() *bytes.Buffer {
	if t.OrderType == "boleto" {
		return t.getXMLBoleto()
	}
	if t.skipLineItems && t.skipBillingDetails {
		t.skipOrderDetails = true
	} else {
		t.skipOrderDetails = false
	}
	vals := make(map[string]interface{}, 0)
	vals["User"] = User
	vals["Password"] = Password
	// <Transaction>
	vals["CC"] = t.CardNumber
	vals["ExpiryDate"] = t.CardExpiryDate.String()
	vals["CCVC2"] = t.CVC2
	vals["TipoTransacao"] = t.OrderType
	vals["Method"] = "auth" // pre ou auth // mostra se é pré autorização ou não!
	vals["MReference"] = t.OrderID
	vals["Amount"] = fmt.Sprintf("%v", t.OrderMoneyAmount)
	//
	vals["VMLocation"] = t.location
	vals["VChannel"] = "W" // ???
	vals["VuserId"] = t.Buyer.BuyerID
	vals["Vmail"] = t.Buyer.Email
	vals["VIP"] = t.Buyer.IPv4
	vals["PFone1"] = t.Buyer.Phone1
	vals["PFone2"] = t.Buyer.Phone2
	vals["PDateBirth"] = t.Buyer.DOB.String()
	vals["PIDNumber"] = t.Buyer.RG
	vals["PIDType"] = "2" // RG ?
	vals["PFName"] = t.Buyer.Name
	vals["PLName"] = t.Buyer.Surname
	//
	vals["Address1"] = t.Buyer.Address.Line1
	vals["Address2"] = t.Buyer.Address.Line2
	vals["AddressCity"] = t.Buyer.Address.City
	vals["AddressState"] = t.Buyer.Address.State
	vals["AddressCountry"] = t.Buyer.Address.Country
	vals["AddressZIP"] = t.Buyer.Address.ZIP
	//
	vals["DeliveryDate"] = t.DeliveryDetails.Date.String()
	vals["DeliveryMethod"] = t.DeliveryDetails.DeliveryKind
	vals["DeliveryAddress1"] = t.DeliveryDetails.Address.Line1
	vals["DeliveryAddress2"] = t.DeliveryDetails.Address.Line2
	vals["DeliveryCity"] = t.DeliveryDetails.Address.City
	vals["DeliveryState"] = t.DeliveryDetails.Address.State
	vals["DeliveryCountry"] = t.DeliveryDetails.Address.Country
	vals["DeliveryZIP"] = t.DeliveryDetails.Address.ZIP
	//
	vals["Cobranca1"] = t.BillingAddress.Line1
	vals["Cobranca2"] = t.BillingAddress.Line2
	vals["CobrancaCity"] = t.BillingAddress.City
	vals["CobrancaState"] = t.BillingAddress.State
	vals["CobrancaCountry"] = t.BillingAddress.Country
	vals["CobrancaZIP"] = t.BillingAddress.ZIP

	if t.Installments > 1 {
		//TODO: support interest
		vals["TipoParcelado"] = "zero_interest"
		vals["NumParcelas"] = t.Installments
	}

	//
	// items
	i0 := make([]map[string]interface{}, len(t.Products))
	for k, v := range t.Products {
		i0[k] = make(map[string]interface{})
		i0[k]["ProductCode"] = v.ProductId
		i0[k]["ProductDescription"] = v.Description
		i0[k]["ProductCategory"] = v.Category
		i0[k]["OrderQuantity"] = fmt.Sprintf("%v", v.Quantity)
		i0[k]["UnitPrice"] = fmt.Sprintf("%v", v.MoneyAmount)
		i0[k]["ProductRisk"] = v.Risk
	}
	vals["OrderItems"] = i0

	vals["SkipRisk"] = t.skipRisk
	vals["SkipAddressDetails"] = t.skipAddressDetails
	vals["SkipBillingDetails"] = t.skipBillingDetails
	vals["SkipLineItems"] = t.skipLineItems
	vals["SkipOrderDetails"] = t.skipOrderDetails
	vals["SkipShippingDetails"] = t.skipShippingDetails

	vals["UseThreeDSecure"] = t.threedSecure
	vals["ThreeDSecure"] = t.ThreeDSecureData

	// </Transaction>
	tpl := template.New("xml")
	if t.OrderType == "debit" {
		vals["ThreeDSecure"] = struct {
			MerchantURL      string
			PurchaseDesc     string
			PurchaseDatetime string
		}{
			t.DebitInfo.URL,
			t.DebitInfo.Description,
			t.DebitInfo.DateTime.Sprint(),
		}
		tpl = template.Must(tpl.Parse(tpl_request_dc))
	} else {
		tpl = template.Must(tpl.Parse(tpl_request_cc))
	}
	buffer := new(bytes.Buffer)
	tpl.Execute(buffer, vals)
	return buffer
}

func (t *Transaction) Submit(xmlw ...io.Writer) (*TransactionResponse, error) {
	buffer := t.GetXML()
	if Verbose {
		log.Println("XML TO SUBMIT")
		log.Println(URL())
		log.Println(buffer.String())
	}
	resp, err := http.Post(URL(), "application/xml", buffer)
	if err != nil {
		return nil, err
	}
	tr := &TransactionResponse{}
	buffer.Truncate(0)
	defer resp.Body.Close()
	io.Copy(buffer, resp.Body)
	bs := buffer.Bytes()
	if Verbose {
		log.Println(string(bs))
	}
	if len(xmlw) > 0 {
		xmlw[0].Write(bs)
	}
	err = xml.Unmarshal(bs, tr)
	return tr, err
}

/* ~~~~~ TransactionProduct ~~~~~ */

func (t *TransactionProduct) SetHighRisk() *TransactionProduct {
	t.Risk = "High"
	return t
}

func (t *TransactionProduct) SetMediumRisk() *TransactionProduct {
	t.Risk = "Medium"
	return t
}

func (t *TransactionProduct) SetLowRisk() *TransactionProduct {
	t.Risk = "Low"
	return t
}

/* ~~~~~ CardExpDate ~~~~~ */

func (c *CardExpDate) String() string {
	return c.Month + "/" + c.Year
}

/* ~~~~~ TransactionDate ~~~~~ */

func (t *TransactionDate) String() string {
	if len(t.Year) < 1 || len(t.Month) < 1 || len(t.Day) < 1 {
		return ""
	}
	for len(t.Year) < 4 {
		t.Year = "0" + t.Year
	}
	for len(t.Month) < 2 {
		t.Month = "0" + t.Month
	}
	for len(t.Day) < 2 {
		t.Day = "0" + t.Day
	}
	return t.Year + "-" + t.Month + "-" + t.Day
}
