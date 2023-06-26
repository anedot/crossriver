package crossriver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	achPaymentsPath = "/ach/v1/payments"
)

// GET ach/v1/payments/:id
func (c *Client) GetAchPayment(ctx context.Context, paymentId string) (*AchPaymentResponse, error) {
	req, err := c.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.ApiBase, achPaymentsPath+"/", paymentId),
		nil,
	)

	response := &AchPaymentResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)

	return response, err
}

// POST ach/v1/payments
func (c *Client) CreateAchPayment(ctx context.Context, payment Payment) (*AchPaymentResponse, error) {
	createAchPaymentRequest, err := json.Marshal(payment)

	req, err := c.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.ApiBase, achPaymentsPath),
		createAchPaymentRequest,
	)

	response := &AchPaymentResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)

	return response, err
}

// POST ach/v1/payments/:id/cancel
func (c *Client) CancelAchPayment(ctx context.Context, paymentId string) (*AchPaymentResponse, error) {
	req, err := c.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s", c.ApiBase, achPaymentsPath+"/", paymentId, "/cancel"),
		nil,
	)

	response := &AchPaymentResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)

	return response, err
}
