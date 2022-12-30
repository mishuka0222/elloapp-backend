package mtproto

func MakeChatPhotoByPhoto(photo *Photo) (chatPhoto *ChatPhoto) {
	// chatPhoto#1c6e1c11 flags:#
	//	has_video:flags.0?true
	//	photo_id:long
	//	stripped_thumb:flags.1?bytes
	//	dc_id:int = ChatPhoto;
	switch photo.GetPredicateName() {
	case Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) == 0 {
			// TODO(@benqi): log
			chatPhoto = MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
		} else {
			// TODO(@benqi): check PhotoSize is photoSizeEmpty
			chatPhoto = MakeTLChatPhoto(&ChatPhoto{
				HasVideo:      len(photo.VideoSizes) > 0,
				PhotoId:       photo.GetId(),
				StrippedThumb: nil,
				DcId:          1,
			}).To_ChatPhoto()
		}
	default:
		chatPhoto = MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
	}
	return

}

func MakeUserProfilePhotoByPhoto(photo *Photo) (userProfilePhoto *UserProfilePhoto) {
	switch photo.GetPredicateName() {
	case Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) > 0 {
			userProfilePhoto = MakeTLUserProfilePhoto(&UserProfilePhoto{
				HasVideo:      len(photo.VideoSizes) > 0,
				PhotoId:       photo.GetId(),
				StrippedThumb: nil,
				DcId:          1,
			}).To_UserProfilePhoto()
		}
	case Predicate_photoEmpty:
		userProfilePhoto = MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}
	return
}
