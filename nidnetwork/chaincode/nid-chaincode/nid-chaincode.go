package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SampleNetwork struct {
}

var bcFunctions = map[string]func(shim.ChaincodeStubInterface, []string) pb.Response{
	"add_person_info":    addUser,
	"query_person_info":  queryUser,
	"remove_person_info": deleteUser,
	"update_person_info": updateUser,
}

func (s *SampleNetwork) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Initializing the network")
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Don't pass any arguments for init method")
	}
	return shim.Success(nil)
}

func (s *SampleNetwork) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fun, args := stub.GetFunctionAndParameters()

	if fun == "init" {
		return s.Init(stub)
	}

	bcFunc := bcFunctions[fun]

	if bcFunc == nil {
		return shim.Error("Invalid invoke function.")
	}

	return bcFunc(stub, args)

	return shim.Error("No such function found!!")
}

func main() {

	network := new(SampleNetwork)

	err := shim.Start(network)

	if err != nil {
		fmt.Printf("Error starting the network: %s", err)
	}
}
