package bfapi

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/tarb/util/www"
)

//
const SessionTimeout = 8 * time.Hour

//
var (
	ErrBusy                = errors.New("login resource busy")
	ErrAlreadyLoggedIn     = errors.New("already logged in")
	ErrNotLoggedIn         = errors.New("not logged in")
	ErrSubExpired          = errors.New("subscription expired")
	ErrSubActivationFailed = errors.New("subscription activation failed")
)

//
func (c *Client) VendorLogin(username, password string) (Token, error) {
	// if logged in already then error out - require logout first
	if c.LoggedIn() {
		return Token{}, ErrAlreadyLoggedIn
	}

	if !c.sem.TryAcquire(1) {
		return Token{}, ErrBusy
	}
	defer c.sem.Release(1)

	token, err := c.login(username, password)
	if err != nil {
		return Token{}, err
	}

	t := Token{Type: SessionToken, Token: token}
	c.token.Store(t)

	vcid, err := c.GetVendorClientID()
	if err != nil {
		return Token{}, err
	}
	sh, err := c.GetApplicationSubscriptionHistory(vcid, c.appKey)
	if err != nil {
		return Token{}, err
	}

	// store the session token
	now := FromStdTime(time.Now())
	t = Token{Type: SessionToken, Token: token, Logged: now, Update: now, VcID: vcid, Sub: sh.GetBest()}
	c.token.Store(t)

	// Not necessarily login success
	// result.SessionToken can be empty here with non nil err value
	return t, nil
}

// Login is a singleflight call - logs in with the supplied username and password
// params are username, password - betfair account login details
func (c *Client) Login(username, password string) (Token, error) {
	// if logged in already then error out - require logout first
	if c.LoggedIn() {
		return Token{}, ErrAlreadyLoggedIn
	}

	if !c.sem.TryAcquire(1) {
		return Token{}, ErrBusy
	}
	defer c.sem.Release(1)

	token, err := c.login(username, password)
	if err != nil {
		return Token{}, err
	}

	now := FromStdTime(time.Now())
	t := Token{Type: SessionToken, Token: token, Logged: now, Update: now}
	c.token.Store(t)

	return t, nil
}

func (c *Client) login(username, password string) (string, error) {
	if c.certificate == nil {
		var result LoginResult
		err := c.client.Build(http.MethodPost, scheme, accountHost, login).
			WithFormBody(func(form url.Values) {
				form.Set("username", username)
				form.Set("password", password)
			}).
			WithHeaders(func(h http.Header) {
				h.Set("X-Application", c.appKey)
				h.Set("Accept", "application/json")
				h.Set("Connection", "keep-alive")
			}).
			DoWithRetry(5, www.LinearJitterDelay).
			CollectJSON(&result)

		if err != nil {
			return "", err
		} else if result.Status != StatusSuccess {
			return "", LoginError(result.Error)
		}

		return result.SessionToken, nil

	} else {
		var certResult CertLoginResult
		err := c.client.Build(http.MethodPost, scheme, accountHost, certLogin).
			WithFormBody(func(form url.Values) {
				form.Set("username", username)
				form.Set("password", password)
			}).
			WithHeaders(func(h http.Header) {
				h.Set("X-Application", c.appKey)
				h.Set("Accept", "application/json")
				h.Set("Connection", "keep-alive")
			}).
			DoWithRetry(5, www.LinearJitterDelay).
			CollectJSON(&certResult)

		if err != nil {
			return "", err
		} else if certResult.Status != StatusSuccess {
			return "", LoginError(certResult.Status)
		}

		return certResult.SessionToken, nil
	}
}

// KeepAlive error if subscription is no longer active for example
// return false, and err.Expired
func (c *Client) KeepAlive() error {
	if !c.sem.TryAcquire(1) {
		return ErrBusy
	}
	defer c.sem.Release(1)

	now := FromStdTime(time.Now())
	token := c.Token()

	var result LoginResult
	err := c.client.Build(http.MethodGet, scheme, accountHost, keepAlive).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	// was login successful
	if err != nil {
		return err
	} else if result.Status != "SUCCESS" {
		// Connection worked but keepAlive was not successful
		return errors.New("Session KeepAlive Failed. " + result.Error)
	}

	c.token.Store(Token{SessionToken, token.Token, token.Logged, now, token.VcID, token.Sub})

	return nil
}

//
func (c *Client) ActiveSub() (bool, error) {
	token := c.Token()
	now := FromStdTime(time.Now())

	if token.ActiveSub() {

		if token.Sub.ExpiryDateTime < now {
			sh, err := c.GetApplicationSubscriptionHistory(c.Token().VcID, c.appKey)
			if err != nil {
				return false, err
			}

			t := Token{SessionToken, token.Token, token.Logged, token.Update, token.VcID, sh.GetBest()}
			if active := t.ActiveSub(); !active {
				c.logout()
				return false, ErrSubExpired
			}

			c.token.Store(t)
			token = t
		}

	}

	return false, ErrSubExpired
}

//
func (c *Client) ActivateSub(token string) error {
	if !c.sem.TryAcquire(1) {
		return ErrBusy
	}
	defer c.sem.Release(1)

	t := c.Token()
	if err := c.ActivateApplicationSubscription(token); err != nil {
		return err
	}

	sh, err := c.GetApplicationSubscriptionHistory(t.VcID, c.appKey)
	if err != nil {
		return err
	}

	t = Token{t.Type, t.Token, t.Logged, t.Update, t.VcID, sh.GetBest()}

	if active := t.ActiveSub(); !active {
		return ErrSubActivationFailed
	}
	c.token.Store(t)

	return nil
}

//
func (c *Client) Logout() error {
	if !c.sem.TryAcquire(1) {
		return ErrBusy
	}
	defer c.sem.Release(1)

	return c.logout()
}

//
func (c *Client) logout() error {
	if !c.LoggedIn() {
		return ErrNotLoggedIn
	}

	var result LoginResult
	err := c.client.Build(http.MethodGet, scheme, accountHost, logout).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	if err != nil {
		return err
	} else if result.Status != "SUCCESS" {
		// Connection worked but logout was not successful
		return errors.New("Session Logout Failed. " + result.Error)
	}

	c.token.Store(Token{})

	return nil
}
