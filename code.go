package bestmixer

import "fmt"

// CodeService is an interface for interfacing with the code endpoints
// of the BestMixer API.
type CodeService interface {
	Info(bmCode string) (*Code, error)
}

// CodeServiceOp handles communication with the code related methods of
// the BestMixer API.
type CodeServiceOp struct {
	client *Client
}

// CodeResource represents the response from code/info api
type CodeResource struct {
	Code *Code `json:"min_service_fee"`
}

// Code represents a BestMixer code
type Code struct {
	BTC float64 `json:"btc"`
	BCH float64 `json:"bch"`
	LTC float64 `json:"ltc"`
}

type codeRequest struct {
	APIKey string `json:"api_key"`
	BMCode string `json:"bm_code"`
}

// Info returns BestMixer code related information.
func (c *CodeServiceOp) Info(bmCode string) (*Code, error) {
	path := fmt.Sprintf("%s/code/info", APIURL)
	resource := new(CodeResource)
	err := c.client.Post(path, &codeRequest{
		APIKey: c.client.apiKey,
		BMCode: bmCode,
	}, resource)
	return resource.Code, err
}
