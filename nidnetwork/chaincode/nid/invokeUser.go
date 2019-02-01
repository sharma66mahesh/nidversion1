package nid

import (
	"encoding/json"
	// "fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


func authUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	input := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	authenticated := false

	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	userKey, err := stub.CreateCompositeKey("User", []string{input.Username})
	if err != nil {
		return shim.Error(err.Error())
	}
	userBytes, _ := stub.GetState(userKey)
	if len(userBytes) == 0 {
		authenticated = false
	} else {
		user := User{}
		err := json.Unmarshal(userBytes, &user)
		if err != nil {
			return shim.Error(err.Error())
		}
		authenticated = user.Password == input.Password
	}

	authBytes, _ := json.Marshal(authenticated)
	return shim.Success(authBytes)
}

func getUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	input := struct {
		Username string `json:"username"`
	}{}

	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	userKey, err := stub.CreateCompositeKey("User", []string{input.Username})
	if err != nil {
		return shim.Error(err.Error())
	}
	userBytes, _ := stub.GetState(userKey)
	if len(userBytes) == 0 {
		return shim.Success(nil)
	}

	response := struct {
		Username  string `json:"username"`
		FirstName string `json:"firstName"`
		MiddleName string `json:"middleName,omitempty"`
		LastName  string `json:"lastName"`
	}{}
	err = json.Unmarshal(userBytes, &response)
	if err != nil {
		return shim.Error(err.Error())
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(responseBytes)
}

//CreateUser defined
func CreateUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	user:=User{}
		
	err := json.Unmarshal([]byte(args[0]), &user)
	if err != nil {
		return shim.Error(err.Error())
	}
	if user.Password != user.ConfirmPassword{
		return shim.Error(err.Error())
	}

	groupKey, err := stub.CreateCompositeKey("UserGroup", []string{user.GroupName})
	if err != nil {
		return shim.Error(err.Error())
	}
	groupBytes, _ := stub.GetState(groupKey)
	if len(groupBytes) == 0 {
		return shim.Error("User group doesnt exists.")
	}
	
	userKey, err := stub.CreateCompositeKey("User", []string{user.Username})
	if err != nil {
		return shim.Error(err.Error())
	}
	userBytes, _ := stub.GetState(userKey)
	if len(userBytes) != 0 {
		return shim.Error(err.Error())
	}

	//value for the ledger
	userAsBytes, err := json.Marshal(user)

	if err != nil {
		return shim.Error(err.Error())
	}
		//Put the key value pair in ledger
		err = stub.PutState(userKey, userAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	
		return shim.Success(userAsBytes)
}
//CreateUserGroup defined
func CreateUserGroup(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	group:=UserGroup{}
		
	err := json.Unmarshal([]byte(args[0]), &group)
	if err != nil {
		return shim.Error(err.Error())
	}

	
	groupKey, err := stub.CreateCompositeKey("UserGroup", []string{group.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	groupBytes, _ := stub.GetState(groupKey)
	if len(groupBytes) != 0 {
		return shim.Error(err.Error())
	}

	//value for the ledger
	groupAsBytes, err := json.Marshal(group)

	if err != nil {
		return shim.Error(err.Error())
	}
		//Put the key value pair in ledger
		err = stub.PutState(groupKey, groupAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	
		return shim.Success(groupAsBytes)
}

//GetUserGroups defined
func GetUserGroups(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("UserGroup", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			GroupKey string `json:"groupKey"`
			Name        string `json:"groupName"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.GroupKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetUserList defined
func GetUserList(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("User", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	results := []interface{}{}
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			UserKey string `json:"userKey"`
			Name        string `json:"username"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.UserKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}
