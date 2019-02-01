/*
 * Copyright 2018 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/nid"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("main")

// EnrollmentChaincode implementation
type EnrollmentChaincode struct {
	testMode bool
}

//Mapping the invoke functions
var bcFunctions = map[string]func(shim.ChaincodeStubInterface, []string) pb.Response{
	"province_create":         nid.CreateProvince,
	"district_create":         nid.CreateDistrict,
	"municipality_create":     nid.CreateMunicipality,
	"applicantform_create":    nid.CreateApplicantForm,
	"address_list":            nid.GetAllAddress,
	"province_list":           nid.GetAllProvinces,
	"district_list":           nid.GetAllDistrictOfProvince,
	"municipality_list":       nid.GetAllMunicipalityOfDistrict,
	"sex_create":              nid.CreateSex,
	"maritalStatus_create":    nid.CreateMaritalStatus,
	"citizenshipType_create":  nid.CreateCitizenshipType,
	"municipalityType_create": nid.CreateMunicipalityType,
	"sex_list":                nid.GetSex,
	"maritalStatus_list":      nid.GetMaritalStatus,
	"citizenshipType_list":    nid.GetCitizenshipType,
	"municipalityType_list":   nid.GetMunicipalityType,
	"userGroup_create":        nid.CreateUserGroup,
	"userGroup_list":          nid.GetUserGroups,
	"user_create":             nid.CreateUser,
	"user_list":               nid.GetUserList,
}

//Init implementation for initialising the chaincode
func (t *EnrollmentChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	var err error
	fmt.Println("Initializing Enrollment")
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		err = fmt.Errorf("No arguments expected but found %d", len(args))
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// Invoke Function accept blockchain code invocations.
func (t *EnrollmentChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "init" {
		return t.Init(stub)
	}
	bcFunc := bcFunctions[function]
	if bcFunc == nil {
		return shim.Error("Invalid invoke function.")
	}
	return bcFunc(stub, args)
}

func main() {
	logger.SetLevel(shim.LogInfo)

	twc := new(EnrollmentChaincode)
	twc.testMode = false
	err := shim.Start(twc)
	if err != nil {
		fmt.Printf("Error starting Enrollment chaincode: %s", err)
	}
}
