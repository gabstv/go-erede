package erede

const (
	OrderTypeCredit = "credit"
)

type Transaction struct {
	CardNumber       string
	CardExpiryDate   CardExpDate
	CVC2             string
	OrderID          string
	OrderMoneyAmount float64
	OrderType        string
	ShippingAddress  TransactionAddress
	BillingAddress   TransactionAddress
	Products         []TransactionProduct
}

type CardExpDate struct {
	Month string
	Year  string
}

type TransactionBuyer struct {
	Location string
	BuyerID  string
	Email    string
	IPv4     string
	Phone1   string
	Phone2   string
	DOB      TransactionDate
	RG       string
	Name     string
	Surname  string
	Address  TransactionAddress
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
