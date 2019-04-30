package bestmixer

import "gopkg.in/jarcoal/httpmock.v1"

var client *Client

func setup() {
	client = New("my-api-key", nil)
	httpmock.ActivateNonDefault(client.httpClient)
}

func teardown() {
	httpmock.DeactivateAndReset()
}
