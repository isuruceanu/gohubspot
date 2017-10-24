package gohubspot

import "testing"
import "fmt"

func TestContactListsService_GetContactLists(t *testing.T) {
	setup()
	defer teardown()

	lists, err := testclient.ContactLists.GetContactLists()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(lists)
}

func TestContactListService_CreateContactList(t *testing.T) {
	setup()
	defer teardown()

	list, err := testclient.ContactLists.CreateContactList("GoHubspot List1")

	if err != nil {
		t.Error(err)
	}

	fmt.Println(list)
	t.Error("ddd")
}
