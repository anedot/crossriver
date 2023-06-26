package crossriver

import (
	"context"
	"fmt"
	"net/http"
)

const (
	achBatchesPath = "/ach/v1/client-batches"
)

// GET ach/v1/client-batches/:batch_id
func (c *Client) GetAchBatch(ctx context.Context, batchId string) (*AchBatchResponse, error) {
	req, err := c.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s%s", c.ApiBase, achBatchesPath+"/", batchId),
		nil,
	)

	response := &AchBatchResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)
	return response, err
}

// POST ach/v1/client-batches
func (c *Client) CreateAchBatch(ctx context.Context, payments []Payment) (*AchBatchResponse, error) {
	type createAchBatchRequest struct {
		Payments []Payment `json:"payments"`
	}

	req, err := c.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", c.ApiBase, achBatchesPath),
		createAchBatchRequest{Payments: payments},
	)

	response := &AchBatchResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)
	return response, err
}

// POST ach/v1/client-batches/:batch_id/cancel
func (c *Client) CancelAchBatch(ctx context.Context, batchId string) (*AchBatchResponse, error) {
	req, err := c.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s%s%s", c.ApiBase, achBatchesPath+"/", batchId, "/cancel"),
		nil,
	)

	response := &AchBatchResponse{}

	if err != nil {
		return response, err
	}

	err = c.SendWithAuth(req, response)
	return response, err
}
