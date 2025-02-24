package blog

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHomeHandler(t *testing.T) {
	posts := []BlogPost{
		{
			Title:                "First Post",
			Excerpt:              "This is the first post",
			Date:                 time.Date(2024, 2, 24, 12, 0, 0, 0, time.UTC),
			ReadingTimeInSeconds: 60,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := NewHomeHandler{posts}
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("want status %d; got %d", http.StatusOK, w.Code)
	}
}
