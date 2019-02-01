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

package nid

import (
	"time"
)

//Address struct
type Address struct {
	ProvinceKey     string `json:"provinceKey"` //name should be capital
	DistrictKey     string `json:"districtKey"`
	MunicipalityKey string `json:"municipalityKey"`
	WardNumber   int    `json:"wardNumber"`
}

//Name structure
type Name struct {
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

//FamilyDetails struct
type FamilyDetails struct {
	NationalIdentityNumber string  `json:"nationalIdentityNumber"`
	CitizenshipNumber      string  `json:"citizenshipNumber"`
	Name                   Name    `json:"name"`
	Nationality            string  `json:"nationality"`
	PermanentAddress       Address `json:"permanentAddress"`
}

//Province struct
type Province struct {
	Name string `json:"provinceName"`
}

//District struct
type District struct {
	Name         string `json:"districtName"`
	ProvinceKey string `json:"provinceKey"`
}

//Municipality struct
type Municipality struct {
	Name             string `json:"municipalityName"`
	TotalWards       int    `json:"totalWards"`
	MunicipalityType string `json:"municipalityType"`
	DistrictKey     string `json:"districtKey"`
}
//Sex struct
type Sex struct {
	Sex string `json:"sex"`
}

//CitizenshipType struct
type CitizenshipType struct {
	CitizenshipType string `json:"citizenshipType"`
}

//MaritalStatus struct
type MaritalStatus struct {
	MaritalStatus string `json:"maritalStatus"`
}
//MunicipalityType struct
type MunicipalityType struct {
	MunicipalityType string `json:"municipalityType"`

}
//User struct
type User struct {
	
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	MiddleName string `json:"middleName,omitempty"`
	LastName  string `json:"lastName"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	GroupName string `json:"groupName"`
	

}

//UserGroup struct 
type UserGroup struct{
	
	Name string `json:"groupName"`
	Permissions []string `json:"permissions"`
	}






// //variable declarations
// var sex = [3]string{"Male", "Female", "Others"}
// var citizenshipType = [6]string{"Janmasidha", "Janmako Adharma", "Bansaj", "Sammanartha", "Angikrit", "Baibahik Angikrit"}
// var maritalStatus = [5]string{"Married", "Single", "Widow", "Widower", "Divorced"}
// var municipalityType = [4]string{"Gaupalika", "Nagarpalika", "Upa Mahanagarpalika", "Mahanagarpalika"}

//userForm struct
type applicantForm struct {
	NationalIdentityNumber string `json:"nationalIdentityNumber"`
	//Image string `json:"image"`
	//Fingerprint string `json:"fingerprint"`
	ApplicantName Name      `json:"applicantName"`
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
	PermanentAddress Address `json:"permanentAddress"`
	//TemporaryAddress address `json:"temporaryAddress"`
	//FatherDetails familyDetails `json:"fatherDetails"`
	//MotherDetails familyDetails `json:"motherDetails"`
	//GrandFatherDetails familyDetails `json:"grandFatherDetails"`
	//SpouseDetails familyDetails `json:"spouseDetails,omitempty"`
}

//variable declaration
var permissionList = []string{
								"CAN_CREATE_PROVINCE",
								"CAN_VIEW_PROVINCE",
								"CAN_UPDATE_PROVINCE",
								"CAN_DELETE_PROVINCE",
								"CAN_CREATE_DISTRICT",
								"CAN_VIEW_DISTRICT",
								"CAN_UPDATE_DISTRICT",
								"CAN_DELETE_DISTRICT",
								"CAN_CREATE_MUNICIPALITY",
								"CAN_VIEW_MUNICIPALITY",
								"CAN_UPDATE_MUNICIPALITY",
								"CAN_DELETE_MUNICIPALITY",
								"CAN_VIEW_ALL_ADDRESS",
								"CAN_CREATE_APPLICANTFORM",
								"CAN_VIEW_APPLICANTFORM",
								"CAN_UPDATE_APPLICANTFORM",
								"CAN_DELETE_APPLICANTFORM",
								
								}