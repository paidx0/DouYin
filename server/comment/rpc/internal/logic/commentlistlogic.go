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

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论列表
func (l *CommentListLogic) CommentList(in *__.CommentListReq) (*__.CommentListResp, error) {
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.CommentListResp{StatusCode: 1}, err
	}

	// 处理业务
	video := new(models.Video)
	has, err := global.DBEngine.Where("vid = ?", in.VideoId).Get(video)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.CommentListResp{StatusCode: 1}, err
	}
	if !has {
		return &__.CommentListResp{StatusCode: 1}, errors.New("视频不存在")
	}

	// 视频评论数
	cnt := video.CommentCount
	if cnt < 1 {
		return &__.CommentListResp{StatusCode: 1}, errors.New("视频评论空")
	}

	commentList := make([]*__.Comment, 0, cnt)

	// 返回视频评论列表，包括评论ID、评论内容和时间、及评论者的信息、及我是否关注这个评论者
	global.DBEngine.Table("comment").Where("comment.video_key = ? and comment.deleted_at IS NULL", video.VideoKey).
		Join("LEFT", "user", "user.user_key = comment.user_key").Where("user.deleted_at IS NULL").
		Join("LEFT", "userfocuson", "userfocuson.to_user_key = comment.user_key").Where("(userfocuson.user_key = ? and userfocuson.deleted_at IS NULL) or 1", uc.Userkey).
		Select("comment.cid,comment.comment_text,comment.created_at,user.uid ,user.username ,user.follow_count ,user.follower_count,userfocuson.is_follow").
		Find(&commentList)

	return &__.CommentListResp{
		Cnt:         int64(cnt),
		StatusCode:  0,
		CommentList: commentList,
	}, nil
}
