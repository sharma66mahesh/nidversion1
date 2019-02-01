// u must have ../

package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/nid"
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
		fmt.Println("Invoke", string(args[0]), "failed", string(res.Message))
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

type Province1 struct {
	UUID     string `json:"provinceUUID"`
	Name     string `json:"provinceName"`
	Username string `json:"username"`
}

var nid1 = new(EnrollmentChaincode)
var stub = shim.NewMockStub("Enrollment Chaincode", nid1)

//Check Init function
func TestInit(t *testing.T) {
	checkInit(t, stub, getInitArguments())
}

func TestProvinceCreate(t *testing.T) {

	TestUserCreate(t)

	province1 := Province1{UUID: String(uuid.Must(uuid.NewV4())), Name: "Province 1", Username:"kailehok"}
	province1AsBytes, err := json.Marshal(province1)
	username := struct {
		Username string `json:"username"`
	}{Username: "kailehok"}
	userBytes, err := json.Marshal(username)

	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}

	checkInvoke(t, stub, [][]byte{[]byte("province_create"), province1AsBytes})
	fmt.Println("abc")
	checkInvoke(t, stub, [][]byte{[]byte("province_list"), userBytes})

}

func TestUserGroupCreate(t *testing.T) {
	userGroup1 := nid.UserGroup{Name: "Manager", Permissions: []string{ "CAN_CREATE_DISTRICT", "CAN_CREATE_MUNICIPALITY", "CAN_VIEW_PROVINCE"}}
	groupBytes, err := json.Marshal(userGroup1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}
	userGroup2 := nid.UserGroup{Name: "Manager1", Permissions: []string{ "CAN_CREATE_DISTRICT", "CAN_CREATE_MUNICIPALITY", "CAN_VIEW_PROVINCE"}}
	groupBytes1, err := json.Marshal(userGroup2)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}
	checkInvoke(t, stub, [][]byte{[]byte("userGroup_create"), groupBytes})
	checkInvoke(t, stub, [][]byte{[]byte("userGroup_create"), groupBytes1})
	fmt.Println("abc")
	checkInvoke(t, stub, [][]byte{[]byte("userGroup_list")})

}

func TestUserCreate(t *testing.T) {

	TestUserGroupCreate(t)

	user1 := nid.User{Username: "kailehok", FirstName: "Kailash", MiddleName: "Sharan", LastName: "Baral", Password: "abc", ConfirmPassword: "abc", GroupName: "Manager"}
	user2 := nid.User{Username: "kailehok1", FirstName: "Kailash", MiddleName: "Sharan", LastName: "Baral", Password: "abc", ConfirmPassword: "abc", GroupName: "Manager"}
	userBytes1, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Error in marshalling")
		t.FailNow()
	}
	userBytes2, err := json.Marshal(user2)
	checkInvoke(t, stub, [][]byte{[]byte("user_create"), userBytes1})
	checkInvoke(t, stub, [][]byte{[]byte("user_create"), userBytes2})
	fmt.Println("abc")
	checkInvoke(t, stub, [][]byte{[]byte("user_list")})

}
