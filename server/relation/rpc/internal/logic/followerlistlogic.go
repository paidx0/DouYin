package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"

	"DouYin/server/relation/pb"
	"DouYin/server/relation/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户粉丝列表
func (l *FollowerListLogic) FollowerList(in *__.ListReq) (*__.ListResp, error) {
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.ListResp{StatusCode: 1}, err
	}

	// 处理业务
	user := new(models.User)
	has, err := global.DBEngine.Where("uid = ?", in.UserId).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.ListResp{StatusCode: 1}, err
	}
	if !has {
		return &__.ListResp{StatusCode: 1}, errors.New("列表该用户不存在")
	}

	// 粉丝用户数
	cnt := user.FollowerCount
	if cnt < 1 {
		return &__.ListResp{StatusCode: 1}, errors.New("用户粉丝列表空")
	}

	userList := make([]*__.User, 0, cnt)

	// 返回该用户粉丝列表：包括用户信息，及我是否关注这些用户
	err = global.DBEngine.Table("userfocuson").Where("userfocuson.to_user_key = ? and userfocuson.deleted_at IS NULL", user.UserKey).
		Join("LEFT", "user", "user.user_key = userfocuson.user_key").Where("user.deleted_at IS NULL").
		Join("LEFT", fmt.Sprintf(`(SELECT is_follow,to_user_key from userfocuson where user_key = "%s" AND deleted_at IS NULL) myuserfocuson`, uc.Userkey), "myuserfocuson.to_user_key = userfocuson.user_key").
		Select("user.uid ,user.username ,user.follow_count ,user.follower_count, myuserfocuson.is_follow").
		Find(&userList)

	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.ListResp{StatusCode: 1}, err
	}

	return &__.ListResp{
		StatusCode: 0,
		Cnt:        int64(cnt),
		UserList:   userList,
	}, nil
}
