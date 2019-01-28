package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func addUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var addr = Address{Province: args[3], District: args[4], City: args[5]}

	var userInfo = UserForm{FirstName: args[1], LastName: args[2], Address: addr}

	userInfoAsBytes, _ := json.Marshal(userInfo)

	stub.PutState(args[0], userInfoAsBytes) //args[0] is the unique id that we'll supply via args

	return shim.Success(nil)
}

func queryUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	userInfoAsBytes, _ := stub.GetState(args[0])

	return shim.Success(userInfoAsBytes)
}

func deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	err := stub.DelState(args[0])

	if err == nil {
		return shim.Success(nil)
	} else {
		return shim.Error(err.Error())
	}
}

func updateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	idArray := []string{args[0]}

	deleteUser(stub, idArray)

	addUser(stub, args)

	return shim.Success(nil)
}
