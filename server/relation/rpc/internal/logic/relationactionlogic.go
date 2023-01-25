package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"xorm.io/xorm"

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
	// 验证 Token，拿到用户ID和Key
	userClaim, err := utils.CheckToken(in.Token)
	if err != nil {
		return nil, err
	}

	// 处理业务
	// 判断对方用户是否存在 并拿到对方用户信息 toUser
	toUser := new(models.User)
	has, err := global.DBEngine.Where("uid = ? ", in.ToUserId).Get(toUser)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("对方用户不存在")
	}
	// 不能对自己操作
	if toUser.Uid == userClaim.Id {
		return nil, errors.New("不能关注或取消关注自己")
	}

	// 关注操作
	if in.ActionType == 1 {
		// 是否已经关注了该用户
		has, err := global.DBEngine.
			Where("user_key = ? and to_user_key = ? and is_follow = 1 and deleted_at IS NULL", userClaim.Userkey, toUser.UserKey).
			Exist(&models.UserFocusOn{})
		if err != nil {
			global.ZAP.Error("数据库查询失败", zap.Error(err))
			return nil, err
		}
		if has {
			return &__.ActionResp{
				StatusCode: global.Error,
				StatusMsg:  "已关注对方，关注失败",
			}, nil
		}

		// 开启事务
		// 1.自己 user 表 follow_count + 1
		// 2.对方 user 表 follower_count + 1
		// 3.用户关注表添加记录
		session := global.DBEngine.NewSession()
		defer func(session *xorm.Session) {
			errs := session.Close()
			if errs != nil {
				global.ZAP.Error("关闭事务失败", zap.Error(errs))
				return
			}
		}(session)

		if errs := session.Begin(); errs != nil {
			global.ZAP.Error("开启事务失败", zap.Error(errs))
			return nil, errs
		}
		// 第一步
		if _, err := session.Exec("update `user` set follow_count = follow_count + 1 WHERE uid = ?", userClaim.Id); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		// 第二步
		if _, err := session.Exec("update `user` set follower_count = follower_count + 1 WHERE uid = ?", in.ToUserId); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		// 第三步
		ufo := &models.UserFocusOn{
			UserKey:   userClaim.Userkey,
			ToUserKey: toUser.UserKey,
			IsFollow:  true,
		}
		if _, err := session.Insert(ufo); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		if errs := session.Commit(); errs != nil {
			global.ZAP.Error("事务提交失败", zap.Error(errs))
			return nil, errs
		}
		return &__.ActionResp{
			StatusCode: global.Success,
			StatusMsg:  "关注成功",
		}, nil
	}
	// 取消关注
	if in.ActionType == 2 {
		// 是否没关注该用户
		has, err := global.DBEngine.
			Where("user_key = ? and to_user_key = ? and is_follow = 1 and deleted_at IS NULL", userClaim.Userkey, toUser.UserKey).
			Exist(&models.UserFocusOn{})
		if err != nil {
			global.ZAP.Error("数据库查询失败", zap.Error(err))
			return nil, err
		}
		if !has {
			return nil, errors.New("尚未关注对方，取消关注失败")
		}

		// 开启事务
		// 1.自己 user 表 follow_count - 1
		// 2.对方 user 表 follower_count - 1
		// 3.用户关注表 is_follow = 0
		// 4.用户关注表删除记录
		session := global.DBEngine.NewSession()
		defer func(session *xorm.Session) {
			errs := session.Close()
			if errs != nil {
				global.ZAP.Error("关闭事务失败", zap.Error(errs))
				return
			}
		}(session)
		if err := session.Begin(); err != nil {
			global.ZAP.Error("开启事务失败", zap.Error(err))
			return nil, err
		}
		// 第一步
		if _, err := session.Exec("update `user` set follow_count = follow_count - 1 WHERE uid = ?", userClaim.Id); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		// 第二步
		if _, err := session.Exec("update `user` set follower_count = follower_count - 1 WHERE uid = ?", in.ToUserId); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		// 第三步
		if _, err := session.
			Where("user_key = ? and to_user_key = ?", userClaim.Userkey, toUser.UserKey).
			Table(new(models.UserFocusOn)).
			Update(map[string]interface{}{"is_follow": false}); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		if _, err := session.
			Where("user_key = ? and to_user_key = ?", userClaim.Userkey, toUser.UserKey).
			Delete(new(models.UserFocusOn)); err != nil {
			global.ZAP.Error("数据库操作失败", zap.Error(err))
			if errs := session.Rollback(); errs != nil {
				global.ZAP.Error("事务回滚失败", zap.Error(errs))
				return nil, errs
			}
			return nil, err
		}
		if errs := session.Commit(); errs != nil {
			global.ZAP.Error("事务提交失败", zap.Error(errs))
			return nil, errs
		}
		return &__.ActionResp{
			StatusCode: global.Success,
			StatusMsg:  "取消关注成功",
		}, nil
	}
	return &__.ActionResp{
		StatusCode: global.Error,
		StatusMsg:  "参数错误",
	}, nil
}
