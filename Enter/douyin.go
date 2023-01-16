package main

import (
	"DouYin/Enter/internal/middleware"
	"flag"
	"fmt"
	"math"
	"net/http"
	"time"

	"DouYin/Enter/internal/config"
	"DouYin/Enter/internal/handler"
	"DouYin/Enter/internal/svc"
	_ "DouYin/global"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/douyin-api.yaml", "the global file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next(w, r)
			stop := time.Since(start)
			spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stop.Nanoseconds())/1000000.0)))
			datasize := r.Header.Get("Content-Length")
			form := r.Form
			middleware.Logger(spendTime, r, datasize, form)
		}
	})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
