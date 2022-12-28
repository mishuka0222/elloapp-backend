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
---

## Get Feeds List
Returns displayable feed list

### Request
- Service: Feed = 100100
- Method: GetFeedList = 100100
- Data: 
```javascript
null
```

### Response
```javascript
[{ 
    "chat_id":   int64, 
    "photo_id":  int64, 
    "title":     string, 
    "status":    bool, 
    "peer_type": int32 
}]
```
------


## Update Feeds List
Update user feed list

### Request
- Service: Feed = 100100
- Method: UpdateFeedList = 100200
- Data:  
```javascript
[{ 
    "chat_id":   int64, 
    "peer_type": int32 
}]
```
### Response
```javascript
 null
```
---------


## Get History Count
Returns count unread feeds

### Request
- Service: Feed = 100100
- Method: GetHistoryCounter = 100300
- Data:
```javascript
null
```

### Response
```javascript
{ 
    "count": int32 
}
```
---

## Read History
Update user history unread count messages

### Request
- Service: Feed = 100100
- Method: ReadHistory = 100400
- Data:
```javascript
[{
    "chat_id":   int64, 
    "max_id":    int32, 
    "peer_type": int32
}]
```
### Response
```javascript
[{
    "pts":       int32, 
    "pts_count": int32,
    "chat_id":   int64, 
    "peer_type": int32
}]
```
---------



## Get History
Returns array of messages

### Request
- Service: Feed = 100100
- Method: GetHistory = 100500
- Data:
```javascript
// may be null
{
    "before": int32, // If you need read previous messages mast be above zero
    "limit":  int32 // count messages need to return
}
```
### Response
```javascript
// Examle:
[{
    "messages": [{
        "predicate_name": "message",
        "id": 105,
        "peer_id": {
            "predicate_name": "peerChat",
            "chat_id": 4
        },
        "from_id": {
            "predicate_name": "peerUser",
            "user_id": 777010
        },
        "date": 1670259463,
        "media": {
            "predicate_name": "messageMediaPhoto",
            "photo_FLAGPHOTO": {
                "predicate_name": "photo",
                "id": 1599810259570921500,
                "access_hash": 35745184112654140,
                "date": 1670259463,
                "sizes": [
                    {
                        "predicate_name": "photoSize",
                        "type": "s",
                        "w": 60,
                        "h": 90,
                        "size2": 3770
                    },
                    {
                        "predicate_name": "photoSize",
                        "type": "m",
                        "w": 214,
                        "h": 320,
                        "size2": 46527
                    },
                    {
                        "predicate_name": "photoSize",
                        "type": "x",
                        "w": 534,
                        "h": 800,
                        "size2": 284230
                    },
                    {
                        "predicate_name": "photoSize",
                        "type": "y",
                        "w": 854,
                        "h": 1280,
                        "size2": 617087
                    }
                ],
                "dc_id": 1
            }
        },
        "grouped_id": {
            "value": 1599810260606914600
        }
        },{
        "predicate_name": "message",
        "id": 107,
        "peer_id": {
            "predicate_name": "peerChat",
            "chat_id": 4
        },
        "from_id": {
            "predicate_name": "peerUser",
            "user_id": 777017
        },
        "date": 1670259526,
        "media": {
            "predicate_name": "messageMediaDocument",
            "document": {
                "predicate_name": "document",
                "id": 1599810521463263200,
                "access_hash": 35745183158134492,
                "date": 1670259526,
                "mime_type": "video/mp4",
                "size2_INT64": 4725376,
                "thumbs": [
                    {
                        "predicate_name": "photoStrippedSize",
                        "type": "i",
                        "bytes": "ASgg1kgZunFVm3C2GfStFDtzVOUZsVP+wK15tTK2hSsMD7UxOAGAyPpVW5lZ2klKsYwfvccmpInAXyQwBnl6ntioLu1libyMlkB3bj3PrWTN0bcl4OkaZ9c8VFJK/wBk2FBtwBmlRDwePxFOdP3IHoKepnoUbhY4toMZH91gOc4yaYhmlnJkcEHAUA8H61KJvtMSMSAVJHPNNeJmGUYcYOBUGt9C5LcQwD53A9qpz6mSpEMf/Am4qDVvviq/8AqrkWQlvIQu3POT/OtBPljZ+5WsyD735/zrUP8AqP8AgNKRcVuf"
                    },
                    {
                        "predicate_name": "photoSize",
                        "type": "m",
                        "w": 254,
                        "h": 320,
                        "size2": 24671
                    }
                ],
                "dc_id": 1,
                "attributes": [
                    {
                        "predicate_name": "documentAttributeVideo",
                        "constructor": 250621158,
                        "w": 672,
                        "h": 848,
                        "supports_streaming": true,
                        "duration": 32
                    },
                    {
                        "predicate_name": "documentAttributeFilename",
                        "constructor": 358154344,
                        "file_name": "IMG_5082.MOV"
                    }
                ],
                "size2_INT32": 4725376
            }
        }
        },{
        "predicate_name": "message",
        "id": 108,
        "peer_id": {
            "predicate_name": "peerChat",
            "chat_id": 4
        },
        "from_id": {
            "predicate_name": "peerUser",
            "user_id": 777010
        },
        "date": 1670259655,
        "message": "Есть висть"
        },{
        "predicate_name": "message",
        "id": 109,
        "peer_id": {
            "predicate_name": "peerChat",
            "chat_id": 4
        },
        "from_id": {
            "predicate_name": "peerUser",
            "user_id": 777010
        },
        "date": 1670259719,
        "media": {
            "predicate_name": "messageMediaDocument",
            "document": {
                "predicate_name": "document",
                "id": 1599782111483531300,
                "access_hash": 5947573369625459000,
                "date": 1670252747,
                "mime_type": "audio/mp3",
                "size2_INT64": 5662638,
                "dc_id": 1,
                "attributes": [
                    {
                        "predicate_name": "documentAttributeAudio",
                        "constructor": -1739392570,
                        "duration": 235,
                        "title": {
                            "value": " Don't worry, be happy"
                        },
                        "performer": {
                            "value": "Bobby Mcferryn "
                        }
                    },
                    {
                        "predicate_name": "documentAttributeFilename",
                        "constructor": 358154344,
                        "file_name": "08-Don't worry, be happy.mp3"
                    }
                ],
                "size2_INT32": 5662638
            }
        }
        },{
        "predicate_name": "message",
        "id": 110,
        "peer_id": {
            "predicate_name": "peerChat",
            "chat_id": 4
        },
        "from_id": {
            "predicate_name": "peerUser",
            "user_id": 777010
        },
        "reply_to": {
            "predicate_name": "messageReplyHeader",
            "reply_to_msg_id": 107
        },
        "date": 1670259861,
        "message": "Тест"
        }],
    "chats": [{
        "predicate_name": "chat",
        "id": 2,
        "creator": true,
        "title": "Hello Merehead",
        "photo": {
            "predicate_name": "chatPhoto",
            "photo_id": 1599797183391993900,
            "dc_id": 1
        },
        "participants_count_INT32": 8,
        "date": 1670256159,
        "version": 9,
        "default_banned_rights": {
            "predicate_name": "chatBannedRights",
            "until_date": 2147483647
        }
        },{
        "predicate_name": "chat",
        "id": 7,
        "title": "ElloApp x Merehead",
        "photo": {
            "predicate_name": "chatPhoto",
            "photo_id": 1599868340136841200,
            "dc_id": 1
        },
        "participants_count_INT32": 6,
        "date": 1670272826,
        "version": 6,
        "admin_rights": {
            "predicate_name": "chatAdminRights",
            "change_info": true,
            "delete_messages": true,
            "ban_users": true,
            "invite_users": true,
            "pin_messages": true,
            "manage_call": true,
            "other": true
        },
        "default_banned_rights": {
            "predicate_name": "chatBannedRights",
            "until_date": 2147483647
        }
        }],
    "users": [{
        "predicate_name": "user",
        "id": 777017,
        "access_hash": {
            "value": 8268654797255415000
        },
        "first_name": {
            "value": "Slava"
        },
        "last_name": {
            "value": "Ostrogliad"
        },
        "username": {
            "value": "vostrogliad"
        },
        "photo": {
            "predicate_name": "userProfilePhoto",
            "photo_id": 1599854158049972200,
            "dc_id": 1
        },
        "status": {
            "predicate_name": "userStatusOffline",
            "was_online": 1671129018
        }
    }],
    "chat_id_info": [{
        "chat_id": 13
    },{
        "chat_id": 2,
        "count": 63
    }]
}]
```
---------
