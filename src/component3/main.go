package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(Component3Chaincode))
	if err != nil {
		fmt.Printf("Error starting Bootstrap chaincode: %s", err)
	}

}

//Download entire src.
//make sure gopath includes '/Users/.../go:/Users/....../my-forked-chaincode-bootstrap/'
//go build \ test should now work.
