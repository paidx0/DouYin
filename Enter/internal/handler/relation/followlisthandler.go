package relation

import (
	"net/http"

	"DouYin/Enter/internal/logic/relation"
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := relation.NewFollowlistLogic(r.Context(), svcCtx)
		resp, err := l.Followlist(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
