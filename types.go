package crossriver

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	createPushRegistrationPath = "/webhooks/v1/registrations"
)

type (
	Client struct {
		mu        sync.Mutex
		ApiBase   string
		Audience  string
		Client    *http.Client
		GrantType string
		Id        string
		Log       io.Writer
		PartnerId string
		Secret    string
		Token     string
	}

	BankAccount struct {
		RoutingNumber  string
		AccountNumber  string
		AccountType    string
		Name           string
		Identification string
	}

	Payment struct {
		AccountNumber   string      `json:"accountNumber"`
		Receiver        BankAccount `json:"receiver"`
		SecCode         string      `json:"secCode"`
		Description     string      `json:"description"`
		TransactionType string      `json:"transactionType"`
		Amount          int         `json:"amount"`
		ServiceType     string      `json:"serviceType"`
	}

	Webhook struct {
		PartnerId    string `json:"partnerId"`
		EventName    string `json:"eventName"`
		CallbackUrl  string `json:"callbackUrl"`
		AuthUsername string `json:"authUsername"`
		AuthPassword string `json:"authPassword"`
		Type         string `json:"type"`
		Format       string `json:"format"`
	}

	AchBatchResponse struct {
		Id                 string   `json:"id"`
		ReferenceId        string   `json:"referenceId"`
		Status             string   `json:"status"`
		AccountNumber      string   `json:"accountNumber"`
		PaymentCount       int      `json:"paymentCount"`
		DebitTotal         int      `json:"debitTotal"`
		CreditTotal        int      `json:"creditTotal"`
		ImportCount        int      `json:"importCount"`
		ProductId          string   `json:"productId"`
		PartnerId          string   `json:"partnerId"`
		CreatedAt          string   `json:"createdAt"`
		ImportedAt         string   `json:"importedAt"`
		LastModifiedAt     string   `json:"lastModifiedAt"`
		PaymentIdentifiers []string `json:"paymentIdentifiers"`
	}

	AchPaymentResponse struct {
		Id              string
		AccountNumber   string
		ReferenceId     string
		PaymentType     string
		Direction       string
		Status          string
		Source          string
		PostingType     string
		PostingCode     string
		Posting         string
		Originator      BankAccount
		Receiver        BankAccount
		SecCode         string
		Description     string
		TransactionType string
		Amount          int
		ServiceType     string
		TraceNumber     string
		CreatedAt       time.Time
		LastModifiedAt  time.Time
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	ErrorResponse struct {
		Response *http.Response `json:"-"`
		Errors   []Error        `json:"errors"`
	}

	TokenResponse struct {
		Token     string `json:"access_token"`
		ExpiresIn int    `json:"expires_in"`
		TokenType string `json:"token_type"`
	}

	WebhookResponse struct {
		PartnerId    string `json:"partnerId"`
		EventName    string `json:"eventName"`
		CallbackUrl  string `json:"callbackUrl"`
		AuthUsername string `json:"authUsername"`
		AuthPassword string `json:"authPassword"`
		Type         string `json:"type"`
		Format       string `json:"format"`
	}
)

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Errors)
}
