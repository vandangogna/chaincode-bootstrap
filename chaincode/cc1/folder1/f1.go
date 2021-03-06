package folder1

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// F1Chaincode definition
type F1Chaincode struct {
}

// F1Method1 returns a successful message from the current method
func (t *F1Chaincode) F1Method1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	message := fmt.Sprintf("F1Method1 - chaincode method called successfully")
	fmt.Println(message)
	return shim.Success([]byte(message))
}
