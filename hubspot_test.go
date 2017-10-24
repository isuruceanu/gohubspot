package gohubspot

var (
	testclient *HubspotClient
)

// setup sets up a test client that is configured to talk to test server
func setup() {
	testclient = NewHubspotApiClient("demo")
}

func teardown() {
}
