package bfapi

import (
	json "encoding/json"
	"net/http"
	"time"

	"github.com/tarb/util/www"
)

//
type Menu struct {
	Parent     *Menu
	Decendants int

	ID              json.Number `json:"id"`
	ExchangeID      string      `json:"exchangeId"`
	Type            string      `json:"type"`
	CountryCode     string      `json:"countryCode"`
	MarketStartTime time.Time   `json:"marketStartTime"`
	MarketType      string      `json:"marketType"`
	Name            string      `json:"name"`
	Children        []*Menu     `json:"children"`
}

// GetMenu - Collects JSON Market Menu from Betfair API
func (c *Client) GetMenu() (*Menu, error) {
	var m Menu
	err := c.client.Build(http.MethodGet, scheme, exchangeHost, getMarketMenuJSON).
		DoWithRetry(5, www.LinearJitterDelay).
		CollectJSON(&m)
	if err != nil {
		return nil, err
	}

	var process func(*Menu, *Menu) int
	process = func(m *Menu, p *Menu) int {
		count := 0

		if m.Children != nil {
			for i := range m.Children {
				count += process(m.Children[i], m)
			}
		} else {
			count = 1
		}

		m.Parent = p
		m.Decendants = count
		return count
	}
	process(&m, nil)

	return &m, err
}

// NumMarkets -
func (m *Menu) NumMarkets() int {
	count := 0

	if m.Children != nil {
		for i := range m.Children {
			count += m.Children[i].NumMarkets()
		}
	} else {
		count = 1
	}

	return count
}

// WalkAll -
func (m *Menu) WalkAll(fn func(int, *Menu)) {
	var walk func(*Menu, int, func(int, *Menu))
	walk = func(m *Menu, lvl int, fn func(int, *Menu)) {
		fn(lvl, m)
		if m.Children != nil {
			for i := range m.Children {
				walk(m.Children[i], lvl+1, fn)
			}
		}
	}
	walk(m, 0, fn)
}

// WalkIf -
func (m *Menu) WalkIf(fn func(int, *Menu) bool) {
	var walk func(*Menu, int, func(int, *Menu) bool)
	walk = func(m *Menu, lvl int, fn func(int, *Menu) bool) {
		if fn(lvl, m) && m.Children != nil {
			for i := range m.Children {
				walk(m.Children[i], lvl+1, fn)
			}
		}
	}
	walk(m, 0, fn)
}

// Filter maybe rewrite this so that it handles back processing at the same time
func (m *Menu) Filter(fn func(*Menu) bool) *Menu {
	var copy Menu = *m
	var c *Menu = &copy

	if len(m.Children) > 0 {
		c.Children = nil

		for i := range m.Children {
			cm := m.Children[i].Filter(fn)
			if cm != nil {
				c.Children = append(c.Children, cm)
			}
		}

		if len(c.Children) == 0 {
			c = nil
		}
	} else {
		if !(fn(m)) {
			c = nil
		}
	}

	return c
}
