package main

import (
	"context"

	"github.com/i-nas/i-nas-tools/pkg/http"
	"github.com/i-nas/i-nas-tools/pkg/lifecycle"
	"github.com/i-nas/i-nas-tools/pkg/log"
)

func main() {
	log.Info("start app...")
	ctx := context.Background()
	defer lifecycle.BeforeExit()
	lifecycle.Init()
	http.StartWithCtx(ctx)
}
