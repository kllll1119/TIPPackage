package main

import (
	"TIPPackage"
	"fmt"
)

func main() {
	fmt.Println(TIPPackage.GetAddress("47.56.100.100"))
	//fmt.Println(ippackage.GetAddressShort("47.56.100.100"))

	fmt.Println(TIPPackage.GetAddress("118.113.243.200"))

	fmt.Println(TIPPackage.GetAddress("192.168.50.121"))
}
