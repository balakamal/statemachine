package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Billing struct {
}

type PaymentSource struct {
	Id             string  `json:"id"`
	Type           string  `json:"type"`
	BillingAddress Billing `json:"billing_address"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	Scheme         string  `json:"scheme"`
	Last_4         string  `json:"last_4"`
	Fingerprint    string  `json:"fingerprint"`
	Bin            string  `json:"bin"`
	CardType       string  `json:"card_type"`
	CartCategory   string  `json:"cart_category"`
	Issuer         string  `json:"issuer"`
	IssuerCountry  string  `json:"issuer_country"`
	ProductId      string  `json:"product_id"`
	ProductType    string  `json:"product_type"`
	AvsCheck       string  `json:"avs_check"`
	CvvCheck       string  `json:"cvv_check"`
}

type Customer struct {
	Id string `json:"id"`
}

type PaymentProcessing struct {
	AcquirerTransactionId      string `json:"acquirer_transaction_id"`
	Retrieval_reference_number string `json:"retrieval_reference_number"`
}

type PaymentMetadata struct {
}

type PaymentRisk struct {
	Flagged bool `json:"flagged"`
}

type PaymentData struct {
	ActionId        string            `json:"action_id"`
	PaymentType     string            `json:"payment_type"`
	AuthCode        string            `json:"auth_code"`
	ResponseCode    string            `json:"response_code"`
	ResponseSummary string            `json:"response_summary"`
	SchemeId        string            `json:"scheme_id"`
	Source          PaymentSource     `json:"source"`
	Customers       Customer          `json:"customers"`
	Processing      PaymentProcessing `json:"processing"`
	Amount          int               `json:"amount"`
	Metadata        PaymentMetadata   `json:"metadata"`
	Risk            PaymentRisk       `json:"risk"`
	Id              string            `json:"id"`
	Currency        string            `json:"currency"`
	ProcessedOn     string            `json:"processed_on"`
	Reference       string            `json:"reference"`
}

type Self struct {
	href string `json:"href"`
}

type Payment struct {
	href string `json:"href"`
}

type Link struct {
	Self     Self    `json:"self"`
	Payments Payment `json:"payments"`
}
type Checkout struct {
	Id        string      `json:"id"`
	Type      string      `json:"type"`
	CreatedOn string      `json:"created_on"`
	Data      PaymentData `json:"data"`
	Links     Link        `json:"_links"`
}

func main() {
	file, _ := ioutil.ReadFile("/mappingcheckout/payment_approved.json")

	data := Checkout{}

	_ = json.Unmarshal([]byte(file), &data)
	fmt.Print(data)
}
