package bestmixer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestCreateOrder(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/create", APIURL),
		httpmock.NewStringResponder(
			200,
			`{
				"error": null,
				"data": {
					"order_id": "XXXXXXXXXX",
					"input_address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"min_amount": 0.005,
					"max_amount": 123.456,
					"bm_code": "XXXXXXXXXX",
					"letter_of_guarantee": "...Base64..."
				}
			}`,
		),
	)

	resp, err := client.Order.Create("XXXXXXXXXX", CoinBTC, 1.001, []Output{
		Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 5.5,
			Delay:   33,
		},
		Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 4.5,
			Delay:   0,
		}, Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 90,
			Delay:   121,
		},
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	assert.Equal(t, "XXXXXXXXXX", resp.OrderID)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", resp.InputAddress)
	assert.Equal(t, 0.005, resp.MinAmount)
	assert.Equal(t, 123.456, resp.MaxAmount)
	assert.Equal(t, "XXXXXXXXXX", resp.BmCode)
	assert.Equal(t, "...Base64...", resp.LetterOfGuarantee)

	t.Logf("%#v\n", resp)

}

func TestCreateFixedOrder(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/create_fixed", APIURL),
		httpmock.NewStringResponder(
			200,
			`{
				"error": null,
				"data": {
					"order_id": "XXXXXXXXXX",
					"input_address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"exactly_amount": 0.30915986,
					"bm_code": "XXXXXXXXXX",
					"letter_of_guarantee": "...Base64..."
				}
			}`,
		),
	)

	resp, err := client.Order.CreateFixed("XXXXXXXXXX", CoinBTC, 1.001, []Output{
		Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 0.1,
			Delay:   33,
		},
		Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 0.1,
			Delay:   0,
		}, Output{
			Address: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			Percent: 0.1,
			Delay:   121,
		},
	})
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	assert.Equal(t, "XXXXXXXXXX", resp.OrderID)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", resp.InputAddress)
	assert.Equal(t, 0.30915986, resp.ExactlyAmount)
	assert.Equal(t, "XXXXXXXXXX", resp.BmCode)
	assert.Equal(t, "...Base64...", resp.LetterOfGuarantee)

	t.Logf("%#v\n", resp)

}

func TestInfoOrder(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/order/info", APIURL),
		httpmock.NewStringResponder(
			200,
			`{
				"error": null,
				"data": {
					"status": "Awaiting",
					"input_address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"received": 0,
					"coin": "btc",
					"fixed": 0,
					"min_amount": 0.005,
					"max_amount": 123.456,
					"service_fee": 0.001,
					"fee_per_address": 0.0000001,
					"output": [
						{
							"address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
							"percent": 5.5,
							"delay": 33,
							"delay_countdown": 0,
							"amount": 0,
							"status": "Pending"
						},
						{
							"address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
							"percent": 4.5,
							"delay": 0,
							"delay_countdown": 0,
							"amount": 0,
							"status": "Pending"
						},
						{
							"address": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
							"percent": 90,
							"delay": 121,
							"delay_countdown": 0,
							"amount": 0,
							"status": "Pending"
						}
					]
				}
			}`,
		),
	)

	resp, err := client.Order.Info("XXXXXXXXXX")
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, "Awaiting", resp.Status)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", resp.InputAddress)
	assert.Equal(t, 0, resp.Received)
	assert.Equal(t, "btc", resp.Coin.String())
	assert.Equal(t, 0, resp.Fixed)
	assert.Equal(t, 0.005, resp.MinAmount)
	assert.Equal(t, 123.456, resp.MaxAmount)
	assert.Equal(t, 0.001, resp.ServiceFee)
	assert.Equal(t, 0.0000001, resp.FeePerAddress)
	assert.Equal(t, 3, len(resp.Output))

	t.Logf("%#v\n", resp)

}
