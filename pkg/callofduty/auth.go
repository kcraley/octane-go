package callofduty

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	identityURL   = "https://www.callofduty.com/api/papi-client/crm/cod/v1/identities"
	loginTokenURL = "https://profile.callofduty.com/login"
	loginPostURL  = "https://profile.callofduty.com/do_login?new_SiteId=cod"
)

// Authentication contains the necessary cookie values that
// are required to be passed with every request
type Authentication struct {
	Username  string
	Password  string
	ATKN      *http.Cookie
	SSO       *http.Cookie
	SSOExpiry *http.Cookie
	XSRF      *http.Cookie
}

// SessionExpired validates that the authentication session is not expired
func (a *Authentication) SessionExpired() bool {
	// Check if the SSOExpiry value is greater than the existing
	// epoch time value.  These values are converted and compared
	// in epoch milliseconds.
	expiry, _ := strconv.Atoi(a.SSOExpiry.Value)
	now := int(time.Now().UnixNano() / 1000000)
	if now > expiry {
		return true
	}
	return false
}

// NewSession re-authenticates with the Call of Duty site
func (a *Authentication) NewSession() error {
	// Authenticate and create a new session
	err := a.login()
	if err != nil {
		return err
	}
	return nil
}

// login authenticates to the Call of Duty site with
// a username, password and the XSRF-Token Cookie
func (a *Authentication) login() error {
	var err error

	// Get the XSRF Token
	a.XSRF, err = getXSRFToken()
	if err != nil {
		return err
	}

	// Create a new http client with a non-nil cookie jar
	cookieJar, _ := cookiejar.New(
		&cookiejar.Options{
			PublicSuffixList: publicsuffix.List,
		},
	)
	client := &http.Client{
		Jar:     cookieJar,
		Timeout: 10 * time.Second,
	}

	// Create the payload to send with login information
	form := url.Values{}
	form.Add("username", a.Username)
	form.Add("password", a.Password)
	form.Add("remember_me", "true")
	form.Add("_csrf", a.XSRF.Value)

	// Construct the request with payload and cookie
	req, err := http.NewRequest("POST", loginPostURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.AddCookie(a.XSRF)
	req.AddCookie(&http.Cookie{Name: "new_SiteId", Value: "cod"})
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execute the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Get the necessary cookies for future requests
	for _, cookie := range cookieJar.Cookies(req.URL) {
		switch cookie.Name {
		case "ACT_SSO_COOKIE":
			a.SSO = cookie
		case "ACT_SSO_COOKIE_EXPIRY":
			a.SSOExpiry = cookie
		case "atkn":
			a.ATKN = cookie
		}
	}

	// Return the authentication type with the required values
	return nil
}

// getXSRFToken makes an initial request to the Call of Duty
// login page in order to retrieve the XSRF-TOKEN Cookie
func getXSRFToken() (*http.Cookie, error) {
	// Create default empty http.Cookie
	emtpyCookie := &http.Cookie{}

	// Create a new http client
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequest("GET", loginTokenURL, nil)
	if err != nil {
		return emtpyCookie, err
	}

	// Add empty Header
	req.Header.Set("Cookie", "")

	// Submit the request
	resp, err := client.Do(req)
	if err != nil {
		return emtpyCookie, err
	}
	defer resp.Body.Close()

	// Find and return the XSRF Header
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "XSRF-TOKEN" {
			return cookie, nil
		}
	}
	// Return a default of nothing if not found
	return emtpyCookie, fmt.Errorf("XSRF-TOKEN not found")
}
