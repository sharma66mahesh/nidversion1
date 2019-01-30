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

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"

	cc "nidchaincode"

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
	"province_create":         cc.CreateProvince,
	"district_create":         cc.CreateDistrict,
	"municipality_create":     cc.CreateMunicipality,
	"applicantform_create":    cc.CreateApplicantForm,
	"address_list":            cc.GetAllAddress,
	"province_list":           cc.GetAllProvinces,
	"district_list":           cc.GetAllDistrictOfProvince,
	"municipality_list":       cc.GetAllMunicipalityOfDistrict,
	"sex_create":              cc.CreateSex,
	"maritalStatus_create":    cc.CreateMaritalStatus,
	"citizenshipType_create":  cc.CreateCitizenshipType,
	"municipalityType_create": cc.CreateMunicipalityType,
	"sex_list":                cc.GetSex,
	"maritalStatus_list":      cc.GetMaritalStatus,
	"citizenshipType_list":    cc.GetCitizenshipType,
	"municipalityType_list":   cc.GetMunicipalityType,
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
	msp, err := cid.GetMSPID(stub)

	if err != nil {
		return shim.Error(err.Error())
	}
	if msp != "mohaMSP" && msp != "ecMSP" {
		return shim.Error("You don't have certificate from valid MSP")
	}

	certificate, err := cid.GetX509Certificate(stub)
	if err != nil {
		return shim.Error(err.Error())
	}

	if certificate.Issuer.CommonName != "ca.moha.nid.com" && certificate.Issuer.CommonName != "ca.ec.nid.com" {
		return shim.Error("Incorrect certificate issuer name")
	}

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
