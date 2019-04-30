package bestmixer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// UserAgent custom set when we send http request
	UserAgent = "gobestmixer/1.0.0"
	// APIURL is base URL for BestMixer
	APIURL = "https://bestmixer.io/api/ext"
)

// Coin is cryptocurrency key
type Coin string

// String returns coin string representation
func (c Coin) String() string {
	return string(c)
}

const (
	// CoinBTC is bitcoin coin
	CoinBTC Coin = "btc"
	// CoinBCH is bitcoin cash coin
	CoinBCH Coin = "bch"
	// CoinLTC is litecoin coin
	CoinLTC Coin = "ltc"
	// CoinETH is ethereum coin
	CoinETH Coin = "eth"
)

// Client manages communication with the BestMixer API.
type Client struct {
	httpClient *http.Client
	apiKey     string

	Code  CodeService
	Fee   FeeService
	Order OrderService
}

// Response from API, general usage for checking the error and parsing data
type Response struct {
	Error error       `json:"error"`
	Data  interface{} `json:"data"`
}

// New creates a new BestMixer client with api key
// you can also setup your own http client if needed,
// otherwise it defaults to http.DefaultClient
func New(apiKey string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	cli := &Client{
		apiKey:     apiKey,
		httpClient: client,
	}

	cli.Code = &CodeServiceOp{client: cli}
	cli.Fee = &FeeServiceOp{client: cli}
	cli.Order = &OrderServiceOp{client: cli}
	return cli
}

// NewRequest creates a new http request for BestMixer
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(rBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)
	return req, nil
}

// Do sends an API request and populates the given interface with the parsed response
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("api responded with http status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// CreateAndDo performs a web request to BestMixer with the given method, path,
// body as data and result resource.
func (c *Client) CreateAndDo(method, path string, data, resource interface{}) error {
	req, err := c.NewRequest(method, path, data)
	if err != nil {
		return err
	}

	resp := Response{
		Data: resource,
	}
	err = c.Do(req, &resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// Post performs a POST request for the given path and saves the result in the
// given resource.
func (c *Client) Post(path string, data, resource interface{}) error {
	return c.CreateAndDo("POST", path, data, resource)
}
