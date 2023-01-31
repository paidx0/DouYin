package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/favorite/pb"
	"DouYin/server/favorite/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 赞操作
func (l *FavoriteActionLogic) FavoriteAction(in *__.FavoriteActionReq) (*__.FavoriteActionResp, error) {
	// 检查 UserToken
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.FavoriteActionResp{StatusCode: 1}, err
	}

	// 处理业务
	video := new(models.Video)
	has, err := global.DBEngine.Where("vid = ?", in.VideoId).Get(video)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.FavoriteActionResp{StatusCode: 1}, err
	}
	if !has {
		return &__.FavoriteActionResp{StatusCode: 1}, errors.New("视频不存在")
	}

	// 点赞还是取消点赞，1-点赞，2-取消点赞
	if in.ActionType == 1 {
		userlikevideo := &models.UserLikeVideo{
			UserKey:    uc.Userkey,
			VideoKey:   video.VideoKey,
			IsFavorite: true,
		}
		// 如果已经点赞过，
		has, err := global.DBEngine.Get(userlikevideo)
		if err != nil {
			global.ZAP.Error("数据库查询失败", zap.Error(err))
			return &__.FavoriteActionResp{StatusCode: 1}, err
		}
		// 直接返回成功
		if has {
			return &__.FavoriteActionResp{StatusCode: 0}, nil
		} else {
			// 否则添加进去
			// 开启事务
			session := global.DBEngine.NewSession()
			defer session.Close()
			session.Begin()

			_, err = session.Insert(userlikevideo)
			if err != nil {
				global.ZAP.Error("数据库插入失败", zap.Error(err))
				session.Rollback()
				return &__.FavoriteActionResp{StatusCode: 1}, err
			}
			// 点赞数+1
			_, err = session.Table("video").Where("vid = ?", video.Vid).Update(map[string]interface{}{"favorite_count": video.FavoriteCount + 1})
			if err != nil {
				global.ZAP.Error("数据库更新失败", zap.Error(err))
				session.Rollback()
				return &__.FavoriteActionResp{StatusCode: 1}, err
			}

			// 最后提交事务
			err = session.Commit()
			if err != nil {
				global.ZAP.Error("事务提交失败", zap.Error(err))
				return &__.FavoriteActionResp{StatusCode: 1}, err
			}
		}
	} else if in.ActionType == 2 {
		// 点赞数大于0才做-1操作，避免点赞数减到负数
		if video.FavoriteCount <= 0 {
			return &__.FavoriteActionResp{StatusCode: 0}, nil
		}

		session := global.DBEngine.NewSession()
		defer session.Close()
		session.Begin()

		_, err := session.Delete(&models.UserLikeVideo{
			UserKey:  uc.Userkey,
			VideoKey: video.VideoKey,
		})
		if err != nil {
			global.ZAP.Error("数据库删除失败", zap.Error(err))
			session.Rollback()
			return &__.FavoriteActionResp{StatusCode: 1}, err
		}

		// 点赞数-1
		_, err = global.DBEngine.Table("video").Where("vid = ?", video.Vid).Update(map[string]interface{}{"favorite_count": video.FavoriteCount - 1})
		if err != nil {
			global.ZAP.Error("数据库更新失败", zap.Error(err))
			session.Rollback()
			return &__.FavoriteActionResp{StatusCode: 1}, err
		}

		// 最后提交事务
		err = session.Commit()
		if err != nil {
			global.ZAP.Error("事务提交失败", zap.Error(err))
			return &__.FavoriteActionResp{StatusCode: 1}, err
		}
	} else {
		return &__.FavoriteActionResp{StatusCode: 1}, nil
	}

	return &__.FavoriteActionResp{StatusCode: 0}, nil
}
