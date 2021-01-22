package bfapi

import (
	"errors"
	"net/http"

	"github.com/tarb/util/www"
)

//
func (c *Client) ListMarketCatalogue(lmc ListMarketCatalogueArg) ([]MarketCatalogue, error) {
	var result []MarketCatalogue
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listMarketCatalogue).
		WithJSONBody(lmc).
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
func (c *Client) ListMarketBook(lmb ListMarketBookArg) ([]MarketBook, error) {
	var result []MarketBook
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listMarketBook).
		WithJSONBody(lmb).
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
func (c *Client) PlaceOrders(arg PlaceOrderArg) (PlaceExecutionReport, error) {
	var result PlaceExecutionReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, placeOrders).
		WithJSONBody(arg).
		WithHeaders(func(h http.Header) {
			h.Set(c.GetAuth())
			h.Set("X-Application", c.appKey)
			h.Set("Accept", "application/json")
			h.Set("Connection", "keep-alive")
			h.Set("Content-Type", "application/json; charset=utf-8")
		}).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&result)

	if result.Status == "FAILURE" {
		return result, errors.New(result.ErrorCode)
	}

	err = statusToAPINGException(err)

	return result, err
}

//
func (c *Client) CancelOrders(req CancelOrderRequest) (CancelExecutionReport, error) {
	var result CancelExecutionReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, cancelOrders).
		WithJSONBody(req).
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
func (c *Client) ReplaceOrders(req ReplaceOrderRequest) (ReplaceExecutionReport, error) {
	var result ReplaceExecutionReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, replaceOrders).
		WithJSONBody(req).
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
func (c *Client) UpdateOrders(ins UpdateOrderRequest) (UpdateExecutionReport, error) {
	var result UpdateExecutionReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, updateOrders).
		WithJSONBody(ins).
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
func (c *Client) ListClearedOrders(arg ListClearedOrdersArgs) (ClearedOrderSummaryReport, error) {
	var result ClearedOrderSummaryReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listClearedOrders).
		WithJSONBody(arg).
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
func (c *Client) ListCurrentOrders(arg ListCurrentOrdersArgs) (CurrentOrderSummaryReport, error) {
	var result CurrentOrderSummaryReport
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listCurrentOrders).
		WithJSONBody(arg).
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
func (c *Client) ListMarketTypes(arg ListArg) ([]MarketTypeResult, error) {
	var result []MarketTypeResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listMarketTypes).
		WithJSONBody(arg).
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
func (c *Client) ListEventTypes(arg ListArg) ([]EventTypeResult, error) {
	var result []EventTypeResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listEventTypes).
		WithJSONBody(arg).
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
func (c *Client) ListCompetitions(arg ListArg) ([]CompetitionResult, error) {
	var result []CompetitionResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listCompetitions).
		WithJSONBody(arg).
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
func (c *Client) ListEvents(arg ListArg) ([]EventResult, error) {
	var result []EventResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listEvents).
		WithJSONBody(arg).
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
func (c *Client) ListCountries(arg ListArg) ([]CountryCodeResult, error) {
	var result []CountryCodeResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listCountries).
		WithJSONBody(arg).
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
func (c *Client) ListVenues(arg ListArg) ([]VenueResult, error) {
	var result []VenueResult
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listVenues).
		WithJSONBody(arg).
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
func (c *Client) ListMarketProfitAndLoss(arg ListMarketProfitAndLossArg) (*MarketProfitAndLoss, error) {
	var result MarketProfitAndLoss
	err := c.client.Build(http.MethodPost, scheme, exchangeHost, listMarketProfitAndLoss).
		WithJSONBody(arg).
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

	return &result, err
}
