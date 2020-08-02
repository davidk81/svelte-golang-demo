# api definitions

list of api provided by backend service & sample data format

- TODO: use openapi/swagger generation tools
```
- auth service
	- /api/v1/session
		- POST(login)
		request:
		{
			"username": "marco",
			"password": ""
		}
		response:
		{
			"name": "Marco Polo",
			"username": "marco",
			"roles": ["nurse","admin"]
		}
		/DELETE (logout)
	- /api/v1/regiser
		- POST (register new user)
		request:
		{
			"name": "Marco Polo",
			"username": "marco",
			"password": "",
			"roles": ["nurse","admin"]
		}
		response:
		{
			"username": "marco",
			"roles": ["nurse","admin"]
		}

- user service
	- /api/v1/users (not implemented)
		- GET
		response:		
		[{
			"name": "Marco Polo",
			"username": "marco",
			"roles": ["nurse","admin"]
		},{
			"name": "Nancy Kim",
			"username": "nancy",
			"roles": ["nurse"]
		}]
	- /api/v1/user?user-id=
		- POST/GET/DELETE (not implemented)

- patient service		
	- /api/v1/patients
		- GET (list)
		response:		
		[{
			"name": "Patrick",
			"patientid": "patient1",
			"created" : "2020-08-02T11:15:08.259739Z"
		},{
			"name": "Nancy Kim",
			"patientid": "patient2",
			"created" : "2020-08-02T11:15:08.259739Z"
		}]
	- /api/v1/patient?patient-id=
		- GET
		response:		
		{
			"name": "Patrick",
			"patientid": "patient1",
			"created" : "2020-08-02T11:15:08.259739Z"
		}
	- /api/v1/patient-notes?patient-id=
		- GET (list)
		response:		
		[{
			"noteid":"6c7c14ee-d4b1-11ea-9733-0242ac130004",
			"patientid":"patient1",
			"userid":"nurse1",
			"note":"slight fever 38C",
			"created":"2020-08-02T11:15:08.259739Z"
		},
		{
			"noteid":"6e9c4cd3-d4b1-11ea-9733-0242ac130004",
			"patientid":"patient1",
			"userid":"nurse1",
			"note":"looks good ot me",
			"created":"2020-08-02T11:15:11.826367Z"
		}]
	- /api/v1/patient-note?note-id=
		- POST (new health note)
		request:
		{
			"patientid":"patient1",
			"userid":"nurse1",
			"note":"efaef"
		}		
		response:
		{
			"noteid":"6e9c4cd3-d4b1-11ea-9733-0242ac130004",
			"patientid":"patient1",
			"userid":"nurse1",
			"note":"efaef",
			"created":"2020-08-02T11:15:11.826367Z"
		}
```