# AuthorizationCustomize

## Register users, confirm email and sign in methods

### Request
- Service: AuthorizationCustomize = 100100
- Method: SignUp = 100100
- Data:
```javascript
[
    {
        username:      "makhmudov",
        password:      "password",
        gender:        "Male",
        date_of_birth: "1998-06-10T00:00:00+0000",
        email:         "azeezmakhmudov@gmail.com",
        phone:         "",
        country_code:  "UZB",
        avatar:        "",
        first_name:    "",
        last_name:     "",
    }
]
```

### Response
```javascript
[{ 
    email: "azeezmakhmudov@gmail.com",
    confirmation_expire: 1673550211,
}]
```
------


- Service: AuthorizationCustomize = 100100
- Method: Confirmmation = 100300
- Data:
```javascript
[
    {
        username_or_email: "makhmudov", // or "azeezmakhmudov@gmail.com"
        code: "718293"
    }
]
```

### Response
```javascript
[{ 
    message: "success",
}]
```
------

- Service: AuthorizationCustomize = 100100
- Method: SignIn = 100200
- Data:
```javascript
[
    {
        username: "makhmudov",
        password: "password"
    }
]
```

### Response
```javascript
[
    {
        "predicate_name": "auth_authorization",
        "user": {
            "predicate_name": "user",
            "id": 777100,
            "self": true,
            "contact": true,
            "mutual_contact": true,
            "access_hash": {
                "value": 6144279944477390169
            },
            "status": {
                "predicate_name": "userStatusOffline",
                "was_online": 1673548820
            }
        }
    }
]
```
------