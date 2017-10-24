package gohubspot

import (
	"fmt"
	"testing"
	"time"
)

func TestContactService_CreateContact(t *testing.T) {
	setup()
	defer teardown()

	var now UnixTime
	now.Time = time.Now().UTC()

	contact := Properties{}

	contact.AddProperty("Email", "ion.suruceanu@gapsquare.com")
	contact.AddProperty("NAME", "Ion Suruceanu")
	contact.AddProperty("firstname", "Ion")
	contact.AddProperty("lastname", "Suruceanu")
	contact.AddProperty("lastlogindatetime", &now)

	fmt.Println(contact)

	ip, err := testclient.Contacts.Create(contact)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ip)

	t.Error("Just to test")
}
