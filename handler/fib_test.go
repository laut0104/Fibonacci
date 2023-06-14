package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Handler を Test すべきなのかそれとも getFib()のみテストすればいいのか
func TestGetFib(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		expectedBody   string
	}{
		//改行文字入るのどうにかならないの？
		{"Test Case 1", "GET", "https://783d-43-234-17-93.ngrok-free.app/fib", http.StatusBadRequest, http.StatusText(http.StatusBadRequest) + "\n"},
		{"Test Case 2", "GET", "https://783d-43-234-17-93.ngrok-free.app/fib?n=1", http.StatusOK, `{"result":1}`},
		{"Test Case 3", "GET", "https://783d-43-234-17-93.ngrok-free.app/fib?n=99", http.StatusOK, `{"result":218922995834555169026}`},
		{"Test Case 4", "GET", "https://783d-43-234-17-93.ngrok-free.app/fib?n=0", http.StatusBadRequest, http.StatusText(http.StatusBadRequest) + "\n"},
		{"Test Case 5", "GET", "https://783d-43-234-17-93.ngrok-free.app/fib?n=-1", http.StatusBadRequest, http.StatusText(http.StatusBadRequest) + "\n"},
		{"Test Case 6", "POST", "https://783d-43-234-17-93.ngrok-free.app/fib", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed) + "\n"},
		// {"Test Case 3", "GET", "https://783d-43-234-17-93.ngrok-free.app/unknown", http.StatusNotFound, http.StatusText(http.StatusNotFound)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Fib)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, rr.Body.String())
			}
		})
	}
}
