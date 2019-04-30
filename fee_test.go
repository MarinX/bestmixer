package bestmixer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestFee(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/fee/info", APIURL),
		httpmock.NewStringResponder(
			200,
			`{
				"error": null,
				"data": {
					"fee_per_address": {
						"btc": 0.000123,
						"bch": 0.000456,
						"ltc": 0.000789,
						"eth": 0.000321
					}
				}
			}`,
		),
	)

	resp, err := client.Fee.Info()
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, 0.000123, resp.BTC)
	assert.Equal(t, 0.000456, resp.BCH)
	assert.Equal(t, 0.000789, resp.LTC)
	assert.Equal(t, 0.000321, resp.ETH)

	t.Logf("%#v\n", resp)
}
