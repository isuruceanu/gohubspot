package gohubspot

import "testing"
import "fmt"
import "gohubspot/contactproperty"

func TestContactPropertiesService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	props, err := testclient.ContactProperties.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(*props) == 0 {
		t.Errorf("Retuns zero properties")
	}
}

func TestContactPropertiesService_Create(t *testing.T) {
	setup()
	defer teardown()

	property := ContactProperty{Name: "lastlogindatetime",
		Label: "Last Login DateTime", DataType: contactproperty.DateTime,
		FormField: false, GroupName: "contactinformation"}

	prop, err := testclient.ContactProperties.Create(property)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(prop)
	t.Error("Just test")
}
