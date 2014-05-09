package erede

import (
	"encoding/xml"
)

const (
	RStatSuccess                       = 1
	RStatSocketWriteError              = 2
	RStatTimeout                       = 3
	RStatEditError                     = 5
	RStatCommsError                    = 6
	RStatUnauthorized                  = 7
	RStatCurrencyError                 = 9
	RStatAuthError                     = 10
	RStatInvalidAuthCode               = 12
	RStatTypeFieldMissing              = 13
	RStatDBServerError                 = 14
	RStatInvalidType                   = 15
	RStatCannotFulfillTransaction      = 19
	RStatDuplicateTransactionReference = 20
	RStatInvalidCardType               = 21
	RStatInvalidReference              = 22
	RStatExpiryDateInvalid             = 23
	RStatCardExpired                   = 24
	RStatCardNumberInvalid             = 25
	RStatCardNumberWrongLength         = 26
	RStatIssueNumberError              = 27
	RStatStartDateError                = 28
	RStatCardNotValidYet               = 29
	RStatStartDateAfterExpiryDate      = 30
	//TODO: fill more errors
	RStatCurrencyNotSupportedByCard = 59
	RStatInvalidXML                 = 60
	//TODO: fill more errors
	RStatPaymentGatewayBusy = 440
	//TODO: fill more errors
	RStatInvalidTransactionType    = 473
	RStatInvalidValueForMerchantID = 480
)

const (
	RRejectedServiceUnauthorized = 51
	//TODO: fill more errors
	RRejectedInvalidVendor   = 57
	RRejectedUnauthorized    = 58
	RRejectedInvalidPassword = 65
	//TODO: fill more errors
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

func GetRespRejectionDescription(code int) string {
	switch code {
	case 51:
		return "Produto ou serviço não habilitado para o estabelecimento. Entre em contato com a Rede."
	case 53:
		return "Transação não permitida para o emissor. Entre em contato com a Rede."
	case 56:
		return "Erro nos dados informados. Tente novamente."
	case 57:
		return "Estabelecimento inválido."
	case 58:
		return "Transação não autorizada. Contate o emissor."
	case 65:
		return "Senha inválida. Tente novamente."
	case 69:
		return "Transação não permitida para este produto ou serviço."
	case 72:
		return "Contate o emissor."
	case 74:
		return "Falha na comunicação. Tente novamente."
	case 79:
		return "Cartão expirado. Transação não pode ser resubmetida. Contate o emissor."
	case 80:
		return "Transação não autorizada. Contate o emissor. (Saldo Insuficiente)"
	case 81:
		return "Produto ou serviço não habilitado para o emissor (AVS)."
	case 82:
		return "Transação não autorizada para cartão de débito."
	case 83:
		return "Transação não autorizada. Problemas com cartão. Contate o emissor."
	case 84:
		return "Transação não autorizada. Transação não pode ser resubmetida. Contate o emissor."
	}
	return "ERRO!"
}
