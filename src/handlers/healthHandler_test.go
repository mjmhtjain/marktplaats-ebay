package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {

	t.Run("IF handler is working, THEN expect 200 response", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/health", nil)

		HealthHandler(w, req)
		assert.Equal(t, w.Code, http.StatusOK)
	})
}
