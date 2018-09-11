package folder2

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// F2Chaincode definition
type F2Chaincode struct {
}

// F2Method1 returns a hello message from the current method
func (t *F2Chaincode) F2Method1(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	message := fmt.Sprintf("F2Method1 - chaincode method called sucessfully")
	fmt.Println(message)
	return shim.Success([]byte(message))
}
