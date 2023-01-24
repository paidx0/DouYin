package publish

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"mime/multipart"
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

	fileUrl, err := utils.ResumeUpload(file, fileHeader, filesize)
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
		PlayUrl:  fileUrl,
		CoverUrl: "",
		Name:     fileHeader.Filename,
		Tag:      path.Ext(fileHeader.Filename),
		VideoKey: uuid,
	}
	userpulishvideo := &models.UserPulishVideo{
		UserKey:  uc.Userkey,
		VideoKey: uuid,
	}

	_, err = global.DBEngine.Insert(video)
	_, err = global.DBEngine.Insert(userpulishvideo)
	if err != nil {
		global.ZAP.Error("数据库插入失败", zap.Error(err))
		return &types.PublishActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}, err
	}

	resp = &types.PublishActionResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
	}
	return
}
