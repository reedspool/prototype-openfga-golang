package oauth

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"golang.org/x/oauth2"
)

func testGenerateStateToken() string {
	return "TestStatePleaseIgnore"
}

func TestActualGenerateStateToken(t *testing.T) {
	state := GenerateStateToken()
	if len(state) < 10 {
		t.Errorf("state not long enough: %q", state)
	}

	if match, err := regexp.MatchString("[^a-zA-Z]", state); match || err != nil {
		t.Errorf("Err? %v or state contained non-alphabetic characters: %q", err, state)
	}
}

func TestOauthLoginFinalize(t *testing.T) {
	t.Errorf("Untested")
}

func TestOauthLoginRedirect(t *testing.T) {
	mockOauthProviderServer, oauthConfig := createMockOauthServer()
	defer mockOauthProviderServer.Close()

	// Start my server
	myOauth := MyOauth{Config: oauthConfig, GenerateStateToken: testGenerateStateToken}
	myServer := httptest.NewServer(http.HandlerFunc(myOauth.MyLoginHandler))
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}}

	// Expect the login response will have a redirect
	loginResponse, err := client.Get(myServer.URL + "/login")
	if err != nil {
		t.Fatalf("Unexpected error on login: %v", err)
	}
	if loginResponse.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("/login status not OK: %v", loginResponse.Status)
	}
	// location = http://127.0.0.1:42467/auth?access_type=offline&client_id=test-client-id&response_type=code&state=testtest
	location := loginResponse.Header.Get("location")
	url, err := url.Parse(location)
	state := url.Query().Get("state")

	if match, err := regexp.MatchString(oauthConfig.Endpoint.AuthURL, location); !match || err != nil {
		t.Errorf("Err?: %v OR redirected to unexpected place: %v", err, location)
	}

	if state != testGenerateStateToken() {
		t.Errorf("Unexpected state parameter: %v", state)
	}
}

func TestMockOauthServerWithCoreOauth2(t *testing.T) {
	mockServer, config := createMockOauthServer()
	defer mockServer.Close()

	// Test token exchange
	token, err := config.Exchange(context.Background(), "mock-code")
	if err != nil {
		t.Fatalf("Exchange failed: %v", err)
	}

	if token.AccessToken != "mock-access-token" {
		t.Errorf("Expected mock-access-token, got %s", token.AccessToken)
	}
}

func TestMockOauthServer(t *testing.T) {
	mockServer, oauthConfig := createMockOauthServer()
	defer mockServer.Close()

	resp, err := http.Get(oauthConfig.Endpoint.TokenURL)

	if err != nil {
		t.Fatalf("Failed to Get from mock server: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Got non-200 status from mock server: %v", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected json content type but got: %v", resp.Header.Get("Content-Type"))
	}

	b, err := io.ReadAll(resp.Body)
	if match, err := regexp.MatchString("access_token", string(b)); !match {
		t.Errorf("regex matching body error: %v", err)
	}
}

func createMockOauthServer() (*httptest.Server, *oauth2.Config) {
	// Create mock OAuth2 server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		if r.URL.Path == "/token" {
			bufio.NewWriter(w).WriteString("Hello!")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"access_token":  "mock-access-token",
				"refresh_token": "mock-refresh-token",
				"token_type":    "Bearer",
				"expires_in":    3600,
			})
			return
		}

		if r.URL.Path == "/userinfo" {
			json.NewEncoder(w).Encode(map[string]string{
				"email": "test@example.com",
				"name":  "Test User",
			})
			return
		}
	}))

	config := &oauth2.Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		Endpoint: oauth2.Endpoint{
			AuthURL:  mockServer.URL + "/auth",
			TokenURL: mockServer.URL + "/token",
		},
	}

	return mockServer, config
}
