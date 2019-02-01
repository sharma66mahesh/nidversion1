# Enrollment Chaincode

Chaincode for NID enrollment

## Documentation for the functions of chaincode

1. Creation of province
```
function name: province_create
Arguments:
  JSON
JSON structure:
{
    "provinceUUID" : //unique identifier for province 
    "provinceName" : //name of the province
}
```
2. Creation of district
```
function name: district_create
Arguments
  JSON
{
    "districtUUID" : //unique identifier for district
    "districtName" : //Name of the district
    "provinceKey" : //Key for the province in which the district is to be added
}
```
3. Creation of Municipality
```
function name: municipality_create
Argument
JSON
{
    "municipalityUUID" : //unique identifier for municipality
    "municipalityName" :    //Name of the municipality
    "totalWards" ://total wards in the municipality
    "municipalityType" : //type of municipality should be either "Gaupalika", "Nagarpalika", "Upa Mahangarpalika", "Mahanagarpalika"
    "districtKey" : //the key for district in which it is to be added
}
```
4. Creation of applicant form
```
function name: applicantform_create
Argument:
  JSON
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
}
```
5. Get all provinces list
```
function name: province_list
Argument: No arguments required

```

6. Get all districts
```
function name: district_list
Argument:
No argument gives all district
Providing   JSON with province name gives district of province
JSON format
{
    "provinceName" : //name of the province for which districts is to be returned
}
```

7. Get municipalities of a district
```
function name: municipality_list
Argument:
No argument gives all municipality
Providing   JSON with municipality name gives municipalities of district
JSON format
{
    "districtName": //name of the district for which municipality is to be returned
}
```

8. Get all address
```
function name: address_list
Argument:
No arguments required
```


9. Create sex 
```
function name: sex_create
Argument:
JSON
{
  "sex"://name for sex  
}
```

10. Create municpality type 
```
function name: municipalityType_create
Argument:
JSON
{
  "municipalityType"://type for municipality
}
```
11. Create marital status 
```
function name: maritalStatus_create
Argument:
JSON
{
  "maritalStatus"://status of marriage
}
```
12. Create citizenship type 
```
function name: citizenshipType_create
Argument:
JSON
{
  "citizenshipType"://type of citizenship
}
```
13. Get sex values
```
function name: sex_list
Argument:
    no arguments
```
14. Get citizenship types
```
function name: citizenshipType_list
Argument:
    no arguments
```
15. Get marital status
```
function name: maritalStatus_list
Argument:
    no arguments
```
16. Get municipality type
```
function name: municipalityType_list
Argument:
    no arguments
```
