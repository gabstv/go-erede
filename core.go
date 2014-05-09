package erede

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

var (
	User     string
	Password string
	URL      = "https://scommerce.redecard.com.br/Beta/wsTransaction"
)

const (
	OrderTypeCredit = "credit"
)

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
	skipRisk            bool
	skipAddressDetails  bool
	skipShippingDetails bool
	skipBillingDetails  bool
	skipOrderDetails    bool
	skipLineItems       bool
	location            string
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

func (t *Transaction) Submit() (*TransactionResponse, error) {
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
	vals["Method"] = "pre"
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

	// </Transaction>
	tpl := template.New("xml")
	tpl = template.Must(tpl.Parse(tpl_request_cc))
	var buffer bytes.Buffer
	tpl.Execute(&buffer, vals)
	fmt.Println(buffer.String())
	resp, err := http.Post(URL, "application/xml", &buffer)
	if err != nil {
		return nil, err
	}
	tr := &TransactionResponse{}
	buffer.Truncate(0)
	defer resp.Body.Close()
	io.Copy(&buffer, resp.Body)
	err = xml.Unmarshal(buffer.Bytes(), tr)
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
	return t.Year + "-" + t.Month + "-" + t.Day
}
