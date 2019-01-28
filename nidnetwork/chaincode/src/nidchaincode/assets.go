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

package nidchaincode

import (
	"time"
)

//address struct
type address struct {
	ProvinceKey     string `json:"provinceKey"` //name should be capital
	DistrictKey     string `json:"districtKey"`
	MunicipalityKey string `json:"municipalityKey"`
	WardNumber      int    `json:"wardNumber"`
}

//name structure
type name struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

//familyDetails struct
type familyDetails struct {
	NationalIdentityNumber string  `json:"nationalIdentityNumber"`
	CitizenshipNumber      string  `json:"citizenshipNumber"`
	Name                   name    `json:"name"`
	Nationality            string  `json:"nationality"`
	PermanentAddress       address `json:"permanentAddress"`
}

//province struct
type province struct {
	Name string `json:"provinceName"`
}

//district struct
type district struct {
	Name        string `json:"districtName"`
	ProvinceKey string `json:"provinceKey"`
}

//municipality struct
type municipality struct {
	Name             string `json:"municipalityName"`
	TotalWards       int    `json:"totalWards"`
	MunicipalityType string `json:"municipalityType"`
	DistrictKey      string `json:"districtKey"`
}

//sex struct
type sex struct {
	Sex string `json:"sex"`
}

//citizenshipType struct
type citizenshipType struct {
	CitizenshipType string `json:"citizenshipType"`
}

//maritalStatus struct
type maritalStatus struct {
	MaritalStatus string `json:"maritalStatus"`
}

//municipalityType struct
type municipalityType struct {
	MunicipalityType string `json:"municipalityType"`
}

//userForm struct
type applicantForm struct {
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	//Image string `json:"image"`
	//Fingerprint string `json:"fingerprint"`
	ApplicantName name      `json:"applicantName"`
	DateOfBirthBS time.Time `json:"dateOfBirthBS"`
	//CitizenshipNumber string `json:"citizenshipNumber"`
	//CitizenshipIssuedDateBS string `json:"citizenshipIssuedDateBS"`
	CitizenshipType string `json:"citizenshipType"`
	//MobileNumber string `json:"mobileNumber"`
	//Caste string `json:"caste"`
	Sex string `json:"sex"`
	//Religion string `json:"religion"`
	//AcademicQualification string `json:"academicQualification"`
	//Occupation string `json:"occupation"`
	MaritalStatus    string  `json:"maritalStatus"`
	PermanentAddress address `json:"permanentAddress"`
	//TemporaryAddress address `json:"temporaryAddress"`
	//FatherDetails familyDetails `json:"fatherDetails"`
	//MotherDetails familyDetails `json:"motherDetails"`
	//GrandFatherDetails familyDetails `json:"grandFatherDetails"`
	//SpouseDetails familyDetails `json:"spouseDetails,omitempty"`
}
