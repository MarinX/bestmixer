package bestmixer

import "fmt"

// FeeService is an interface for interfacing with the fee endpoints
// of the BestMixer API.
type FeeService interface {
	Info() (*Fee, error)
}

// FeeServiceOp handles communication with the code related methods of
// the BestMixer API.
type FeeServiceOp struct {
	client *Client
}

// FeeResource represents the response from fee/info api
type FeeResource struct {
	Fee *Fee `json:"fee_per_address"`
}

// Fee represents a BestMixer fee
type Fee struct {
	BTC string `json:"btc"`
	BCH string `json:"bch"`
	LTC string `json:"ltc"`
	ETH string `json:"eth"`
}

type feeRequest struct {
	APIKey string `json:"api_key"`
}

// Info returns max. miner's fee for each receiving address.
// To prevent blockchain analysis based on minimal fee amounts,
// a floating miner's fee for transaction processing is being used instead.
func (c *FeeServiceOp) Info() (*Fee, error) {
	path := fmt.Sprintf("%s/fee/info", APIURL)
	resource := new(FeeResource)
	err := c.client.Post(path, &feeRequest{
		APIKey: c.client.apiKey,
	}, resource)
	return resource.Fee, err
}
