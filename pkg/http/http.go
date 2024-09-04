package http

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/i-nas/i-nas-tools/pkg/lifecycle"
	"github.com/i-nas/i-nas-tools/pkg/log"
)

func StartWithCtx(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "up",
		})
	})
	//创建HTTP服务器
	addr := os.Getenv("I_NAS_TOOLS_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	lifecycle.RegisterBeforeExit(
		func() {
			server.Shutdown(ctx)
		},
	)
	//启动HTTP服务器
	log.Infof("ListenAndServe %s", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("err: %s", err)
	}
}

func Start() {
	StartWithCtx(context.Background())

}
