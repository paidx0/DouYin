package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/comment/pb"
	"DouYin/server/comment/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论操作
func (l *CommentActionLogic) CommentAction(in *__.CommentActionReq) (*__.CommentActionResp, error) {
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.CommentActionResp{StatusCode: 1}, err
	}

	// 处理业务
	video := new(models.Video)
	has, err := global.DBEngine.Where("vid = ?", in.VideoId).Get(video)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.CommentActionResp{StatusCode: 1}, err
	}
	if !has {
		return &__.CommentActionResp{StatusCode: 1}, errors.New("视频不存在")
	}

	// 评论还是删除评论，1-发布评论，2-删除评论
	if in.ActionType == 1 && in.CommentText != "" {
		usercomment := &models.Comment{
			UserKey:     uc.Userkey,
			VideoKey:    video.VideoKey,
			CommentText: in.CommentText,
		}
		_, err = global.DBEngine.Insert(usercomment)
		if err != nil {
			global.ZAP.Error("数据库插入失败", zap.Error(err))
			return &__.CommentActionResp{StatusCode: 1}, err
		}
	} else if in.ActionType == 2 && in.CommentId != 0 {
		_, err := global.DBEngine.Delete(&models.Comment{
			UserKey:  uc.Userkey,
			VideoKey: video.VideoKey,
			Cid:      int(in.CommentId),
		})
		if err != nil {
			global.ZAP.Error("数据库删除失败", zap.Error(err))
			return &__.CommentActionResp{StatusCode: 1}, err
		}
		return &__.CommentActionResp{StatusCode: 3}, nil
	} else {
		return &__.CommentActionResp{StatusCode: 1}, nil
	}

	u := new(UserComment)
	_, err = global.DBEngine.Table("comment").
		Where("comment.video_key = ? and comment.user_key = ? and comment.comment_text = ? and comment.deleted_at IS NULL", video.VideoKey, uc.Userkey, in.CommentText).
		Join("LEFT", "user", "comment.user_key = user.user_key").Where("user.deleted_at IS NULL").
		Select("comment.cid,user.uid ,user.username ,user.follow_count ,user.follower_count ").Get(u)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.CommentActionResp{StatusCode: 1}, err
	}
	return &__.CommentActionResp{
		StatusCode: 0,
		Comment: &__.Comment{
			Cid:           u.Cid,
			Uid:           u.Uid,
			Username:      u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
		},
	}, nil
}

type UserComment struct {
	Cid           int64
	Uid           int64
	Username      string
	FollowCount   int64
	FollowerCount int64
}
