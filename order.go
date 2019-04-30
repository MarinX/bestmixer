package bestmixer

import "fmt"

// OrderService is an interface for interfacing with the order endpoints
// of the BestMixer API.
type OrderService interface {
	Create(bmCode string, coin Coin, fee float64, outputs []Output) (*Order, error)
	CreateFixed(bmCode string, coin Coin, fee float64, outputs []Output) (*Order, error)
	Info(orderID string) (*OrderInfo, error)
}

// OrderServiceOp handles communication with the order related methods of
// the BestMixer API.
type OrderServiceOp struct {
	client *Client
}

// Output is object for creating output order request
type Output struct {
	Address string  `json:"address"`
	Percent float64 `json:"percent"`
	Delay   uint    `json:"delay"`
}

// Order represents the response from order/create|create_fixed api
type Order struct {
	OrderID           string  `json:"order_id"`
	InputAddress      string  `json:"input_address"`
	MinAmount         float64 `json:"min_amount,omitempty"`
	MaxAmount         float64 `json:"max_amount,omitempty"`
	ExactlyAmount     float64 `json:"exactly_amount,omitempty"`
	BmCode            string  `json:"bm_code"`
	LetterOfGuarantee string  `json:"letter_of_guarantee"`
}

// OrderInfo represents the response from order/info api
type OrderInfo struct {
	Status        string   `json:"status"`
	InputAddress  string   `json:"input_address"`
	Received      int      `json:"received"`
	Coin          Coin     `json:"coin"`
	Fixed         int      `json:"fixed"`
	MinAmount     float64  `json:"min_amount"`
	MaxAmount     float64  `json:"max_amount"`
	ServiceFee    float64  `json:"service_fee"`
	FeePerAddress float64  `json:"fee_per_address"`
	Output        []Output `json:"output"`
}

type orderRequest struct {
	APIKey string   `json:"api_key"`
	BmCode string   `json:"bm_code"`
	Coin   string   `json:"coin"`
	Fee    float64  `json:"fee"`
	Output []Output `json:"output"`
}

type orderInfoRequest struct {
	APIKey  string `json:"api_key"`
	OrderID string `json:"order_id"`
}

// Create creates order and returns order related information.
func (o *OrderServiceOp) Create(bmCode string, coin Coin, fee float64, outputs []Output) (*Order, error) {
	path := fmt.Sprintf("%s/order/create", APIURL)
	resource := new(Order)
	err := o.client.Post(path, &orderRequest{
		APIKey: o.client.apiKey,
		BmCode: bmCode,
		Coin:   coin.String(),
		Fee:    fee,
		Output: outputs,
	}, resource)
	return resource, err
}

// CreateFixed creates order and returns order related information.
func (o *OrderServiceOp) CreateFixed(bmCode string, coin Coin, fee float64, outputs []Output) (*Order, error) {
	path := fmt.Sprintf("%s/order/create_fixed", APIURL)
	resource := new(Order)
	err := o.client.Post(path, &orderRequest{
		APIKey: o.client.apiKey,
		BmCode: bmCode,
		Coin:   coin.String(),
		Fee:    fee,
		Output: outputs,
	}, resource)
	return resource, err
}

// Info returns order related information.
func (o *OrderServiceOp) Info(orderID string) (*OrderInfo, error) {
	path := fmt.Sprintf("%s/order/info", APIURL)
	resource := new(OrderInfo)
	err := o.client.Post(path, &orderInfoRequest{
		APIKey:  o.client.apiKey,
		OrderID: orderID,
	}, resource)
	return resource, err
}
