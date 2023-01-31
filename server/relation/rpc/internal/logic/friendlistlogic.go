package logic

import (
	"DouYin/global"
	"DouYin/server/relation/pb"
	"DouYin/server/relation/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"fmt"
	"go.uber.org/zap"

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
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.ListResp{StatusCode: 1}, err
	}

	userList := make([]*__.User, 0)

	// 返回该用户好友列表：包括用户信息，好友列表也就是互相关注的列表
	err = global.DBEngine.Table("userfocuson").Where("userfocuson.user_key = ? and userfocuson.deleted_at IS NULL", uc.Userkey).
		Join("INNER", fmt.Sprintf(`(SELECT is_follow,user_key from userfocuson where to_user_key = "%s" AND deleted_at IS NULL) myuserfocuson`, uc.Userkey), "myuserfocuson.user_key = userfocuson.to_user_key").
		Join("INNER", "user", "user.user_key = userfocuson.user_key").Where("user.deleted_at IS NULL").
		Select("user.uid ,user.username ,user.follow_count ,user.follower_count, myuserfocuson.is_follow").
		Find(&userList)

	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.ListResp{StatusCode: 1}, err
	}

	return &__.ListResp{
		StatusCode: 0,
		Cnt:        int64(len(userList)),
		UserList:   userList,
	}, nil
}
