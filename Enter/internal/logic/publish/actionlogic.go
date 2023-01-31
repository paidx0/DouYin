package publish

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLogic) Action(req *types.PublishActionReq, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.PublishActionResp, err error) {
	// 验证 Token
	uc, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}

	// 用户是否存在
	user := new(models.User)
	has, err := global.DBEngine.Where("user_key = ? and username = ?", uc.Userkey, uc.Username).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	filesize := fileHeader.Size
	if req.Title == "" || filesize == 0 {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, nil
	}

	// 视频临时保存到本地
	tempVideo := "./tempdir/" + fileHeader.Filename

	out, err := os.Create(tempVideo)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}

	// 截取封面图第5帧
	err = utils.GetSnapshot(tempVideo, 5)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}

	// 上传封面文件
	pngUrl, err := utils.TempUpload(tempVideo+".png", fileHeader.Filename+".png")
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, nil
	}
	// 上传视频文件
	videoUrl, err := utils.TempUpload(tempVideo, fileHeader.Filename)
	if err != nil {
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, nil
	}

	// 上传成功后，视频信息添加到video和userpulishvideo表
	uuid := utils.UUID()
	video := &models.Video{
		Title:    req.Title,
		PlayUrl:  videoUrl,
		CoverUrl: pngUrl,
		Name:     fileHeader.Filename,
		Tag:      path.Ext(fileHeader.Filename),
		VideoKey: uuid,
	}
	userpulishvideo := &models.UserPulishVideo{
		UserKey:  uc.Userkey,
		VideoKey: uuid,
	}

	// 开启事务
	session := global.DBEngine.NewSession()
	defer session.Close()
	session.Begin()

	_, err = session.Insert(video)
	_, err = session.Insert(userpulishvideo)
	if err != nil {
		global.ZAP.Error("数据库插入失败", zap.Error(err))
		session.Rollback()
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}

	// 最后提交事务
	err = session.Commit()
	if err != nil {
		global.ZAP.Error("事务提交失败", zap.Error(err))
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}

	// 最后删除掉临时文件
	os.Remove(tempVideo)
	os.Remove(tempVideo + ".png")

	resp = &types.PublishActionResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
	}
	return
}
