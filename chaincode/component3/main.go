package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(BootstrapChaincode))
	if err != nil {
		fmt.Printf("Error starting Sample chaincode: %s", err)
	}

}