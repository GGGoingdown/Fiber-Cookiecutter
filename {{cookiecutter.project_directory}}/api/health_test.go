package api

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthCheckAPI(t *testing.T) {
	endpoint := "/health"

	server := newTestServer(t)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	require.NoError(t, err)
	res, err := server.api.Test(request, -1)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
