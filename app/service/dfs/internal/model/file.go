package model

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// upload.saveFilePart#b304a621 file_id:long file_part:int bytes:bytes = Bool;
//

// CheckFileParts
// FILE_PARTS_INVALID - Invalid number of parts. The value is not between 1..3000
func CheckFileParts(fileParts int32) (err error) {
	if fileParts < 1 || fileParts > 3000 {
		err = mtproto.ErrFilePartsInvalid
	}
	return
}

// CheckFilePart
// FILE_PART_INVALID: The file part number is invalid. The value is not between 0 and 2,999.
func CheckFilePart(filePart int32) (err error) {
	if filePart < 0 || filePart > 2900 {
		err = mtproto.ErrFilePartInvalid
	}
	return
}

// CheckFilePartSize
// FILE_PART_SIZE_INVALID - 512KB cannot be evenly divided by part_size
func CheckFilePartSize(partSize int32) (err error) {
	// part_size % 1024 = 0 (divisible by 1KB)
	// 524288 % part_size = 0 (512KB must be evenly divisible by part_size)
	if partSize%1024 != 0 {
		err = mtproto.ErrFilePartLengthInvalid
	} else if 524288%partSize != 0 {
		err = mtproto.ErrFilePartSizeInvalid
	} else if partSize > 524288 {
		// FILE_PART_TOO_BIG: The size limit (512 KB) for the content of the file part has been exceeded
		err = mtproto.ErrFilePartTooBig
	}

	return
}
