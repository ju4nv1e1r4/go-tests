package address

import (
	"strings"
)

// AddressType: Verify if addres type is valid
func AddressType(address string) string {
	validTypes := []string{"street", "avenue", "road", "highway"}

	lowerCaseAddress := strings.ToLower(address)
	firstWorld := strings.Split(lowerCaseAddress, " ")[0]

	validTypeAddress := false
	for _, typeAddress := range validTypes{
		if typeAddress == firstWorld {
			validTypeAddress = true
		}
	}

	if validTypeAddress {
		return strings.Title(firstWorld)
		
	}

	return "Invalid Type"
}
