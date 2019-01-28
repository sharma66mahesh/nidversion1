package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	uuid "github.com/satori/go.uuid"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status == shim.OK {
		fmt.Println("Invoke Success", string(res.Payload))
	} else {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func String(u uuid.UUID) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:8], u[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], u[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], u[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}

func getInitArguments() [][]byte {
	return [][]byte{[]byte("init")}
}

type Province struct {
	UUID string `json:"provinceUUID"`
	Name string `json:"provinceName"`
}

var nid = new(EnrollmentChaincode)
var stub = shim.NewMockStub("Enrollment Chaincode", nid)

//Check Init function
func TestInit(t *testing.T) {
	checkInit(t, stub, getInitArguments())
}

func TestProvince(t *testing.T) {
	province1 := Province{UUID: String(uuid.Must(uuid.NewV4())), Name: "Province 1"}
	province1AsBytes, err := json.Marshal(province1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("province_create"), province1AsBytes})
	checkInvoke(t,stub,[][]byte{[]byte("getProvinces")})

}
type Sex struct {
	Sex string `json:"sex"`
}
func TestSex(t *testing.T) {
	sex1 := Sex{Sex: "Male"}
	sex1AsBytes, err := json.Marshal(sex1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("sex_create"), sex1AsBytes})
	checkInvoke(t,stub,[][]byte{[]byte("sex_list")})

}

type MunicipalityType struct {
	MunicipalityType string `json:"municipalityType"`
}
func TestMunicipalityType(t *testing.T) {
	mun1 := MunicipalityType{MunicipalityType: "Mahanagarpalika"}
	mun1AsBytes, err := json.Marshal(mun1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("municipalityType_create"), mun1AsBytes})
	checkInvoke(t,stub,[][]byte{[]byte("municipalityType_list")})

}

type CitizenshipType struct {
	CitizenshipType string `json:"citizenshipType"`
}
func TestCitizenshipType(t *testing.T) {
	cit1 := CitizenshipType{CitizenshipType: "Bansaj"}
	cit1AsBytes, err := json.Marshal(cit1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("citizenshipType_create"), cit1AsBytes})
	checkInvoke(t,stub,[][]byte{[]byte("citizenshipType_list")})

}

type MaritalStatusType struct {
	MaritalStatusType string `json:"maritalStatus"`
}
func TestMaritalStatus(t *testing.T) {
	mar1 := MaritalStatusType{MaritalStatusType: "Married"}
	mar1AsBytes, err := json.Marshal(mar1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("maritalStatus_create"), mar1AsBytes})
	checkInvoke(t,stub,[][]byte{[]byte("maritalStatus_list")})

}