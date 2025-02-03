package address

import "testing"

// Unit test
func TestAddressType(t *testing.T)  {
	AddressForTest := "Street One"
	addressExpected := "Street"

	AddressReceived := AddressType(AddressForTest)
	if AddressReceived != addressExpected {
		t.Errorf("Unexpected type of address!\n Expected: %s \n Received: %s",
		addressExpected,
		AddressReceived,
	)
	}
}