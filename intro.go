package main

import (
	"fmt"
	"intro/address"
)


func main()  {
	addresValidation := address.AddressType("Street One")
	fmt.Println(addresValidation)
}
