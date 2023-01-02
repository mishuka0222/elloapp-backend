package dao

import (
	"context"
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/container2"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/mr"
)

const (
	contactListKeyPrefix = "user_contact_list"
	contactKeyPrefix     = "user_contact"
)

func genContactListCacheKey(userId int64) string {
	return fmt.Sprintf("%s_%d", contactListKeyPrefix, userId)
}

func genContactCacheKey(selfId, contactId int64) string {
	return fmt.Sprintf("%s_%d_%d", contactKeyPrefix, selfId, contactId)
}

//func (d *Dao) GetUserContactIdList(ctx context.Context, id int64) (bool, []int64) {
//	var (
//		contactIdList []int64
//		keyMiss       bool
//	)
//
//	err := d.CachedConn.QueryRow(
//		ctx,
//		&contactIdList,
//		genContactListCacheKey(id),
//		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
//			idList, err := d.UserContactsDAO.SelectUserContactIdList(ctx, id)
//			if err != nil {
//				return err
//			}
//			*v.(*[]int64) = idList
//			keyMiss = true
//			return nil
//		})
//	if err != nil {
//		// return []int64{}
//	}
//
//	return keyMiss, contactIdList
//}

func (d *Dao) GetUserContactList(ctx context.Context, id int64) []*user.ContactData {
	cacheUserData := d.GetCacheUserData(ctx, id)
	if len(cacheUserData.GetContactIdList()) == 0 {
		return nil
	}

	return d.getContactListByIdList(ctx, id, cacheUserData.GetContactIdList())
}

func (d *Dao) GetUserContact(ctx context.Context, id, contactId int64) *user.ContactData {
	contacts := d.GetUserContactListByIdList(ctx, id, contactId)
	if len(contacts) == 0 {
		return nil
	}

	return contacts[0]
}

func (d *Dao) GetUserContactListByIdList(ctx context.Context, id int64, contactId ...int64) []*user.ContactData {
	cacheUserData := d.GetCacheUserData(ctx, id)
	idList := cacheUserData.GetContactIdList()
	if len(idList) == 0 {
		return nil
	}

	idList2 := make([]int64, 0, len(idList))
	for _, id2 := range contactId {
		if ok, _ := container2.Contains(id2, idList); !ok {
			idList2 = append(idList2, id2)
		}
	}
	if len(idList2) == 0 {
		return nil
	}

	return d.getContactListByIdList(ctx, id, idList2)
}

func (d *Dao) getContactListByIdList(ctx context.Context, id int64, idList []int64) []*user.ContactData {
	contactList := make([]*user.ContactData, len(idList))
	mr.ForEach(
		func(source chan<- interface{}) {
			for i, v := range idList {
				source <- idxId{i, v}
			}
		},
		func(item interface{}) {
			idx := item.(idxId)
			do := new(dataobject.UserContactsDO)
			err := d.CachedConn.QueryRow(
				ctx,
				do,
				genContactCacheKey(id, idx.id),
				func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
					do2, _ := d.UserContactsDAO.SelectContact(ctx, id, idx.id)
					if do2 == nil {
						return sqlc.ErrNotFound
					}
					*v.(*dataobject.UserContactsDO) = *do2
					return nil
				})
			if err == nil {
				contactList[idx.idx] = user.MakeTLContactData(&user.ContactData{
					UserId:        id,
					ContactUserId: do.ContactUserId,
					FirstName:     mtproto.MakeFlagsString(do.ContactFirstName),
					LastName:      mtproto.MakeFlagsString(do.ContactLastName),
					MutualContact: do.Mutual,
				}).To_ContactData()
			}
		})

	for _, v := range contactList {
		if v == nil {
			// has hole, internal error
			return nil
		}
	}

	return contactList
}

func (d *Dao) DeleteUserContact(ctx context.Context, id int64, contactId int64) {
	_, affected, _ := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(
				ctx,
				d.DB,
				func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					_, err := d.UserContactsDAO.DeleteContactsTx(tx, id, []int64{contactId})
					if err != nil {
						result.Err = err
						return
					}

					affected, err := d.UserContactsDAO.UpdateMutualTx(tx, false, contactId, id)
					result.Data = affected
					result.Err = err
				})
			return 0, tR.Data.(int64), tR.Err
		},
		genCacheUserDataCacheKey(id))

	if affected != 0 {
		d.CachedConn.DelCache(ctx, genContactCacheKey(contactId, id))
	}
}

func (d *Dao) PutUserContact(ctx context.Context, changeMutual bool, do *dataobject.UserContactsDO) error {
	keys := []string{
		genCacheUserDataCacheKey(do.OwnerUserId),
		genContactCacheKey(do.OwnerUserId, do.ContactUserId),
	}
	if changeMutual {
		keys = append(keys, genContactCacheKey(do.ContactUserId, do.OwnerUserId))
	}

	_, _, err := d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			tR := sqlx.TxWrapper(
				ctx,
				conn,
				func(tx *sqlx.Tx, result *sqlx.StoreResult) {
					if changeMutual {
						_, result.Err = d.UserContactsDAO.UpdateMutualTx(tx, true, do.ContactUserId, do.OwnerUserId)
						if result.Err != nil {
							return
						}
					}
					_, _, result.Err = d.UserContactsDAO.InsertOrUpdateTx(tx, do)
					if result.Err != nil {
						return
					}

					// // TODO(@benqi): set addPhonePrivacyException
					// if mtproto.FromBool(in.AddPhonePrivacyException) {
					// 	//
					// }
				})

			return 0, 0, tR.Err
		},
		keys...)

	return err
}
