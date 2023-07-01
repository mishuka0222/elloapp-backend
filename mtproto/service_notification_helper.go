package mtproto

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
)

/*
   { updateServiceNotification
     flags: 2 [INT],
     popup: [ SKIPPED BY BIT 0 IN FIELD flags ],
     inbox_date: 1585709991 [INT],
     type: "auth1245652777904255014_1585709991" [STRING],
     message: "New login. Dear benqi, we detected a login into your account from a new device on 01/04/2020 at 02:59:51 UTC.\n\nDevice: Android\nLocation: Central, Hong Kong (IP = 52.229.158.115)\n\nIf this wasn't you, you can terminate that session in Settings > Devices (or Privacy & Security > Active Sessions)." [STRING],
     media: { messageMediaEmpty },
     entities: [ vector<0x0>
       { messageEntityBold
         offset: 0 [INT],
         length: 10 [INT],
       },
       { messageEntityUrl
         offset: 162 [INT],
         length: 14 [INT],
       },
       { messageEntityBold
         offset: 233 [INT],
         length: 18 [INT],
       },
       { messageEntityBold
         offset: 256 [INT],
         length: 36 [INT],
       },
     ],
   },
*/

func MakeSignInServiceNotification(user *User, authId int64, client, region, clientIp string) *Update {
	now := time.Now()
	notification := MakeTLUpdateServiceNotification(&Update{
		Popup:          false,
		InboxDate:      &types.Int32Value{Value: int32(now.Unix())},
		Type:           fmt.Sprintf("auth%d_%d", authId, now.Unix()),
		Message_STRING: "",
		Media:          MakeTLMessageMediaEmpty(nil).To_MessageMedia(),
		Entities:       nil,
	}).To_Update()

	notification.Message_STRING, notification.Entities = MakeTextAndMessageEntities([]MessageBuildEntry{
		{
			Text:       "",
			Param:      "New login.",
			EntityType: Predicate_messageEntityBold,
		},
		{
			Text: fmt.Sprintf("Dear %s, we detected a login into your account from a new device on %s.\n\nDevice: %s\nLocation: %s (IP = ",
				GetUserName(user),
				now.UTC(),
				client,
				region),
			Param:      clientIp,
			EntityType: Predicate_messageEntityTextUrl,
		},
		{
			Text:       ")\n\nIf this wasn't you, you can terminate that session in ",
			Param:      "Settings > Devices",
			EntityType: Predicate_messageEntityBold,
		},
		{
			Text:       " (or ",
			Param:      "Privacy & Security > Active Sessions",
			EntityType: Predicate_messageEntityBold,
		},
		{
			Text: ").",
		},
	})

	return notification
}
