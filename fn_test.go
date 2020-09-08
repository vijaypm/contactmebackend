package contactmebackend

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSendEmail(t *testing.T) {
	// Mock data with the 'body' which will go over the REST, and 'wantStatus' which is the expected status for the data sent in 'body' field
	form := url.Values{}
	form.Add("name", "Dummy")
	form.Add("email", "dummy@example.com")
	form.Add("message", "Would love to meet over a beer!")
	// httptest.NewRequest package will form the mock request object which we can pass on to our function directly
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// mock response object
	rr := httptest.NewRecorder()
	SendEmail(rr, req)
	if got := rr.Result().StatusCode; got != http.StatusOK {
		t.Errorf("SendEmail(%q) = %d, want %d", form, got, http.StatusOK)
	}
}
