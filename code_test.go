package bestmixer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestCode(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("%s/code/info", APIURL),
		httpmock.NewStringResponder(
			200,
			`{
				"error": null,
				"data": {
					"min_service_fee": {
						"btc": 1.0000,
						"bch": 3.0000,
						"ltc": 3.0000        }
				}
			}`,
		),
	)

	resp, err := client.Code.Info("bm_string")
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	assert.Equal(t, 1.0000, resp.BTC)
	assert.Equal(t, 3.0000, resp.BCH)
	assert.Equal(t, 3.0000, resp.LTC)

	t.Logf("%#v\n", resp)
}
