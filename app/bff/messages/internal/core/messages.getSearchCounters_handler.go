package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetSearchCounters
// messages.getSearchCounters#732eef00 peer:InputPeer filters:Vector<MessagesFilter> = Vector<messages.SearchCounter>;
func (c *MessagesCore) MessagesGetSearchCounters(in *mtproto.TLMessagesGetSearchCounters) (*mtproto.Vector_Messages_SearchCounter, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
	)
	counters := &mtproto.Vector_Messages_SearchCounter{
		Datas: make([]*mtproto.Messages_SearchCounter, 0, len(in.GetFilters())),
	}

	for _, filter := range in.GetFilters() {
		/*
			{
			    "constructor": "CRC32_messages_getSearchCounters",
			    "peer": {
			        "predicate_name": "inputPeerChat",
			        "constructor": "CRC32_inputPeerChat",
			        "chat_id": 10053
			    },
			    "filters": [
			        {
			            "predicate_name": "inputMessagesFilterPhotoVideo",
			            "constructor": "CRC32_inputMessagesFilterPhotoVideo"
			        },
			        {
			            "predicate_name": "inputMessagesFilterDocument",
			            "constructor": "CRC32_inputMessagesFilterDocument"
			        },
			        {
			            "predicate_name": "inputMessagesFilterRoundVoice",
			            "constructor": "CRC32_inputMessagesFilterRoundVoice"
			        },
			        {
			            "predicate_name": "inputMessagesFilterUrl",
			            "constructor": "CRC32_inputMessagesFilterUrl"
			        },
			        {
			            "predicate_name": "inputMessagesFilterMusic",
			            "constructor": "CRC32_inputMessagesFilterMusic"
			        },
			        {
			            "predicate_name": "inputMessagesFilterGif",
			            "constructor": "CRC32_inputMessagesFilterGif"
			        }
			    ]
			}
		*/

		/*
		 */
		var (
			fType = mtproto.FromMessagesFilter(filter)
			mType int32
		)

		switch fType {
		case mtproto.FilterPhotoVideo:
			mType = mtproto.MEDIA_PHOTOVIDEO
		case mtproto.FilterDocument:
			mType = mtproto.MEDIA_FILE
		case mtproto.FilterUrl:
			mType = mtproto.MEDIA_URL
		case mtproto.FilterGif:
			mType = mtproto.MEDIA_GIF
		case mtproto.FilterMusic:
			mType = mtproto.MEDIA_MUSIC
		case mtproto.FilterRoundVoice:
			mType = mtproto.MEDIA_AUDIO
		default:
			/*
				[
				    {
				        "predicate_name": "inputMessagesFilterPhotos",
				        "constructor": "CRC32_inputMessagesFilterPhotos"
				    },
				    {
				        "predicate_name": "inputMessagesFilterVideo",
				        "constructor": "CRC32_inputMessagesFilterVideo"
				    }
				]
			*/
			counter := mtproto.MakeTLMessagesSearchCounter(&mtproto.Messages_SearchCounter{
				Inexact: false,
				Filter:  filter,
				Count:   0,
			}).To_Messages_SearchCounter()
			counters.Datas = append(counters.Datas, counter)
			continue
		}

		sz, _ := c.svcCtx.Dao.MessageClient.MessageGetSearchCounter(
			c.ctx,
			&message.TLMessageGetSearchCounter{
				UserId:    c.MD.UserId,
				PeerType:  peer.PeerType,
				PeerId:    peer.PeerId,
				MediaType: mType,
			})

		counter := mtproto.MakeTLMessagesSearchCounter(&mtproto.Messages_SearchCounter{
			Inexact: false,
			Filter:  filter,
			Count:   sz.GetV(),
		}).To_Messages_SearchCounter()
		counters.Datas = append(counters.Datas, counter)
		c.Logger.Infof("messages.getSearchCounters - result: %s", counter)
	}

	return counters, nil
}
