package bfapi

import (
	"errors"
	"net/http"

	"github.com/tarb/util/www"
)

//
func (c *Client) GetAccountFunds() (AccountFundsResponse, error) {
	var result AccountFundsResponse
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getAccountFunds).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) GetAccountDetails() (AccountDetailsResponse, error) {
	var result AccountDetailsResponse
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getAccountDetails).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) ListCurrencyRates() ([]CurrencyRate, error) {
	var result []CurrencyRate
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listCurrencyRates).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) GetApplicationSubscriptionToken(len int, ref string) (string, error) {
	type arg struct {
		SubscriptionLength int    `json:"subscriptionLength,omitempty"` //number of days
		ClientReference    string `json:"clientReference,omitempty"`
	}

	var result string
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getApplicationSubscriptionToken).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg{len, ref}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) GetVendorClientID() (string, error) {
	var result string
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getVendorClientID).
		WithHeaders(func(h http.Header) {
			if c.Token().Type == OAuthToken {
				h.Set("X-Application", c.appKey)
			}

			h.Set(c.GetAuth())
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) UpdateApplicationSubscription(vid string, len int) (string, error) {
	type arg struct {
		VCID string `json:"vendorClientId"`
		Len  int    `json:"subscriptionLength"`
	}

	var result string
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, updateApplicationSubscription).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg{vid, len}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) ActivateApplicationSubscription(token string) error {
	type arg struct {
		Key string `json:"subscriptionToken"`
	}

	var result string
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, activateApplicationSubscription).
		WithJSONBody(arg{token}).
		WithHeaders(func(h http.Header) {
			if c.Token().Type == OAuthToken {
				h.Set("X-Application", c.appKey)
			}

			h.Set(c.GetAuth())
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	if err != nil {
		err = statusToAPINGException(err)
		return err
	}

	if result != "SUCCESS" {
		return errors.New("activation failed")
	}

	return nil
}

// GetApplicationSubscriptionHistory returns the users subscription history
// vid - (required) the vendor client id of the user
// appKey - (optional) the appKey to test subscription against. If not specified the client appKey is used
func (c *Client) GetApplicationSubscriptionHistory(vid, appkey string) (SubscriptionHistory, error) {
	var arg = struct {
		VCID   string `json:"vendorClientId"`
		AppKey string `json:"applicationKey"`
	}{}
	arg.AppKey = appkey
	arg.VCID = vid

	var result SubscriptionHistory
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getApplicationSubscriptionHistory).
		WithHeaders(func(h http.Header) {
			if c.Token().Type == OAuthToken {
				h.Set("X-Application", c.appKey)
			}

			h.Set(c.GetAuth())
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) CreateDeveloperAppKeys(name string) (DeveloperApp, error) {
	type arg struct {
		AppName string `json:"appName"`
	}

	var result DeveloperApp
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, createDeveloperAppKeys).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg{name}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	var result []DeveloperApp
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getDeveloperAppKeys).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return result, err
}

// ListApplicationSubscriptionTokens has paramater values ALL, ACTIVATED, UNACTIVATED, CANCELLED, EXPIRED
func (c *Client) ListApplicationSubscriptionTokens(status string) ([]ApplicationSubscription, error) {
	var arg = struct {
		Status string `json:"subscriptionStatus,omitempty"`
	}{
		Status: status,
	}

	var result []ApplicationSubscription
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listApplicationSubscriptionTokens).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)
	return result, err
}

// GetToken - get web vendor token info
func (c *Client) GetToken(arg TokenArg) (*VendorAccessTokenInfo, error) {
	var result VendorAccessTokenInfo
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, getToken).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		WithJSONBody(arg).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)
	err = statusToAPINGException(err)

	return &result, err
}
