package nidchaincode

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//CreateProvince creates the province provided its provinceUUID and provinceName
func CreateProvince(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		UUID string `json:"provinceUUID"`
	}{}
	pr := province{}

	errp := json.Unmarshal([]byte(args[0]), &partial)
	if errp != nil {
		return shim.Error(errp.Error())
	}

	errf := json.Unmarshal([]byte(args[0]), &pr)
	if errf != nil {
		return shim.Error(errf.Error())
	}

	if partial.UUID == "" {
		return shim.Error("Expected UUID for the province")
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	provinceKey, err := stub.CreateCompositeKey("Province", []string{partial.UUID})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	provinceAsBytes, err := json.Marshal(pr)
	if err != nil {
		return shim.Error("Error marshaling province structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(provinceKey, provinceAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(provinceAsBytes)
}

//CreateDistrict creates district provided its districtUUID, districtName and provinceKey
func CreateDistrict(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		UUID string `json:"districtUUID"`
	}{}
	dt := district{}
	pr := province{}

	err := json.Unmarshal([]byte(args[0]), &partial)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &dt)
	if err != nil {
		return shim.Error(err.Error())
	}
	if partial.UUID == "" {
		return shim.Error("Expected UUID for the district")
	}

	val, err := stub.GetState(dt.ProvinceKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if val == nil {
		return shim.Error("First create the province in which you want to add district \n")
	}

	err = json.Unmarshal(val, &pr)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Marshal the data and put into the ledger
	// key for the ledger
	districtKey, err := stub.CreateCompositeKey("District", []string{pr.Name, partial.UUID})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	districtAsBytes, err := json.Marshal(dt)
	if err != nil {
		return shim.Error("Error marshaling district structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(districtKey, districtAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//CreateMunicipality creates the municipality provided municipalityUUID, municipalityName, municipalityType, totalWards and districtKey
func CreateMunicipality(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Check for correct number of arguments
	if len(args) != 1 {
		return shim.Error("Invalid argument count\n")
	}

	//Take the required structs from the argument
	partial := struct {
		UUID string `json:"municipalityUUID"`
	}{}
	mn := municipality{}
	dt := district{}

	err := json.Unmarshal([]byte(args[0]), &partial)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = json.Unmarshal([]byte(args[0]), &mn)
	if err != nil {
		return shim.Error(err.Error())
	}

	if partial.UUID == "" {
		return shim.Error("Expected UUID for the municipality")
	}

	//Check if the district exists for the municipality to be added

	val, err := stub.GetState(mn.DistrictKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if val == nil {
		return shim.Error("First create the district to add municipality in it \n")
	}

	err = json.Unmarshal(val, &dt)

	//Check if correct municipality type is provided
	municipalityTypeKey, err := stub.CreateCompositeKey("MunicipalityType",[]string{mn.MunicipalityType})

	val, err = stub.GetState(municipalityTypeKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	if val == nil {
		return shim.Error("Invalid municipality type \n")
	}


	//Marshal the data and put into the ledger
	// key for the ledger
	municipalityKey, err := stub.CreateCompositeKey("Municipality", []string{dt.Name, partial.UUID})
	if err != nil {
		return shim.Error(err.Error())
	}
	//value for the ledger
	municipalityAsBytes, err := json.Marshal(mn)
	if err != nil {
		return shim.Error("Error marshaling municipality structure")
	}
	//Put the key value pair in ledger
	err = stub.PutState(municipalityKey, municipalityAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//GetAllProvinces returns all provinces
func GetAllProvinces(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	resultsIterator, err := stub.GetStateByPartialCompositeKey("Province", []string{})
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
			ProvinceKey string `json:"provinceKey"`
			Name        string `json:"provinceName"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.ProvinceKey = kvResult.Key
		results = append(results, result)
	}
	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetAllDistrictOfProvince returns all districts of a province or all districts
func GetAllDistrictOfProvince(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	input := struct {
		ProvinceName string `json:"provinceName"`
	}{}
	if len(args) != 0 && len(args) != 1 {
		return shim.Error("Expected zero or one argument")
	}
	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	filterByProvincename := len(input.ProvinceName) > 0

	var resultsIterator shim.StateQueryIteratorInterface

	//filtering by province name if required
	if filterByProvincename {
		resultsIterator, err = stub.GetStateByPartialCompositeKey("District", []string{input.ProvinceName})
	} else {
		resultsIterator, err = stub.GetStateByPartialCompositeKey("District", []string{})
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	results := []interface{}{}
	//Iterate over the results
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response first
		result := struct {
			DistrictKey string `json:"districtKey"`
			Name        string `json:"districtName"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.DistrictKey = kvResult.Key

		results = append(results, result)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

//GetAllMunicipalityOfDistrict returns all municipality of a district or all municipalities
func GetAllMunicipalityOfDistrict(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	input := struct {
		DistrictName string `json:"districtName"`
	}{}

	if len(args) != 1 && len(args) != 0 {
		return shim.Error("Expected zero or one argument")
	}
	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	filterByDistrictName := len(input.DistrictName) > 0

	var resultsIterator shim.StateQueryIteratorInterface

	//filtering by province name if required
	if filterByDistrictName {
		resultsIterator, err = stub.GetStateByPartialCompositeKey("District", []string{input.DistrictName})
	} else {
		resultsIterator, err = stub.GetStateByPartialCompositeKey("District", []string{})
	}
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	results := []interface{}{}
	//Iterate over the results
	for resultsIterator.HasNext() {
		kvResult, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		//Construct response first
		result := struct {
			MunicipalityKey  string `json:"municipalityKey"`
			Name             string `json:"municipalityName"`
			TotalWards       int    `json:"totalWards"`
			MunicipalityType string `json:"municipalityType"`
		}{}

		err = json.Unmarshal(kvResult.Value, &result)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Fetch Key
		result.MunicipalityKey = kvResult.Key
		results = append(results, result)
	}

	resultsAsBytes, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsAsBytes)
}

type municipalityResponse struct {
	MunicipalityKey  string `json:"municipalityKey"`
	MunicipalityName string `json:"municipalityName"`
	TotalWards       string `json:"totalWards"`
	MunicipalityType string `json:"municipalityType"`
}
type districtResponse struct {
	DistrictKey    string                 `json:"districtKey"`
	DistrictName   string                 `json:"districtName"`
	Municipalities []municipalityResponse `json:"municipalities"`
}

type provinceResponse struct {
	ProvinceKey  string             `json:"provinceKey"`
	ProvinceName string             `json:"provinceName"`
	Districts    []districtResponse `json:"districts"`
}

//GetAllAddress returns all provinces, districts and municipalities
func GetAllAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Expected no arguments")
	}

	//Get all the values of province for running loop
	provinceIterator, err := stub.GetStateByPartialCompositeKey("Province", []string{})
	if err != nil {
		shim.Error(err.Error())
	}
	defer provinceIterator.Close()

	provinceResults := []provinceResponse{}
	for provinceIterator.HasNext() {
		kvResult, err := provinceIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		provinceResult := provinceResponse{}

		//Get the value of province
		err = json.Unmarshal(kvResult.Value, &provinceResult)
		if err != nil {
			return shim.Error(err.Error())
		}

		//Get the key of province
		provinceResult.ProvinceKey = kvResult.Key

		//Get all districts of the province
		districtIterator, err := stub.GetStateByPartialCompositeKey("District", []string{provinceResult.ProvinceName})
		if err != nil {
			return shim.Error(err.Error())
		}
		defer districtIterator.Close()

		districtResults := []districtResponse{}
		for districtIterator.HasNext() {
			districtKvResult, err := districtIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}

			districtResult := districtResponse{}

			//Get the key of district
			districtResult.DistrictKey = districtKvResult.Key

			//Get the value of district
			err = json.Unmarshal(districtKvResult.Value, &districtResult)
			if err != nil {
				return shim.Error(err.Error())
			}

			//Get all the municipalities of district
			municipalityIterator, err := stub.GetStateByPartialCompositeKey("Municipality", []string{districtResult.DistrictName})
			if err != nil {
				return shim.Error(err.Error())
			}
			defer municipalityIterator.Close()

			municipalityResults := []municipalityResponse{}
			for municipalityIterator.HasNext() {
				municipalityKvResult, err := municipalityIterator.Next()
				if err != nil {
					return shim.Error(err.Error())
				}

				municipalityResult := municipalityResponse{}

				//Get the key of municipality
				municipalityResult.MunicipalityKey = municipalityKvResult.Key

				//Get the value of municipality
				err = json.Unmarshal(municipalityKvResult.Value, &municipalityResult)
				if err != nil {
					return shim.Error(err.Error())
				}
				municipalityResults = append(municipalityResults, municipalityResult)
			}
			districtResult.Municipalities = municipalityResults
			districtResults = append(districtResults, districtResult)
		}
		provinceResult.Districts = districtResults
		provinceResults = append(provinceResults, provinceResult)
	}

	provinceResultsAsBytes, err := json.Marshal(provinceResults)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(provinceResultsAsBytes)
}
