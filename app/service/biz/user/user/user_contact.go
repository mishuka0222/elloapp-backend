package user

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

type BlockedList []*mtproto.PeerBlocked

// Len len
func (m BlockedList) Len() int {
	return len(m)
}
func (m BlockedList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m BlockedList) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Date < m[j].Date
}

type Contact struct {
	SelfUserId    int32  `json:"self_user_id"`
	ContactUserId int32  `json:"contact_user_id"`
	PhoneNumber   string `json:"phone_number"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	MutualContact bool   `json:"mutual_contact"`
}
