package middleware

import (
	"net/http"
	"strings"

	"DouYin/global"
	"github.com/zeromicro/go-zero/core/jsonx"
	"go.uber.org/zap"
)

// Logger 全局日志中间件
func Logger(spendTime string, r *http.Request, datasize string, form map[string][]string) {
	path := r.RequestURI
	method := r.Method
	ip := r.RemoteAddr
	host := r.Host
	formV := make(map[string]string, len(form))
	for k, v := range form {
		formV[k] = strings.Join(v, " ")
	}
	marshal, _ := jsonx.Marshal(formV)
	global.ZAP.Info(path,
		zap.String("Method:", method),
		// zap.String("Status:", status),
		zap.String("SpendTime", spendTime),
		zap.String("Ip:", ip),
		zap.String("Host:", host),
		zap.String("DataSize", datasize),
		zap.String("RawQuery:", string(marshal)),
	)
}
