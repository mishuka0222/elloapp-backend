# Feeds

## Base request 
Base structure all request
### Request
```javascript
{
	"service": int32
	"method":  int32
	"data":    []byte
}
```

``````````````
## Delete user account
Returns displayable message and status 

### Request
- Service: Accounts = 100100
- Method: AccountsDelete = 100100
- Data: 
```javascript
null
```

### Response
```javascript
{
    "status": bool;
  	"message": string;
}
```
``````````````

``````````````
## Change password in profile
Returns displayable message and status 

### Request
- Service: Accounts = 100100
- Method: AccountsChangePassword = 100200
- Data: 
```javascript
{
	"prev_pass": string;
  	"new_pass": string;
}
```

### Response
```javascript
{
    "status": bool;
  	"message": string;
}
```
``````````````

``````````````
## Change Email in profile
Returns displayable message and status 

### Request
- Service: Accounts = 100100
- Method: AccountsChangeEmail = 100300
- Data: 
```javascript
{
	"new_email": string;
}
```

### Response
```javascript
{
    "status": bool;
  	"message": string;
}
```
``````````````

``````````````
## Change user information in profile and save
Returns displayable message and status 

### Request
- Service: Accounts = 100100
- Method: AccountsChangeProfile = 100400
- Data: 
```javascript
{
	"first_name": string;
	"last_name": string;
	"user_name": string;
	"bio": string;
	"gender": string;
	"birthday": string;
	"country_code": string;
}
```

### Response
```javascript
{
    "status": bool;
  	"message": string;
}
```
``````````````

``````````````
## Change password during login
Returns displayable message and status 

### Request
- Service: Accounts = 100100
- Method: AccountsForgotPassword = 100500
- Data: 
```javascript
{
	"new_pass": string;
}
```

### Response
```javascript
{
    "status": bool;
  	"message": string;
}
```
``````````````



