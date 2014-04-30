package erede

import (
	"encoding/xml"
)

type TransactionResponse struct {
	XMLName         xml.Name           `xml:"Response"`
	QueryTxnResult  RespQueryTxnResult `xml:"QueryTxnResult"`
	ExtendedRespMsg string             `xml:"extended_response_message"`
	ExtendedStatus  string             `xml:"extended_status"`
	Mode            string             `xml:"mode"`
	Reason          string             `xml:"reason"`
	Status          int                `xml:"status"`
	Time            int64              `xml:"time"`
}

type RespQueryTxnResult struct {
	Card                 TrRespCard `xml:"Card"`
	Acquirer             string     `xml:"acquirer"`
	AuthHostRef          int        `xml:"auth_host_reference"`
	AuthCode             int        `xml:"authcode"`
	GatewayRef           string     `xml:"gateway_reference"`
	Environment          string     `xml:"environment"`
	FulfillDate          string     `xml:"fulfill_date"`
	FulfillTimestamp     int64      `xml:"fulfill_timestamp"`
	MerchantRef          int        `xml:"merchant_reference"`
	Reason               string     `xml:"reason"`
	Sent                 string     `xml:"sent"`
	Status               int        `xml:"status"`
	TransactionDate      string     `xml:"transaction_date"`
	TransactionTimestamp int64      `xml:"transaction_timestamp"`
}

type TrRespCard struct {
	Category   string `xml:"card_category"`
	Country    string `xml:"country"`
	ExpiryDate string `xml:"expirydate"`
	Issuer     string `xml:"issuer"`
	PAN        string `xml:"pan"`
	Scheme     string `xml:"scheme"`
}
