package dao

import userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"

const (
	versionField = "0"
)

type idxId struct {
	idx int
	id  int64
}

func removeAllNil(contacts []*userpb.ContactData) []*userpb.ContactData {
	for i := 0; i < len(contacts); {
		if contacts[i] != nil {
			i++
			continue
		}

		if i < len(contacts)-1 {
			copy(contacts[i:], contacts[i+1:])
		}

		contacts[len(contacts)-1] = nil
		contacts = contacts[:len(contacts)-1]
	}

	return contacts
}
