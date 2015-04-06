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
	// 3DS
	RStat3DSPayerVerificationRequired = 150
	RStat3DSInvalidTransactionType    = 151
	RStat3DSManualAuthNotSupported    = 152
	RStat3DSVerifyElmtMissing         = 153
	RStat3DSInvalidVerifyValue        = 154
	RStat3DSFieldMissing              = 155
	RStat3DSInvalidBrowserDeviceCateg = 156
	RStat3DSMerchantNotEnabled        = 157
	RStat3DSSchemeNotSupported        = 158
	RStat3DSVerificationFailed        = 159
	RStat3DSInvalidIssuerResponse     = 160
	RStat3DSAuthFailedCallCentre      = 161
	RStat3DSCardNotEnrolled           = 162
	//TODO: fill more errors
	RStatPaymentGatewayBusy = 440
	//TODO: fill more errors
	RStat3DSRequired               = 471
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
	XMLName         xml.Name        `xml:"Response"`
	Card            TrRespCard      `xml:"Card"`
	CardTxn         TrRespCardTxn   `xml:"CardTxn"`
	BoletoTxn       TrRespBoletoTxn `xml:"BoletoTxn"`
	Acquirer        string          `xml:"acquirer"`
	AuthHostRef     int             `xml:"auth_host_reference"`
	GatewayRef      string          `xml:"gateway_reference"`
	ExtendedRespMsg string          `xml:"extended_response_message"`
	ExtendedStatus  string          `xml:"extended_status"`
	MerchantRef     string          `xml:"merchant_reference"`
	MID             string          `xml:"mid"`
	Mode            string          `xml:"mode"`
	Reason          string          `xml:"reason"`
	Status          int             `xml:"status"`
	Time            int64           `xml:"time"`
}

type QueryResponse struct {
	XMLName        xml.Name `xml:"Response"`
	QueryTxnResult QueryResponseTxnResult
	Mode           string `xml:"mode"`
	Reason         string `xml:"reason"`
	Status         int    `xml:"status"`
	Time           int64  `xml:"time"`
}

type QueryResponseTxnResult struct {
	BoletoTxn            TrRespBoletoTxn `xml:"BoletoTxn,omitempty" json:",omitempty"`
	CardTxn              TrRespCardTxn   `xml:"Card,omitempty" json:",omitempty"` //
	Acquirer             string          `xml:"acquirer"`                         //
	AuthHostRef          int             `xml:"auth_host_reference"`              //
	AuthCode             string          `xml:"authcode"`                         //
	Environment          string          `xml:"environment"`                      //
	FulfillDate          string          `xml:"fulfill_date"`                     //
	FulfillTimestamp     string          `xml:"fulfill_timestamp"`                //
	TransactionDate      string          `xml:"transaction_date"`                 //
	TransactionTimestamp string          `xml:"transaction_timestamp"`            //
	Sent                 string          `xml:"sent"`                             //
	GatewayRef           string          `xml:"gateway_reference"`                //
	MerchantRef          string          `xml:"merchant_reference"`               //
	Reason               string          `xml:"reason"`                           //
	Status               int             `xml:"status"`                           //
}

//TODO: CHECK xid, aav, caavAlgorithm, eci
type TrRespThreeDSecure struct {
	AcsURL        string `xml:"acs_url"`
	PareqMessage  string `xml:"pareq_message"`
	XID           string `xml:"xid"`
	AAV           string `xml:"aav"`
	CAVVAlgorithm string `xml:"cavvAlgorithm"`
	ECI           string `xml:"eci"`
}

type TrRespCard struct {
	AccType string `xml:"card_account_type"`
}

type TrRespCardTxn struct {
	AccType      string              `xml:"card_account_type"`
	CardCategory string              `xml:"card_category"`
	Cv2Avs       TrRespCardTxnCv2AVS `xml:"Cv2Avs" json:",omitempty"`
	CardScheme   string              `xml:"scheme"`
	Country      string              `xml:"country"`
	Issuer       string              `xml:"issuer"`
	ThreeDSecure TrRespThreeDSecure  `xml:"ThreeDSecure,omitempty" json:",omitempty"`
	ExpiryDate   string              `xml:"expirydate"` // mm/yy
	PAN          string              `xml:"pan"`
	AuthCode     string              `xml:"authcode"` // não é NSU!
}

type TrRespBoletoTxn struct {
	Method        string `xml:"method"`
	Language      string `xml:"language"`
	Title         string `xml:"title"`
	Country       string `xml:"country"`
	URL           string `xml:"url"`
	TxnStatus     string `xml:"txn_status"`
	BarcodeNumber string `xml:"barcode_number"`
	// query:
	Amount            float64 `xml:"amount"`
	BillingCity       string  `xml:"billing_city"`
	BillingCountry    string  `xml:"billing_country"`
	BillingPostcode   string  `xml:"billing_postcode"`
	BillingStreet1    string  `xml:"billing_street1"`
	BoletoNumber      string  `xml:"boleto_number"`
	BoletoURL         string  `xml:"boleto_url"`
	CustomerEmail     string  `xml:"customer_email"`
	CustomerIP        string  `xml:"customer_ip"`
	CustomerTelephone string  `xml:"customer_telephone"`
	ExpiryDate        string  `xml:"expiry_date"`
	FisrtName         string  `xml:"first_name"`
	LastName          string  `xml:"last_name"`
	Instructions      string  `xml:"instructions"`
	InterestPerDay    float64 `xml:"interest_per_day"`
	MerchantID        string  `xml:"merchant_id"`
	OrderID           string  `xml:"order_id"`
	OverdueFine       float64 `xml:"overdue_fine"`
	PaymentStatus     string  `xml:"payment_status"`
	ProcessorID       string  `xml:"processor_id"`
	TransactionID     string  `xml:"transaction_id"`
}

type TrRespCardTxnCv2AVS struct {
	Status string `xml:"cv2avs_status"`
	Policy int    `xml:"policy"`
}

// type TransactionResponse2 struct {
// 	XMLName         xml.Name           `xml:"Response"`
// 	QueryTxnResult  RespQueryTxnResult `xml:"QueryTxnResult"`
// 	ExtendedRespMsg string             `xml:"extended_response_message"`
// 	ExtendedStatus  string             `xml:"extended_status"`
// 	Mode            string             `xml:"mode"`
// 	Reason          string             `xml:"reason"`
// 	Status          int                `xml:"status"`
// 	Time            int64              `xml:"time"`
// }
//
// type RespQueryTxnResult struct {
// 	Card                 TrRespCard2 `xml:"Card"`
// 	Acquirer             string      `xml:"acquirer"`
// 	AuthHostRef          int         `xml:"auth_host_reference"`
// 	AuthCode             string      `xml:"authcode"`
// 	GatewayRef           string      `xml:"gateway_reference"`
// 	Environment          string      `xml:"environment"`
// 	FulfillDate          string      `xml:"fulfill_date"`
// 	FulfillTimestamp     int64       `xml:"fulfill_timestamp"`
// 	MerchantRef          int         `xml:"merchant_reference"`
// 	Reason               string      `xml:"reason"`
// 	Sent                 string      `xml:"sent"`
// 	Status               int         `xml:"status"`
// 	TransactionDate      string      `xml:"transaction_date"`
// 	TransactionTimestamp int64       `xml:"transaction_timestamp"`
// }

type TrRespCard2 struct {
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

func GetGenRespDescription(code int) string {
	switch code {
	case 1:
		return "Sucesso."
	case 2:
		return "A comunicação foi interrompida"
	case 3:
		return "Ocorreu um timeout enquanto os detalhes da transação eram lidos"
	case 5:
		return "Um campo foi especificado duas vezes. Foram enviados dados excessivos ou inválidos, um fulfill de pré-autorização falhou ou um campo foi omitido. O argumento oferecerá uma melhor indicação do que exatamente deu errado"
	case 6:
		return "Erro no link de comunicação; reenvie"
	case 9:
		return "A moeda especificada não existe"
	case 10:
		return "O vTID ou senha são incorretos"
	case 12:
		return "O código de autorização fornecido é inválido"
	case 13:
		return "Não foi inserido um tipo de transação"
	case 14:
		return "Os detalhes da transação não foram enviados ao nosso banco de dados"
	case 15:
		return "Foi especificado um tipo de transação inválido"
	case 19:
		return "Houve uma tentativa de fulfill de uma transação que não pode ser confirmada ou que já foi confirmada"
	case 20:
		return "Já foi enviada uma transação bem-sucedida que utiliza este vTID e número de referência"
	case 21:
		return "Este terminal não aceita transações para este tipo de cartão"
	case 22:
		return "Os números de referência devem ter 16 dígitos para transações de fulfill, ou de 6 a 30 dígitos para todas as outras"
	case 23:
		return "Expiry date do cartão inválido."
	case 24:
		return "A data de validade fornecida é anterior à data atual"
	case 25, 26:
		return "Número do cartão inválido"
	}
	return ""
}
