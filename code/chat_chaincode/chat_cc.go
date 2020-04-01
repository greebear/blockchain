package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
// ChatChaincode example simple Chaincode implementation
type ChatChaincode struct {
}

// ===================================================================================
// Main
// ===================================================================================
func main(){
	err := shim.Start(new(ChatChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *ChatChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *ChatChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {
	case "initMember":
		//create a new Member
		return t.initMember(stub, args)
	case "queryMember":
		//query a Member
		return t.queryMember(stub, args)
	case "saveMessage":
		//save a message
		return t.saveMessage(stub, args)
	case "saveMessageUsePubKey":
		// save Encrypted message
		return t.saveMessageUsePubKey(stub, args)
	case "queryMessagesByReceiver":
		//query messages by receiver
		return t.queryMessagesByReceiver(stub, args)
	case "queryMessagesByReceiverUsePriKey":
		//query messages by receiver
		return t.queryMessagesByReceiverUsePriKey(stub, args)

	//case "readMarblePrivateDetails":
	//	//read a marble private details
	//	return t.readMarblePrivateDetails(stub, args)
	//case "transferMarble":
	//	//change owner of a specific marble
	//	return t.transferMarble(stub, args)
	//case "delete":
	//	//delete a marble
	//	return t.delete(stub, args)
	//case "queryMarblesByOwner":
	//	//find marbles for owner X using rich query
	//	return t.queryMarblesByOwner(stub, args)
	//case "queryMarbles":
	//	//find marbles based on an ad hoc rich query
	//	return t.queryMarbles(stub, args)
	//case "getMarblesByRange":
	//	//get marbles based on range query
	//	return t.getMarblesByRange(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}
}

