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

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注和取消操作
func (l *RelationActionLogic) RelationAction(in *__.ActionReq) (*__.ActionResp, error) {
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.ActionResp{StatusCode: 1}, err
	}

	// 处理业务
	touser := new(models.User)
	has, err := global.DBEngine.Where("uid = ?", in.ToUserId).Get(touser)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.ActionResp{StatusCode: 1}, err
	}
	if !has {
		return &__.ActionResp{StatusCode: 1}, errors.New("列表该用户不存在")
	}
	// 不能关注自己
	if touser.UserKey == uc.Userkey {
		return &__.ActionResp{StatusCode: 1}, errors.New("用户不允许关注自己")
	}
	myuser := new(models.User)
	global.DBEngine.Where("uid = ?", uc.Id).Get(myuser)

	// 关注还是取消关注，1-关注，2-取消关注
	if in.ActionType == 1 {
		userfocuson := &models.UserFocusOn{
			UserKey:   uc.Userkey,
			ToUserKey: touser.UserKey,
			IsFollow:  true,
		}
		// 如果已经关注过
		has, err := global.DBEngine.Get(userfocuson)
		if err != nil {
			global.ZAP.Error("数据库查询失败", zap.Error(err))
			return &__.ActionResp{StatusCode: 1}, err
		}
		// 直接返回成功
		if has {

			return &__.ActionResp{StatusCode: 0}, nil
		} else {
			// 否则添加进去
			// 开启事务
			session := global.DBEngine.NewSession()
			defer session.Close()
			session.Begin()

			_, err := session.Insert(userfocuson)
			if err != nil {
				global.ZAP.Error("数据库插入失败", zap.Error(err))
				session.Rollback()
				return &__.ActionResp{StatusCode: 1}, err
			}
			// 更新关注数 +1
			_, err = session.Table("user").Where("uid = ?", myuser.Uid).Update(map[string]interface{}{"follow_count": myuser.FollowCount + 1})
			_, err = session.Table("user").Where("uid = ?", touser.Uid).Update(map[string]interface{}{"follower_count": touser.FollowerCount + 1})
			if err != nil {
				global.ZAP.Error("数据库更新失败", zap.Error(err))
				session.Rollback()
				return &__.ActionResp{StatusCode: 1}, err
			}
			// 最后提交事务
			err = session.Commit()
			if err != nil {
				global.ZAP.Error("事务提交失败", zap.Error(err))
				return &__.ActionResp{StatusCode: 1}, err
			}
		}
	} else if in.ActionType == 2 {
		session := global.DBEngine.NewSession()
		defer session.Close()
		session.Begin()

		_, err := session.Delete(&models.UserFocusOn{
			UserKey:   uc.Userkey,
			ToUserKey: touser.UserKey,
		})
		if err != nil {
			global.ZAP.Error("数据删除失败", zap.Error(err))
			session.Rollback()
			return &__.ActionResp{StatusCode: 1}, err
		}
		// 更新关注数 -1
		_, err = session.Table("user").Where("uid = ?", myuser.Uid).Update(map[string]interface{}{"follow_count": myuser.FollowCount - 1})
		_, err = session.Table("user").Where("uid = ?", touser.Uid).Update(map[string]interface{}{"follower_count": touser.FollowerCount - 1})
		if err != nil {
			global.ZAP.Error("数据库更新失败", zap.Error(err))
			session.Rollback()
			return &__.ActionResp{StatusCode: 1}, err
		}
		// 最后提交事务
		err = session.Commit()
		if err != nil {
			global.ZAP.Error("事务提交失败", zap.Error(err))
			return &__.ActionResp{StatusCode: 1}, err
		}
	} else {
		return &__.ActionResp{StatusCode: 1}, nil
	}

	return &__.ActionResp{StatusCode: 0}, nil
}
