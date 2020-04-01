package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"io/ioutil"

	pb "github.com/hyperledger/fabric/protos/peer"
)

type member struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	PublicKey  string `json:"publickey"`
	PrivateKey string `json:"privatekey"`
}

// ============================================================
// initMember - create public/private key for new member , store public key into chaincode state
// ============================================================
func (t *ChatChaincode) initMember(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	type memberTransientInput struct {
		Name  string `json:"name"` //the fieldtags are needed to keep case from bouncing around
	}

	// ==== Input sanitation ====
	fmt.Println("- start init key for new member")

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Member information must be passed in transient map.")
	}

	// ==== Get transient ====
	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	// ==== Check transient  ====
	if _, ok := transMap["member"]; !ok {
		return shim.Error("member must be a key in the transient map")
	}
	if len(transMap["member"]) == 0 {
		return shim.Error("member value in the transient map must be a non-empty JSON string")
	}

	// ==== Decode transient and Check value ====
	var memberInput memberTransientInput
	err = json.Unmarshal(transMap["member"], &memberInput)
	if err != nil {
		return shim.Error("Failed to decode member json: " + err.Error())
	}
	if len(memberInput.Name) == 0 {
		return shim.Error("name field must be a non-empty string")
	}

	// ==== Check if member already exists ====
	memberAsBytes, err := stub.GetPrivateData("collectionMembers", memberInput.Name)
	if err != nil {
		return shim.Error("Failed to get member: " + err.Error())
	} else if memberAsBytes != nil {
		fmt.Println("This member already exists: " + memberInput.Name)
		return shim.Error("This member already exists: " + memberInput.Name)
	}

	// ==== Generate public/private key ====
	GetEccKey()
	publicKey, err := ioutil.ReadFile("eccpublic.pem")
	if err != nil {
		return shim.Error("File reading error: " + err.Error())
	}
	fmt.Println(string(publicKey))

	privateKey, err := ioutil.ReadFile("eccprivate.pem")
	if err != nil {
		return shim.Error("File reading error: " + err.Error())
	}
	fmt.Println(string(privateKey))

	// ==== Create member object, marshal to JSON====
	member := &member{
		ObjectType:  "member",
		Name:        memberInput.Name,
		PublicKey:   string(publicKey),
		PrivateKey:  string(privateKey),
	}
	memberJSONasBytes, err := json.Marshal(member)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save member to state ====
	err = stub.PutPrivateData("collectionMembers", memberInput.Name, memberJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Member saved. Return success ====
	fmt.Println("- end init member")
	return shim.Success(memberJSONasBytes)
}


// ===============================================
// queryMember - query a member from chaincode state
// ===============================================
func (t *ChatChaincode) queryMember(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the member to query")
	}

	name = args[0]

	// ==== Get member from chaincode state ====
	valAsbytes, err := stub.GetPrivateData("collectionMembers", name)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Member does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
