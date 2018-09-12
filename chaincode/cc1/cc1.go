package main

import (
	"fmt"
	"my-forked-chaincode-bootstrap/chaincode/cc1/folder1"
	"my-forked-chaincode-bootstrap/chaincode/cc1/folder2"
	"my-forked-chaincode-bootstrap/chaincode/cc1/folder3"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CC1Chaincode chaincode
type CC1Chaincode struct {
	folder1.F1Chaincode
	folder2.F2Chaincode
	folder3.F3Chaincode
}

// Init chaincode
func (t *CC1Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### CC1Chaincode Init ###########")
	return shim.Success(nil)
}

// Invoke chaincode functions
func (t *CC1Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### CC1Chaincode Invoke ###########")
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "Method1":
		fmt.Println("########## Calling Method1 ##############")
		return t.Method1(stub, args)
	case "F1Method1":
		fmt.Println("########## Calling F1Method1 ##############")
		return t.F1Method1(stub, args)
	case "F2Method1":
		fmt.Println("########## Calling F2Method1 ##############")
		return t.F2Method1(stub, args)
	case "F3Method1":
		fmt.Println("########## Calling F3Method1 ##############")
		return t.F3Method1(stub, args)
	}

	errorMsg := fmt.Sprintf("Unknown action '%s'. Please check the first argument, the function name.", function)
	return shim.Error(errorMsg)
}

//Method1 returns a successful message from the current method
func (t *CC1Chaincode) Method1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Method1() called successfully")
	return shim.Success([]byte("Method1() called successfully"))
}
