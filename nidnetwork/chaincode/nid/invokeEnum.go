package nid

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CreateSex defined
func CreateSex(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		Sex string `json:"sex"`
	}{}

	errp := json.Unmarshal([]byte(args[0]), &partial)
	if errp != nil {
		return shim.Error(errp.Error())
	}

	if partial.Sex == "" {
		return shim.Error("Didnt get any value for sex")
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	sexKey, err := stub.CreateCompositeKey("Sex", []string{partial.Sex})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	sexAsBytes, err := json.Marshal(partial)
	if err != nil {
		return shim.Error("Error marshaling province structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(sexKey, sexAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(sexAsBytes)

}

//CreateCitizenshipType defined
func CreateCitizenshipType(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		CitizenshipType string `json:"citizenshipType"`
	}{}

	errp := json.Unmarshal([]byte(args[0]), &partial)
	if errp != nil {
		return shim.Error(errp.Error())
	}

	if partial.CitizenshipType == "" {
		return shim.Error("Didnt get any value for citizenshipType")
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	citizenKey, err := stub.CreateCompositeKey("CitizenshipType", []string{partial.CitizenshipType})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	citizenAsBytes, err := json.Marshal(partial)
	if err != nil {
		return shim.Error("Error marshaling province structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(citizenKey, citizenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(citizenAsBytes)

}

//CreateMaritalStatus defined
func CreateMaritalStatus(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		MaritalStatus string `json:"maritalstatus"`
	}{}

	errp := json.Unmarshal([]byte(args[0]), &partial)
	if errp != nil {
		return shim.Error(errp.Error())
	}

	if partial.MaritalStatus == "" {
		return shim.Error("Didnt get any value for maritalStatus")
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	maritalKey, err := stub.CreateCompositeKey("MaritalStatus", []string{partial.MaritalStatus})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	maritalAsBytes, err := json.Marshal(partial)
	if err != nil {
		return shim.Error("Error marshaling province structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(maritalKey, maritalAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(maritalAsBytes)

}

//CreateMunicipalityType defined
func CreateMunicipalityType(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		MunicipalityType string `json:"municipalityType"`
	}{}

	errp := json.Unmarshal([]byte(args[0]), &partial)
	if errp != nil {
		return shim.Error(errp.Error())
	}

	if partial.MunicipalityType == "" {
		return shim.Error("Didnt get any value for municiplityType")
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	municipalityKey, err := stub.CreateCompositeKey("MunicipalityType", []string{partial.MunicipalityType})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	municipalityAsBytes, err := json.Marshal(partial)
	if err != nil {
		return shim.Error("Error marshaling province structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(municipalityKey, municipalityAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(municipalityAsBytes)

}

//GetSex defined
func GetSex(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("Sex",[]string{})
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
			SexKey string `json:"sexKey"`
			Sex    string `json:"sex"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.SexKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetMaritalStatus defined
func GetMaritalStatus(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("MaritalStatus",[]string{})
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
			MaritalKey string `json:"maritalKey"`
			MaritalStatus string `json:"maritalStatus"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.MaritalKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetMunicipalityType defined
func GetMunicipalityType(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("MunicipalityType",[]string{})
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
			MunicipalityKey string `json:"municipalityKey"`
			MunicipalityType    string `json:"municipalityType"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.MunicipalityType = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetCitizenshipType defined
func GetCitizenshipType(stub shim.ChaincodeStubInterface,  args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("CitizenshipType",[]string{})
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
			CitizenKey string `json:"citizenKey"`
			CitizenshipType    string `json:"citizenshipType"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.CitizenKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}