package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// TODO(@benqi): remove to baselib
func CalcMd5File(filename string) (string, error) {
	// fileName := core.NBFS_DATA_PATH + m.data.FilePath
	f, err := os.Open(filename)
	if err != nil {
		// log.Error(err.Error())
		return "", err
	}

	defer f.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		// fmt.Println("Copy", err)
		// log.Errorf("Copy - %v", err)
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}
