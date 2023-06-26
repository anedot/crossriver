package crossriver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const createWebhookPath = "/webhooks/v1/registrations"

// POST /webhooks/v1/registrations
func (c *Client) CreateWebhook(ctx context.Context, webhook Webhook) (*WebhookResponse, error) {
	// set PartnerId from client configuration
	webhook.PartnerId = c.PartnerId
	createWebhookRequest, err := json.Marshal(webhook)

	req, err := c.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.ApiBase, createWebhookPath),
		createWebhookRequest,
	)

	response := &WebhookResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)

	return response, err
}
