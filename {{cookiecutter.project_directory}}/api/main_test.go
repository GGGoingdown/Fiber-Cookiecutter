package api

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/GGGoingdown/Fiber-Cookiecutter/config"
	"github.com/GGGoingdown/Fiber-Cookiecutter/utils"
)

func newTestServer(t *testing.T) *Server {
	cfg, err := config.NewConfig("../")
	require.NoError(t, err)
	logger := utils.NewLogHandler("../storage/logs/", cfg.ToZapLogLevel())
	server := NewServer(cfg, logger)
	return server
}
