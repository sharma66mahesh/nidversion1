package nidchaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CreateApplicantForm creates the applicant form supplied by client in blockchain
/*
Sex must be "MALE", "FEMALE" or "OTHERS"
Marital Status must be "MARRIED", "SINGLE", "WIDOW", "WIDOWER", "DIVORCED"
Citizenship type must be "JANMASIDHA", "JANMAKOADHARMA", "BANSHAJ", "SAMMANARTHA", "ANGIKRIT", "BAIBAHIKANGIKRIT"
Municipality type must be "GAUPALIKA", "NAGARPALIKA", "UPAMAHANAGARPALIKA", "MAHANAGARPALIKA"
*/
func CreateApplicantForm(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check for attribute permission
	err := cid.AssertAttributeValue(stub, "CAN_CREATE_APPLICANTFORM", "true")
	if err != nil {
		return shim.Error(err.Error())
	}

	var hello int
	//Check for number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count \n")
	}

	//Take all the required data to the application struct
	applicant := applicantForm{}

	err = json.Unmarshal([]byte(args[0]), &applicant)
	if err != nil {
		return shim.Error(err.Error())
	}
	//check for Citizenship type
	resultsIterator, err := stub.GetStateByPartialCompositeKey("CitizenshipType", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			CitizenshipType string `json:"citizenShip"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.CitizenshipType == applicant.CitizenshipType {
			hello = 1

			break
		}

	}
	if hello != 1 {
		return shim.Error("Citizenship Type Invalid")
	}

	//check for Sex type
	resultsIterator, err = stub.GetStateByPartialCompositeKey("Sex", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			Sex string `json:"sex"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.Sex == applicant.Sex {
			hello = 2

			break
		}

	}
	if hello != 2 {
		return shim.Error("Sex Type Invalid")
	}

	//check for MaritalStatus type
	resultsIterator, err = stub.GetStateByPartialCompositeKey("MaritalStatus", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			MaritalStatus string `json:"maritalStatus"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.MaritalStatus == applicant.MaritalStatus {
			hello = 3

			break
		}

	}
	if hello != 3 {
		return shim.Error("Marital Status Type Invalid")
	}

	//check for Province key
	resultsIterator, err = stub.GetStateByPartialCompositeKey("Province", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			ProvinceKey string `json:"provinceKey"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.ProvinceKey == applicant.PermanentAddress.ProvinceKey {
			hello = 4

			break
		}

	}
	if hello != 4 {
		return shim.Error("Province doesnt exist.")
	}

	//check for District key
	resultsIterator, err = stub.GetStateByPartialCompositeKey("District", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			DistrictKey string `json:"districtKey"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.DistrictKey == applicant.PermanentAddress.DistrictKey {
			hello = 5

			break
		}

	}
	if hello != 5 {
		return shim.Error("District doesnt exist.")
	}

	//check for Municipality key
	resultsIterator, err = stub.GetStateByPartialCompositeKey("Municipality", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response struct
		result := struct {
			MunicipalityKey string `json:"municipalityKey"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		if result.MunicipalityKey == applicant.PermanentAddress.MunicipalityKey {
			hello = 6

			break
		}

	}
	if hello != 6 {
		return shim.Error("Municipality doesnt exist.")
	}

	//generate key value pair and insert to ledger
	applicantFormKey, err := stub.CreateCompositeKey("Applicant", []string{applicant.NationalIdentityNumber})
	if err != nil {
		return shim.Error(err.Error())
	}

	applicantFormAsBytes, err := json.Marshal(applicant)
	if err != nil {
		return shim.Error("Error marshalling applicant form structure")
	}

	//Put the key value pair in ledger
	err = stub.PutState(applicantFormKey, applicantFormAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(applicantFormAsBytes)
}
