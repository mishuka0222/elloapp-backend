/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2022-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package dialog

const (
	Predicate_updateDraftMessage                       = "updateDraftMessage"
	Predicate_dialogExt                                = "dialogExt"
	Predicate_dialogPinnedExt                          = "dialogPinnedExt"
	Predicate_dialogFilterExt                          = "dialogFilterExt"
	Predicate_dialog_saveDraftMessage                  = "dialog_saveDraftMessage"
	Predicate_dialog_clearDraftMessage                 = "dialog_clearDraftMessage"
	Predicate_dialog_getAllDrafts                      = "dialog_getAllDrafts"
	Predicate_dialog_clearAllDrafts                    = "dialog_clearAllDrafts"
	Predicate_dialog_markDialogUnread                  = "dialog_markDialogUnread"
	Predicate_dialog_toggleDialogPin                   = "dialog_toggleDialogPin"
	Predicate_dialog_getDialogUnreadMarkList           = "dialog_getDialogUnreadMarkList"
	Predicate_dialog_getDialogsByOffsetDate            = "dialog_getDialogsByOffsetDate"
	Predicate_dialog_getDialogs                        = "dialog_getDialogs"
	Predicate_dialog_getDialogsByIdList                = "dialog_getDialogsByIdList"
	Predicate_dialog_getDialogsCount                   = "dialog_getDialogsCount"
	Predicate_dialog_getPinnedDialogs                  = "dialog_getPinnedDialogs"
	Predicate_dialog_reorderPinnedDialogs              = "dialog_reorderPinnedDialogs"
	Predicate_dialog_getDialogById                     = "dialog_getDialogById"
	Predicate_dialog_getTopMessage                     = "dialog_getTopMessage"
	Predicate_dialog_updateReadInbox                   = "dialog_updateReadInbox"
	Predicate_dialog_updateReadOutbox                  = "dialog_updateReadOutbox"
	Predicate_dialog_insertOrUpdateDialog              = "dialog_insertOrUpdateDialog"
	Predicate_dialog_deleteDialog                      = "dialog_deleteDialog"
	Predicate_dialog_getUserPinnedMessage              = "dialog_getUserPinnedMessage"
	Predicate_dialog_updateUserPinnedMessage           = "dialog_updateUserPinnedMessage"
	Predicate_dialog_insertOrUpdateDialogFilter        = "dialog_insertOrUpdateDialogFilter"
	Predicate_dialog_deleteDialogFilter                = "dialog_deleteDialogFilter"
	Predicate_dialog_updateDialogFiltersOrder          = "dialog_updateDialogFiltersOrder"
	Predicate_dialog_getDialogFilters                  = "dialog_getDialogFilters"
	Predicate_dialog_getDialogFolder                   = "dialog_getDialogFolder"
	Predicate_dialog_editPeerFolders                   = "dialog_editPeerFolders"
	Predicate_dialog_getChannelMessageReadParticipants = "dialog_getChannelMessageReadParticipants"
	Predicate_dialog_setChatTheme                      = "dialog_setChatTheme"
	Predicate_dialog_setHistoryTTL                     = "dialog_setHistoryTTL"
)

var clazzNameRegisters2 = map[string]map[int]int32{
	Predicate_updateDraftMessage: {
		0: -155335502, // 0xf6bdc4b2

	},
	Predicate_dialogExt: {
		0: -1109809056, // 0xbdd9a860

	},
	Predicate_dialogPinnedExt: {
		0: 245834284, // 0xea7222c

	},
	Predicate_dialogFilterExt: {
		0: -1891683854, // 0x8f3f31f2

	},
	Predicate_dialog_saveDraftMessage: {
		0: 1321916826, // 0x4ecad99a

	},
	Predicate_dialog_clearDraftMessage: {
		0: -76500326, // 0xfb70b29a

	},
	Predicate_dialog_getAllDrafts: {
		0: -1394716698, // 0xacde4fe6

	},
	Predicate_dialog_clearAllDrafts: {
		0: 1102614780, // 0x41b890fc

	},
	Predicate_dialog_markDialogUnread: {
		0: 1160941838, // 0x4532910e

	},
	Predicate_dialog_toggleDialogPin: {
		0: -2038504145, // 0x867ee52f

	},
	Predicate_dialog_getDialogUnreadMarkList: {
		0: -893634316, // 0xcabc38f4

	},
	Predicate_dialog_getDialogsByOffsetDate: {
		0: -1652652540, // 0x9d7e8604

	},
	Predicate_dialog_getDialogs: {
		0: -2046091754, // 0x860b1e16

	},
	Predicate_dialog_getDialogsByIdList: {
		0: -1390049167, // 0xad258871

	},
	Predicate_dialog_getDialogsCount: {
		0: -533089179, // 0xe039b465

	},
	Predicate_dialog_getPinnedDialogs: {
		0: -1463673931, // 0xa8c21bb5

	},
	Predicate_dialog_reorderPinnedDialogs: {
		0: -18664089, // 0xfee33567

	},
	Predicate_dialog_getDialogById: {
		0: -1587594251, // 0xa15f3bf5

	},
	Predicate_dialog_getTopMessage: {
		0: -92425614, // 0xfa7db272

	},
	Predicate_dialog_updateReadInbox: {
		0: 489158840, // 0x1d27f8b8

	},
	Predicate_dialog_updateReadOutbox: {
		0: 1483799934, // 0x5870fd7e

	},
	Predicate_dialog_insertOrUpdateDialog: {
		0: -317723281, // 0xed0fed6f

	},
	Predicate_dialog_deleteDialog: {
		0: 28515811, // 0x1b31de3

	},
	Predicate_dialog_getUserPinnedMessage: {
		0: -1885617487, // 0x8f9bc2b1

	},
	Predicate_dialog_updateUserPinnedMessage: {
		0: 371388970, // 0x1622f22a

	},
	Predicate_dialog_insertOrUpdateDialogFilter: {
		0: 178824068, // 0xaa8a384

	},
	Predicate_dialog_deleteDialogFilter: {
		0: 31276695, // 0x1dd3e97

	},
	Predicate_dialog_updateDialogFiltersOrder: {
		0: -1321465025, // 0xb13c0b3f

	},
	Predicate_dialog_getDialogFilters: {
		0: 1818717244, // 0x6c676c3c

	},
	Predicate_dialog_getDialogFolder: {
		0: 1092325045, // 0x411b8eb5

	},
	Predicate_dialog_editPeerFolders: {
		0: 608601754, // 0x2446869a

	},
	Predicate_dialog_getChannelMessageReadParticipants: {
		0: 683494715, // 0x28bd4d3b

	},
	Predicate_dialog_setChatTheme: {
		0: -374431190, // 0xe9aea22a

	},
	Predicate_dialog_setHistoryTTL: {
		0: 165263532, // 0x9d9b8ac

	},
}

var clazzIdNameRegisters2 = map[int32]string{
	-155335502:  Predicate_updateDraftMessage,                       // 0xf6bdc4b2
	-1109809056: Predicate_dialogExt,                                // 0xbdd9a860
	245834284:   Predicate_dialogPinnedExt,                          // 0xea7222c
	-1891683854: Predicate_dialogFilterExt,                          // 0x8f3f31f2
	1321916826:  Predicate_dialog_saveDraftMessage,                  // 0x4ecad99a
	-76500326:   Predicate_dialog_clearDraftMessage,                 // 0xfb70b29a
	-1394716698: Predicate_dialog_getAllDrafts,                      // 0xacde4fe6
	1102614780:  Predicate_dialog_clearAllDrafts,                    // 0x41b890fc
	1160941838:  Predicate_dialog_markDialogUnread,                  // 0x4532910e
	-2038504145: Predicate_dialog_toggleDialogPin,                   // 0x867ee52f
	-893634316:  Predicate_dialog_getDialogUnreadMarkList,           // 0xcabc38f4
	-1652652540: Predicate_dialog_getDialogsByOffsetDate,            // 0x9d7e8604
	-2046091754: Predicate_dialog_getDialogs,                        // 0x860b1e16
	-1390049167: Predicate_dialog_getDialogsByIdList,                // 0xad258871
	-533089179:  Predicate_dialog_getDialogsCount,                   // 0xe039b465
	-1463673931: Predicate_dialog_getPinnedDialogs,                  // 0xa8c21bb5
	-18664089:   Predicate_dialog_reorderPinnedDialogs,              // 0xfee33567
	-1587594251: Predicate_dialog_getDialogById,                     // 0xa15f3bf5
	-92425614:   Predicate_dialog_getTopMessage,                     // 0xfa7db272
	489158840:   Predicate_dialog_updateReadInbox,                   // 0x1d27f8b8
	1483799934:  Predicate_dialog_updateReadOutbox,                  // 0x5870fd7e
	-317723281:  Predicate_dialog_insertOrUpdateDialog,              // 0xed0fed6f
	28515811:    Predicate_dialog_deleteDialog,                      // 0x1b31de3
	-1885617487: Predicate_dialog_getUserPinnedMessage,              // 0x8f9bc2b1
	371388970:   Predicate_dialog_updateUserPinnedMessage,           // 0x1622f22a
	178824068:   Predicate_dialog_insertOrUpdateDialogFilter,        // 0xaa8a384
	31276695:    Predicate_dialog_deleteDialogFilter,                // 0x1dd3e97
	-1321465025: Predicate_dialog_updateDialogFiltersOrder,          // 0xb13c0b3f
	1818717244:  Predicate_dialog_getDialogFilters,                  // 0x6c676c3c
	1092325045:  Predicate_dialog_getDialogFolder,                   // 0x411b8eb5
	608601754:   Predicate_dialog_editPeerFolders,                   // 0x2446869a
	683494715:   Predicate_dialog_getChannelMessageReadParticipants, // 0x28bd4d3b
	-374431190:  Predicate_dialog_setChatTheme,                      // 0xe9aea22a
	165263532:   Predicate_dialog_setHistoryTTL,                     // 0x9d9b8ac

}

func GetClazzID(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}
