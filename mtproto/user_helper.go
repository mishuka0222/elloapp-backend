package mtproto

func ToUserIdByInput(userSelfId int64, inputUser *InputUser) int64 {
	switch inputUser.PredicateName {
	case Predicate_inputUserEmpty:
		return 0
	case Predicate_inputUserSelf:
		return userSelfId
	case Predicate_inputUser:
		return inputUser.UserId
	default:
		// log.Errorf("invalid inputUser classID - %v", inputUser)
		return 0
	}
}

func ToUserIdListByInput(userSelfId int64, inputUsers []*InputUser) []int64 {
	idList := make([]int64, 0, len(inputUsers))
	for _, user := range inputUsers {
		id := ToUserIdByInput(userSelfId, user)
		if id > 0 {
			idList = append(idList, id)
		} else {
			// ignore in
		}
	}
	return idList
}

func isUserDeleted(user *User) bool {
	return user == nil || user.PredicateName == Predicate_userEmpty || user.Deleted
}

func isUserContact(user *User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func isUserSelf(user *User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func GetUserName(user *User) string {
	if user == nil || isUserDeleted(user) {
		return "Deleted Account"
	}

	firstName := user.GetFirstName().GetValue()
	lastName := user.GetLastName().GetValue()

	if firstName == "" && lastName == "" {
		return ""
	} else if firstName == "" {
		return lastName
	} else if lastName == "" {
		return firstName
	} else {
		return firstName + " " + lastName
	}
}
