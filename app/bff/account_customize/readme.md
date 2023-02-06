# AccountCustomize

### Request
- Service: AccountCustomize = 100300
- Method: AccountChangeEmail = 100100
- Data:
```javascript
{
    new_email: String
}
```

### Response
```javascript
{
    Status: bool
    Message: String
}
```
------

### Request
- Service: AccountCustomize = 100300
- Method: AccountChangePassword = 100200
- Data:
```javascript
{
    prev_pass: String
    new_pass: String
}
```

### Response
```javascript
{
    Status: bool
    Message: String
}
```
------

### Request
- Service: AccountCustomize = 100300
- Method: AccountChangeProfile = 100300
- Data:
```javascript
{
    first_name: String
    last_name: String
    user_name: String
    bio: String
    gender: String
    birthday: String
    country_code: String
}
```

### Response
```javascript
{
    Status: bool
    Message: String
}
```
------

### Request
- Service: AccountCustomize = 100300
- Method: AccountDelete = 100400
- Data:
```javascript
{}
```

### Response
```javascript
{
    Status: bool
    Message: String
}
```
------

### Request
- Service: AccountCustomize = 100300
- Method: AccountConfirmationSend = 100600
- Data:
```javascript
{
    email: String
}
```

### Response
```javascript
{
    status: Bool
    message: String
    email: String
    confirmation_expire: int64
}
```
------

### Request
- Service: AccountCustomize = 100300
- Method: AccountConfirmationConfirm = 100700
- Data:
```javascript
{
    email: String
    code: String
}
```

### Response
```javascript
{
    Status: bool
    Message: String
}
```
------