package utils

import (
	"DouYin/global"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
	"strconv"
	"time"
)

// todo 视频上传到七牛云

// FileUpload 普通上传
func FileUpload(file multipart.File, fileHeader *multipart.FileHeader, filesize int64) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.CONFIG.QiNiu.Bucket,
	}
	mac := qbox.NewMac(global.CONFIG.QiNiu.AccessKey, global.CONFIG.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 指定路径上传+时间戳，保证视频唯一性
	key := global.CONFIG.QiNiu.Path + strconv.FormatInt(time.Now().Unix(), 10) + "-" + fileHeader.Filename
	err = formUploader.Put(context.Background(), &ret, upToken, key, file, filesize, &putExtra)
	if err != nil {
		global.ZAP.Error("视频投稿上传失败", zap.Error(err))
		return "", err
	}

	url = global.CONFIG.QiNiu.QiniuServer + ret.Key
	return url, nil
}

// ResumeUpload 断点续传
func ResumeUpload(file multipart.File, fileHeader *multipart.FileHeader, filesize int64) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.CONFIG.QiNiu.Bucket,
	}
	mac := qbox.NewMac(global.CONFIG.QiNiu.AccessKey, global.CONFIG.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}

	recorder, err := storage.NewFileRecorder("./tempdir")
	if err != nil {
		global.ZAP.Error("分块上传临时目录错误", zap.Error(err))
		return
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}

	// 指定路径上传+时间戳，保证视频唯一性
	key := global.CONFIG.QiNiu.Path + strconv.FormatInt(time.Now().Unix(), 10) + "-" + fileHeader.Filename
	err = resumeUploader.Put(context.Background(), &ret, upToken, key, file, filesize, &putExtra)
	if err != nil {
		global.ZAP.Error("视频投稿上传失败", zap.Error(err))
		return
	}

	url = global.CONFIG.QiNiu.QiniuServer + ret.Key
	return url, nil
}

// TempUpload 临时文件上传
func TempUpload(filePath string, fileName string) (url string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.CONFIG.QiNiu.Bucket,
	}
	mac := qbox.NewMac(global.CONFIG.QiNiu.AccessKey, global.CONFIG.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}

	recorder, err := storage.NewFileRecorder("./tempdir")
	if err != nil {
		global.ZAP.Error("分块上传临时目录错误", zap.Error(err))
		return
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}

	// 指定路径上传+时间戳，保证视频唯一性
	key := global.CONFIG.QiNiu.Path + strconv.FormatInt(time.Now().Unix(), 10) + "-" + fileName
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		global.ZAP.Error("视频投稿上传失败", zap.Error(err))
		return
	}

	url = global.CONFIG.QiNiu.QiniuServer + ret.Key
	return url, nil
}
