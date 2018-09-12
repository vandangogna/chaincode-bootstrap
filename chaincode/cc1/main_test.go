package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// TxID is just a dummuy transactional ID for test cases
const TxID = "mockTxID"

// MockStubUUID is an UUID value used accross all invocations to MockInit() method
const MockStubUUID = "d2490ad8-3901-11e8-b467-0ed5f89f654B"

// CheckInvoke is a common utilities method for test cases
func CheckInvoke(t *testing.T, stub *shim.MockStub, function string, args []byte) pb.Response {
	mockInvokeArgs := [][]byte{[]byte(function), args}
	res := stub.MockInvoke(MockStubUUID, mockInvokeArgs)
	if res.Status != shim.OK {
		fmt.Println("Invocation of", function, "function failed: ", string(res.Message))
		t.FailNow()
	}
	return res
}

func TestInit(t *testing.T) {
	args := [][]byte{}
	scc := new(CC1Chaincode)
	stub := shim.NewMockStub("TestInit", scc)

	res := stub.MockInit(MockStubUUID, args)
	if res.Status != shim.OK {
		fmt.Println("Initialization of chaincode failed: ", string(res.Message))
		t.FailNow()
	}
}

func TestMethod1(t *testing.T) {
	// Create mock object
	scc := new(CC1Chaincode)
	//create a mock stub
	stub := shim.NewMockStub("TestMethod1", scc)
	// Invoke actual test method
	CheckInvoke(t, stub, "Method1", []byte{})
}

func TestF1Method1(t *testing.T) {
	// Create mock object
	scc := new(CC1Chaincode)
	//create a mock stub
	stub := shim.NewMockStub("TestF1Method1", scc)
	// Invoke actual test method
	CheckInvoke(t, stub, "F1Method1", []byte{})
}

func TestF2Method1(t *testing.T) {
	// Create mock object
	scc := new(CC1Chaincode)
	//create a mock stub
	stub := shim.NewMockStub("TestF2Method1", scc)
	// Invoke actual test method
	CheckInvoke(t, stub, "F2Method1", []byte{})
}

func TestF3Method1(t *testing.T) {
	// Create mock object
	scc := new(CC1Chaincode)
	//create a mock stub
	stub := shim.NewMockStub("TestF3Method1", scc)
	// Invoke actual test method
	CheckInvoke(t, stub, "F3Method1", []byte{})
}
