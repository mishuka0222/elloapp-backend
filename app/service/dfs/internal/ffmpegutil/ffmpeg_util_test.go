package ffmpegutil

import (
	"io/ioutil"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestConvertGifToMp4(t *testing.T) {
	ffmpegUtil := NewFFmpegUtil()
	data, duration, err := ffmpegUtil.ConvertToMp4ByPipe("./gsmarena_003.gif", -1, 320)
	if err != nil {
		logx.Errorf("%v", err)
		return
	}
	logx.Infof("duration = %d", duration)
	ioutil.WriteFile("./gsmarena_003.gif.mp4", data, 0644)

	oData, err := ffmpegUtil.GetFirstFrameByPipe(data)
	if err != nil {
		logx.Errorf("%v", err)
		return
	}

	ioutil.WriteFile("./gsmarena_003.gif.jpg", oData, 0644)

	md, _ := ffmpegUtil.GetMetadataByPipe(data)
	if md != nil {
		logx.Infof("%#v", md)
		w, h := GetWidthHeight(md)
		logx.Infof("duration: %d, (w, h): (%d,%d)", GetDuration(md), w, h)
	}
}
