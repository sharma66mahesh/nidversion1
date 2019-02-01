#### NID enrollment version 1

To run the project 
1. Go to web folder
2. Run export IMAGE_TAG = latest in terminal
3. Run startFabric.sh script
4. Install the node modules by running `npm install`
5. Register admin by running `node enrollAdmin.js`
6. Register user by running `node registerUser.js`
7. Start the webserver by running `node app.js`

REST-API

Server is started and api can be called at localhost:4000/api/<function>



#To create user group
1.  http://localhost:4000/api/userGroup_create
2.  POST request
3.  Select Body and pass 
{
	"args":{

        
        "groupName":"Manager",
        "permissions":["CAN_CREATE_PROVINCE","CAN_VIEW_PROVINCE"]
        
    }
}
as JSON(application/json).

#To create user 
1.  http://localhost:4000/api/userGroup_create
2.  POST request
3.  Select Body and pass 
{
	"args":{

        
        "username":"kailehok",
        "firstName":"Kailash",
        "middleName":"Sharan",
        "lastName":"Baral",
        "password":"abc",
        "confirmPassword":"abc",
        "groupName":"Manager"
        
    }
}
as JSON(application/json).

#To list All Province using Postman
1.  http://localhost:4000/api/province_list
2.  GET request
3.  Select Body and pass 
{
	"args":{
        "username":"Manager"
    }
}
as JSON(application/json).


#To create Province

1.  http://localhost:4000/api/province_create
2.  POST request
3.  Select Body and pass 
```
{
    
    "args": {
        "provinceUUID": "001",
        "provinceName": "Province 1",
        "username":"Manager"
    }
}
```
as JSON(application/json).


#To create District

1.  http://localhost:4000/api/District_create
2.  POST request
3.  Select Body and pass 
```
{ 
	"args":{	
		"districtUUID" : "001",
        "districtName" : "Illam",
        "provinceKey" : "\u0000Province\u0000001\u0000" //Province key is generated after invoking create_province
        "username":"Manager"
	}
}
```

#To create Municipality

1.  http://localhost:4000/api/municipality_create
2.  POST request
3.  Select Body and pass 
```
{ 
		"args":{	
		    "municipalityUUID" : "001",
            "municipalityName" :    "Mechi",
            "totalWards" :"13",
            "municipalityType" : "Nagarpalika",
            "districtKey" : "\u0000District\u0000Province 1\u0000001\u0000",
            "username":"Manager"

	}
}
```


#To list All Addresses
1.  http://localhost:4000/api/address_list
2.  GET request
3.  Select Body and pass 
```
{
	"args":{
        "username":"Manager"
    }
}
```
as JSON(application/json).

    Response would look like this
    ```
    {
    "success": true,
    "message": [
        {
            "provinceKey": "\u0000Province\u0000001\u0000",
            "provinceName": "Province 1",
            "districts": [
                {
                    "districtKey": "\u0000District\u0000Province 1\u0000001\u0000",
                    "districtName": "Illam",
                    "municipalities": [
                        {
                            "municipalityKey": "\u0000Municipality\u0000Illam\u0000001\u0000",
                            "municipalityName": "Mechi",
                            "totalWards": 13,
                            "municipalityType": "Nagarpalika"
                        }
                    ]
                }
            ]
        }
        ]
    }
    ```


#To create Sex type 
1.  http://localhost:4000/api/sex_create
2.  POST request
3.  Select Body and pass 
```
{
	"args":{
        "sex":"Male"
        
    }
}
```
as JSON(application/json).

#To create marital status  
1.  http://localhost:4000/api/maritalStatus_create
2.  POST request
3.  Select Body and pass 
```
{
	"args":{
		"maritalStatus":"Married"
      
	}
}
```
as JSON(application/json).

#To Get sex values
1.  http://localhost:4000/api/sex_list
2.  GET request
3.  Select Body and pass 
```
{
    "args":[]
}
```
#To Get citizenship types

1.  http://localhost:4000/api/citizenshipType_list
2.  GET request
3.  Select Body and pass 

```
{
    "args":[]
    
}
```
#To Get marital status
1.  http://localhost:4000/api/maritalStatus_list
2.  GET request
3.  Select Body and pass 
```
{
    "args":[]
        
    
}
```
#To Get municipality type
1.  http://localhost:4000/api/maritalStatus_list
2.  GET request
3.  Select Body and pass 
```
{
    "args":[]
     
    
}
```
#To Create applicant form
1.  http://localhost:4000/api/applicantform_create
2.  POST request
3.  Select Body and pass 
```
{
    "args":
    {
    "nationalIdentityNumber" : //uniue identity number for the applicant:can be of uuid type
	"applicantName" : {
        "firstName" : //first name of the applicant
        "middleName" : //middle name of the applicant, can be empty
        "lastName": //last name of the applicant
    }
	"dateOfBirthBS" : //date of birth
	"sex" : //Sex can be of Male, Female or Others type
    "maritalStatus" : //can be one of "Married", "Single", "Widow", "Widower", "Divorced"
	"permanentAddress": {
        "province" : //province of applicant
        "district" : //district of applicant
        "municipality" : //municipality of applicant
        "wardNumber" : //ward number of applicant
    }
     "username":""
}
```


