package gohubspot

import (
	"testing"
	"time"
)

func TestFormService_SubmitForm(t *testing.T) {
	setup()
	defer teardown()

	testportal := 62515
	formId := "1cdf3a0e-fbc1-4bae-8c88-e2b77b28d3b1"
	var now UnixTime
	now.Time = time.Now().UTC()
	f := testclient.Forms.
		AddOption("firstname", "John").
		AddOption("lastname", "Sur").
		AddOption("email", "surucion@gmail.com").
		AddOption("free_trial_start_date", (&now).String()).
		//	SetHubspotCookie("4f4aa9cc70c4f821af05ccf164c638ba").
		//	SetRemoteIpAddress("239.59.167.65:443").
		SetPageUrl("https://google.com").
		SetPageName("Registration")

	err := f.SubmitForm(testportal, formId)

	if err != nil {
		t.Error(err)
	}

	f = testclient.Forms.
		AddOption("firstname", "John").
		AddOption("lastname", "Sur").
		AddOption("email", "surucion@gmail.com").
		AddOption("free_trial_start_date", (&now).String()).
		//	SetHubspotCookie("4f4aa9cc70c4f821af05ccf164c638ba").
		//	SetRemoteIpAddress("239.59.167.65:443").
		SetPageUrl("https://google.com").
		SetPageName("Registration")

	err = f.SubmitForm(testportal, formId)

	if err != nil {
		t.Error(err)
	}

	t.Error("just error")
}
