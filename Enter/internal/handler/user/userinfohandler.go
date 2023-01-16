package user

import (
	"DouYin/utils"
	"net/http"

	"DouYin/Enter/internal/logic/user"
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 验证 Token
		if req.Token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token为空"))
			return
		}
		userClaim, err := utils.DecodeToken(req.Token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token错误"))
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req, userClaim)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
