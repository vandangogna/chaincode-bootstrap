package folder3

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// F3Chaincode definition
type F3Chaincode struct {
}

// F3Method1 returns a hello message from the current method
func (t *F3Chaincode) F3Method1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	message := fmt.Sprintf("F3Method1 - chaincode method called successfully")
	fmt.Println(message)
	return shim.Success([]byte(message))
}
