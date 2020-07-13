package setup

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
	"rump/internal/logging"
	"rump/internal/server"
	"rump/internal/srvenv"
)

func Setup(ctx context.Context, config interface{}) (*srvenv.SrvEnv, error) {
	logger := logging.FromContext(ctx)
	l := envconfig.OsLookuper()
	env := &srvenv.SrvEnv{}
	if srvConf, ok := config.(*server.Config); ok {
		logger.Info("конфигурируется сервер")
		if err := envconfig.ProcessWith(ctx, srvConf, l); err != nil {
			return nil, fmt.Errorf("не обрабатывается env сервера: %w", err)
		}
		env.SrvConfig = srvConf
	}
	return env, nil
}
