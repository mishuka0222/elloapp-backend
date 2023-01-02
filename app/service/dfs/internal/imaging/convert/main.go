package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/minio/minio-go"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/imaging"
)

var imageFile = flag.String("image", "", "convert image file")
var isABC = flag.Bool("abc", true, "output isABC")

var (
	minioCore *minio.Core
)

func main() {
	flag.Parse()

	if *imageFile == "" {
		flag.Usage()
		return
	}

	ext := filepath.Ext(*imageFile)
	rb, err := ioutil.ReadFile(*imageFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	imaging.ReSizeImage(rb, ext, *isABC, func(szType string, localId int, w, h int32, b []byte) error {
		return ioutil.WriteFile(fmt.Sprintf("%s.%s%s", *imageFile, szType, ext), b, 0644)
	})
}

func init() {
	var err error
	minioCore, err = minio.NewCore(
		"127.0.0.1:9000",
		"TLXH0OZVP0AKOJAZ8DIT",
		"9Sw+Xbhc3aWvxQ78rRgUkTQQLLZ24SyelA+B6Rwe",
		false)
	if err != nil {
		panic("new minio error")
	}
}
