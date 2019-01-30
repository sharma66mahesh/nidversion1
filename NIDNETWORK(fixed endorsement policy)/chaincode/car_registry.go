package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type CarRegistry struct {
}

type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

func (s *CarRegistry) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *CarRegistry) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	if function == "queryCar" {

		return s.queryCar(stub, args)

	} else if function == "createCar" {

		return s.createCar(stub, args)

	} else if function == "changeCarOwner" {

		return s.changeCarOwner(stub, args)
	}

	return shim.Error("Invalid function name.")
}

func (s *CarRegistry) queryCar(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	carAsBytes, _ := stub.GetState(args[0])

	return shim.Success(carAsBytes)
}

func (s *CarRegistry) createCar(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}

	carAsBytes, _ := json.Marshal(car)

	stub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *CarRegistry) changeCarOwner(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	carAsBytes, _ := stub.GetState(args[0])

	car := Car{}
	json.Unmarshal(carAsBytes, &car)

	car.Owner = args[1]
	carAsBytes, _ = json.Marshal(car)

	stub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(CarRegistry))
	if err != nil {
		fmt.Printf("Error creating new chaincode: %s", err)
	}
}
