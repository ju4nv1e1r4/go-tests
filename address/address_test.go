package address

import (
	"testing"
)

type testScenarios struct {
	addressInserted string
	addressExpected string
}

// Unit test
func TestAddressType(t *testing.T)  {
	
	data_receveid := []testScenarios {
		{ "Avenue One", "Avenue" },
		{ "Street One", "Street" },
		{ "Road One", "Road" },
		{ "Longview St", "Invalid Type" },
		{ "", "Invalid Type" },
		{ "STREET ZERO", "Street" },
		{ "Ave Marvelous", "Invalid Type" },
		{ "Highway 2H", "Highway" },
	}

	for _, scenario := range data_receveid {
		address_type := AddressType(scenario.addressInserted)
		if address_type != scenario.addressExpected {
			t.Errorf("Unexpected type of address!\n Expected: %s \n Received: %s",
				address_type,
				scenario.addressExpected,
			)
		}
	}
}
