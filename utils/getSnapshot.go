package utils

import (
	"bytes"
	"fmt"
	"os"

	"DouYin/global"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

// todo 截取视频的某一帧作为视频封面
// 已第5帧作为封面，避免前面是黑色的画面

func GetSnapshot(videoPath string, frameNum int) error {
	buf := bytes.NewBuffer(nil)

	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		global.ZAP.Error("截取缩略图失败", zap.Error(err))
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		global.ZAP.Error("缩略图解析失败", zap.Error(err))
		return err
	}

	err = imaging.Save(img, videoPath+".png")
	if err != nil {
		global.ZAP.Error("保存本地失败", zap.Error(err))
		return err
	}

	return nil
}
