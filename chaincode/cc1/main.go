package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(CC1Chaincode))
	if err != nil {
		fmt.Printf("Error starting Bootstrap chaincode: %s", err)
	}

}
