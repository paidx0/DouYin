package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"

	"DouYin/server/relation/pb"
	"DouYin/server/relation/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户好友列表
func (l *FriendListLogic) FriendList(in *__.ListReq) (*__.ListResp, error) {
	// 验证 Token，拿到用户ID和Key
	userClaim, err := utils.CheckToken(in.Token)
	if err != nil {
		return nil, err
	}

	// 处理业务
	// 判断用户是否存在 并拿到用户信息 user
	user := new(models.User)
	has, err := global.DBEngine.Where("uid = ? ", in.UserID).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	// 关注人数
	cnt := user.FollowCount
	if cnt < 1 {
		return nil, errors.New("好友列表为空")
	}
	friendList := make([]*__.User, 0, cnt)

	// 返回指定用户 的好友的 id，name，follow_count,follower_count,以及 请求的本人是否关注了这些用户
	err = global.DBEngine.Table("user").
		Where("user.user_key = ? AND user.deleted_at is null AND user.user_key = eachother.to_user_key", user.UserKey).
		Join("LEFT", "userfocuson", "user.user_key = userfocuson.user_key ").
		Where("userfocuson.deleted_at is null ").
		Join("LEFT", "userfocuson AS eachother", "userfocuson.to_user_key = eachother.user_key").
		Where("eachother.deleted_at is null").
		Join("LEFT", "user AS touser", "eachother.user_key = touser.user_key").
		Where("touser.deleted_at is null").
		Join("LEFT", "userfocuson  AS isfollow", "touser.user_key = isfollow.to_user_key  AND isfollow.user_key = ?", userClaim.Userkey).
		Where("isfollow.deleted_at is null").
		Select("touser.uid, touser.username, touser.follow_count, touser.follower_count, isfollow.is_follow").
		Find(&friendList)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if len(friendList) == 0 {
		return nil, errors.New("好友列表为空")
	}

	return &__.ListResp{
		StatusCode: global.Success,
		StatusMsg:  "获取好友列表成功",
		UserList:   friendList,
	}, nil
}
