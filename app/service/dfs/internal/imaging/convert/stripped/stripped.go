package main

import (
	"fmt"
	"os"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/internal/imaging"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/bytes2"
)

func main() {
	img, err := imaging.Open("../test002.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(img.Bounds())
	w := 40
	h := 40
	if img.Bounds().Dx() > img.Bounds().Dy() {
		h = 0
	} else {
		w = 0
	}
	dst := imaging.Resize(img, w, h)
	fmt.Println(dst.Bounds())

	sw := bytes2.NewBuffer(make([]byte, 0, 4096))
	err = imaging.EncodeStripped(sw, dst, 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	bb := sw.Bytes()
	fmt.Println(len(bb))

	file, err := os.Create("./test002.jpeg")
	if err != nil {
		// return err
	}
	defer file.Close()

	head := imaging.JpegHeader
	head[164] = bb[1]
	head[166] = bb[2]
	file.Write(head)
	file.Write(bb[3:])
	file.Write(imaging.JpegFooter)
}
