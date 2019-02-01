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

#To list All Province using Postman
1.  http://localhost:4000/api/province_list
2.  GET request
3.  Select Body and pass 
{
	"args":[]
}
as JSON(application/json).


#To create Province

1.  http://localhost:4000/api/province_create
2.  POST request
3.  Select Body and pass 
{
    "success": true,
    "message": "Chaincode invoked",
    "args": {
        "provinceUUID": "001",
        "provinceName": "Province 1"
    }
}
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
	}
}


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
            "districtKey" : "\u0000District\u0000Province 1\u0000001\u0000"
	}
}
```


#To list All Addresses
1.  http://localhost:4000/api/address_list
2.  GET request
3.  Select Body and pass 
{
	"args":[]
}
as JSON(application/json).

    Response would look like this
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
                    "municipalities": []
                }
            ]
        }
        ]
    }


