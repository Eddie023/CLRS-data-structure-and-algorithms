package linkedlist

import (
	"fmt"
	"testing"
)

func TestAddNodeAtEnd(t *testing.T) {

	ll := New()

	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	fmt.Println(ll.ListItems())

	ll.InsertAtLast(50)

	t.Logf("Should add newNode at the end of linkedlist")

	fmt.Println(ll.ListItems())

	if ll.GetLastItem() != 50 {
		t.Errorf("expected %d got %d", 20, ll.GetLastItem())
	}
}
