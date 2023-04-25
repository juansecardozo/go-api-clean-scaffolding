package bootstrap

import (
	"context"
	"github.com/juansecardozo/go-api-clean-scaffolding/internal/platform/server"
)

func Run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	ctx, srv := server.New(
		context.Background(),
		cfg.Host,
		cfg.Port,
		cfg.ShutdownTimeout,
	)
	return srv.Run(ctx)
}
