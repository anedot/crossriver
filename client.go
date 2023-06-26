package crossriver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

const (
	tokenUrl = "https://crbcos-sandbox.auth0.com/oauth/token"
)

func NewClient(id, secret, apiBase, partnerId string) (*Client, error) {
	if id == "" || secret == "" || apiBase == "" || partnerId == "" {
		return nil, errors.New("Id, Secret and ApiBase are required to create a Client")
	}

	return &Client{
		Client:    &http.Client{},
		Id:        id,
		Secret:    secret,
		ApiBase:   apiBase,
		PartnerId: partnerId,
	}, nil
}

func (c *Client) GetToken(ctx context.Context) (*TokenResponse, error) {
	type createTokenRequest struct {
		GrantType string `json:"grant_type"`
		Id        string `json:"client_id"`
		Secret    string `json:"client_secret"`
		Audience  string `json:"audience"`
	}

	req, err := c.NewRequest(
		ctx,
		"POST",
		tokenUrl,
		createTokenRequest{
			Audience:  "https://api.crbcos.com/",
			GrantType: "client_credentials",
			Id:        c.Id,
			Secret:    c.Secret,
		},
	)

	if err != nil {
		return &TokenResponse{}, err
	}

	response := &TokenResponse{}

	c.Send(req, response)

	if response.Token != "" {
		c.Token = response.Token
	}

	return response, err
}

func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {
	req.Header.Set("Authorization", "Bearer "+c.Token)

	// TODO: figure out how to re-fetch token
	// resp, _ := c.Send(req, v)

	// if resp.StatusCode == http.StatusUnauthorized {
	// 	if _, err := c.GetToken(req.Context()); err != nil {
	// 		return err
	// 	}

	// 	req.Header.Set("Authorization", "Bearer "+c.Token)
	// }

	return c.Send(req, v)
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
		data []byte
	)

	// default headers
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{Response: resp}
		data, err = ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, errResp)

		return errResp
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}

func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	return http.NewRequestWithContext(ctx, method, url, buf)
}
