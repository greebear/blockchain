package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"strings"
	"time"

	pb "github.com/hyperledger/fabric/protos/peer"
)

// ============================================================
// saveMessage - save Message into chaincode state
// ============================================================
func (t *ChatChaincode) saveMessage(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	type messageTransientInput struct {
		Receiver    string `json:"receiver"` //the fieldtags are needed to keep case from bouncing around
		Sender  string `json:"sender"`
		Context   string `json:"context"`
	}

	// ==== Input sanitation ====
	fmt.Println("- start to save message")

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Message must be passed in transient map.")
	}

	// ==== Get transient ====
	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	// ==== Check transient  ====
	if _, ok := transMap["message"]; !ok {
		return shim.Error("message must be a key in the transient map")
	}
	if len(transMap["message"]) == 0 {
		return shim.Error("message value in the transient map must be a non-empty JSON string")
	}

	// ==== Decode transient and Check value ====
	var messageInput messageTransientInput
	err = json.Unmarshal(transMap["message"], &messageInput)
	fmt.Println("messageInput", messageInput)
	if len(messageInput.Sender) == 0 {
		return shim.Error("'from' field must be a non-empty string")
	}
	if len(messageInput.Receiver) == 0 {
		return shim.Error("'to' field must be a non-empty string")
	}
	if len(messageInput.Context) == 0 {
		return shim.Error("'context' field must be a non-empty string")
	}

	// ==== Check if message already exists ====
	// TODO: may check Receiver, Sender, Context and SysTime

	// ==== Create message object, marshal to JSON====
	type messageNoEncrypt struct {
		ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
		Receiver     string `json:"receiver"`    //the fieldtags are needed to keep case from bouncing around
		Sender   string `json:"sender"`
		Context    string `json:"context"`
		Date       string `json:"date"`
	}
	message := &messageNoEncrypt{
		ObjectType:  "message",
		Receiver:      messageInput.Receiver,
		Sender:    messageInput.Sender,
		Context:     messageInput.Context,
		Date: 		 time.Now().Format("2006-01-02 15:04:05"),
	}
	messageJSONasBytes, err := json.Marshal(message)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save message to state ====
	timeStampStr := strconv.FormatInt(time.Now().UnixNano(),10)
	err = stub.PutPrivateData("collectionMessages", timeStampStr, messageJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Message saved. Return success ====
	fmt.Println("- end save message")
	return shim.Success(nil)
}
// ============================================================
// saveMessageWithPriKey - save Encrypted Message into chaincode state
// ============================================================
func (t *ChatChaincode) saveMessageUsePubKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var jsonResp string

	// ==== Input sanitation ====
	fmt.Println("- start to save message")

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Message must be passed in transient map.")
	}

	// ==== Get transient ====
	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}

	// ==== Check transient  ====
	if _, ok := transMap["message"]; !ok {
		return shim.Error("message must be a key in the transient map")
	}
	if len(transMap["message"]) == 0 {
		return shim.Error("message value in the transient map must be a non-empty JSON string")
	}

	// ==== Decode transient and Check value ====
	type messageTransientInput struct {
		Receiver    string `json:"receiver"` //the fieldtags are needed to keep case from bouncing around
		Sender  string `json:"sender"`
		Context   string `json:"context"`
	}
	var messageInput messageTransientInput
	err = json.Unmarshal(transMap["message"], &messageInput)

	if len(messageInput.Sender) == 0 {
		return shim.Error("'from' field must be a non-empty string")
	}
	if len(messageInput.Receiver) == 0 {
		return shim.Error("'to' field must be a non-empty string")
	}
	if len(messageInput.Context) == 0 {
		return shim.Error("'context' field must be a non-empty string")
	}

	// ==== Check if message already exists ====
	// TODO: may check Receiver, Sender, Context and SysTime

	// ==== Get member from chaincode state ====
	valAsbytes, err := stub.GetPrivateData("collectionMembers", messageInput.Receiver)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + messageInput.Receiver + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Member does not exist: " + messageInput.Receiver + "\"}"
		return shim.Error(jsonResp)
	}
	// ==== Decode Jsonreturn and Check value ====
	type memberInf struct {
		ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
		Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
		PublicKey  string `json:"publickey"`
		PrivateKey string `json:"privatekey"`
	}
	var memberInfDencode memberInf
	err = json.Unmarshal(valAsbytes, &memberInfDencode)
	if err != nil {
		return shim.Error("Failed to decode member json: " + err.Error())
	}

	// ==== Use Ecc to Encrypt Context ====
	//fmt.Println("【ECC】 Encrypt PublicKey: ", memberInfDencode.PublicKey)
	//fmt.Println("【ECC】 before Encrypt text: ", messageInput.Context)
	cryptText, err := EccEncrypt([]byte(messageInput.Context), []byte(memberInfDencode.PublicKey))
	if err != nil {
		return shim.Error("Failed to EccEncrypt Context: " + err.Error())
	}
	//fmt.Println("【ECC】 after Encrypt text string: ", string(cryptText))
	//fmt.Println("【ECC】 privateKey string : ", memberInfDencode.PrivateKey)
	//fmt.Println("【ECC】 after Encrypt text byte: ", cryptText)
	//fmt.Println("【ECC】 privateKey byte : ", []byte(memberInfDencode.PrivateKey))
	//decryptedText, err := EccDecrypt(cryptText, []byte(memberInfDencode.PrivateKey))
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("【ECC】 Decrypt text: ", string(decryptedText))
	// ==== Create message object, marshal to JSON====
	type messageEncrypt struct {
		ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
		Receiver     string `json:"receiver"`    //the fieldtags are needed to keep case from bouncing around
		Sender   string `json:"sender"`
		Context    []byte `json:"context"`
		Date       string `json:"date"`
	}
	message := &messageEncrypt{
		ObjectType:  "message",
		Receiver:      messageInput.Receiver,
		Sender:    messageInput.Sender,
		Context:     cryptText,
		Date: 		 time.Now().Format("2006-01-02 15:04:05"),
	}
	messageJSONasBytes, err := json.Marshal(message)
	//fmt.Println("【JSON】json bytes: ", messageJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Save message to state ====
	timeStampStr := strconv.FormatInt(time.Now().UnixNano(),10)
	err = stub.PutPrivateData("collectionMessagesEncrypt", timeStampStr, messageJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Message saved. Return success ====
	fmt.Println("- end save message")
	return shim.Success(nil)
}

// ===== Example: Parameterized rich query =================================================
// queryMarblesByOwner queries for marbles based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (Receiver).
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *ChatChaincode) queryMessagesByReceiver(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	receiver := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"message\",\"receiver\":\"%s\"}}", receiver)

	queryResults, err := getQueryResultForQueryString(stub, "collectionMessages", queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *ChatChaincode) queryMessagesByReceiverUsePriKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	receiver := strings.ToLower(args[0])
	privatekey := args[1]


	// ==== Decode transient and Check value ====
	fmt.Println(receiver, "'s privateKey: ", privatekey)

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"message\",\"receiver\":\"%s\"}}", receiver)

	queryResults, err := getDecryptedQueryResultForQueryString(stub, "collectionMessagesEncrypt", queryString, privatekey)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}