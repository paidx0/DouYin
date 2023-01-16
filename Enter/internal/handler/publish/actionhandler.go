package publish

import (
	"net/http"

	"DouYin/Enter/internal/logic/publish"
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := publish.NewActionLogic(r.Context(), svcCtx)
		resp, err := l.Action(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
