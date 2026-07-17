package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProfileHandler(t *testing.T) {
	t.Parallel()

	handler := profileHandler()

	t.Run("serves pprof index", func(t *testing.T) {
		t.Parallel()

		request := httptest.NewRequest(http.MethodGet, "/debug/pprof/", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)

		require.Equal(t, http.StatusOK, response.Code)
		require.Contains(t, response.Body.String(), "Types of profiles available")
	})

	t.Run("rejects unrelated paths", func(t *testing.T) {
		t.Parallel()

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)

		require.Equal(t, http.StatusNotFound, response.Code)
	})
}
