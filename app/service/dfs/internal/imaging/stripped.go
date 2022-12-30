package imaging

import (
	"image"
	"io"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/imaging/jpeg"
)

func EncodeStripped(w io.Writer, img image.Image, quality int) error {
	var (
		rgba *image.RGBA
		err  error
	)

	if nrgba, ok := img.(*image.NRGBA); ok {
		if nrgba.Opaque() {
			rgba = &image.RGBA{
				Pix:    nrgba.Pix,
				Stride: nrgba.Stride,
				Rect:   nrgba.Rect,
			}
		}
	}

	if rgba != nil {
		err = jpeg.EncodeStripped(w, rgba, &jpeg.Options{Quality: quality})
	} else {
		err = jpeg.EncodeStripped(w, img, &jpeg.Options{Quality: quality})
	}

	return err
}
