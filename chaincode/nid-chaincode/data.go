package main

type Address struct {
	Province string `json:"province"` //name should be capital
	District string `json:"district"`
	City     string `json:"city"`
}

type UserForm struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   Address `json:"address"`
}
