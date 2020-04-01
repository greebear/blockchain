package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryCollection string, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetPrivateDataQueryResult(queryCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")

		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	// ==== Decode buffer.String() ==== # use when client get the string from network
	//b := bytes.NewBufferString(buffer.String())
	//var datas []map[string]interface{}
	//err = json.Unmarshal(b.Bytes(),&datas)
	//if err != nil{
	//	panic(err)
	//}
	//for index, data := range datas {
	//	fmt.Println(index, data)
	//}

	return buffer.Bytes(), nil
}

func getDecryptedQueryResultForQueryString(
	stub shim.ChaincodeStubInterface,
	queryCollection string,
	queryString string,
	privateKey string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetPrivateDataQueryResult(queryCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		// ==== Unmarshal Data ====
		type messageEncrypt struct {
			ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
			Receiver     string `json:"receiver"`    //the fieldtags are needed to keep case from bouncing around
			Sender   string `json:"sender"`
			Context    []byte `json:"context"`
			Date       string `json:"date"`
		}
		var messageReturn messageEncrypt
		err = json.Unmarshal(queryResponse.Value, &messageReturn)
		if err != nil{
			fmt.Println("error in unmarshall data")
		}

		// ==== EccDecrypt message ====
		msg, err := EccDecrypt(messageReturn.Context, []byte(privateKey))
		if err != nil {
			fmt.Println(err)
			// ==== Record is a JSON object, so we write as-is ====
			buffer.WriteString(string(queryResponse.Value))
		} else{
			messageReturn.Context = msg
			// ==== Turin into json type ====
			type jsonMessageReconstruct struct {
				ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
				Receiver     string `json:"receiver"`    //the fieldtags are needed to keep case from bouncing around
				Sender   string `json:"sender"`
				Context    string `json:"context"`
				Date       string `json:"date"`
			}
			jsonMessage := &jsonMessageReconstruct{
				ObjectType:  messageReturn.ObjectType,
				Receiver:      messageReturn.Receiver,
				Sender:    messageReturn.Sender,
				Context:     string(messageReturn.Context),
				Date: 		 messageReturn.Date,
			}
			messageJsonBytes, err := json.Marshal(jsonMessage)
			if err != nil{
				fmt.Println("error in Marshal data")
			}
			buffer.WriteString(string(messageJsonBytes))
		}
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())
	// ==== Decode buffer.String() ==== # use when client get the string from network
	//b := bytes.NewBufferString(buffer.String())
	//var datas []map[string]interface{}
	//err = json.Unmarshal(b.Bytes(),&datas)
	//if err != nil{
	//	panic(err)
	//}
	//for index, data := range datas {
	//	fmt.Println(index, data)
	//}

	return buffer.Bytes(), nil
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}