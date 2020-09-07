package contactmebackend

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	// Mock data with the 'body' which will go over the REST, and 'wantStatus' which is the expected status for the data sent in 'body' field
	tests := []struct {
		body       string
		wantStatus int
	}{
		// This one has all the necessary fields for successful request
		{body: `{"name": "Dummy", "email": "dummy@example.com", "message": "Would love to meet over a beer!"}`, wantStatus: http.StatusOK},
		// This one will fail the request because emailAddress is missing.
		{body: `{"name": "Dummy", "message": "Would love to meet over a beer!"}`, wantStatus: http.StatusOK},
	}
	for _, test := range tests {
		// httptest.NewRequest package will form the mock request object which we can pass on to our function directly
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))
		req.Header.Add("Content-Type", "application/json")
		// mock response object
		rr := httptest.NewRecorder()
		SendEmail(rr, req)
		if got := rr.Result().StatusCode; got != test.wantStatus {
			t.Errorf("SendEmail(%q) = %d, want %d", test.body, got, test.wantStatus)
		}
	}
}
