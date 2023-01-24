package publish

import (
	"DouYin/Enter/internal/logic/publish"
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.uber.org/zap"
	"net/http"
)

func ActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := publish.NewActionLogic(r.Context(), svcCtx)

		// 拿到视频数据
		file, fileHeader, err := r.FormFile("data")
		if err != nil {
			global.ZAP.Error("投稿视频数据错误", zap.Error(err))
			httpx.Error(w, err)
			return
		}

		resp, err := l.Action(&req, file, fileHeader)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
