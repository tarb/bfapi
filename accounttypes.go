package bfapi

// AccountFundsResponse - response from getAccountFunds
type AccountFundsResponse struct {
	AvailableBalance   float64 `json:"availableToBetBalance"` //Amount available to bet.
	Exposure           float64 `json:"exposure"`              //Current exposure.
	RetainedCommission float64 `json:"retainedCommission"`    //Sum of retained commission.
	ExposureLimit      float64 `json:"exposureLimit"`         //Exposure limit.
	DiscountRate       float64 `json:"discountRate"`          //User Discount Rate.
	PointsBalance      int     `json:"pointsBalance"`         //The Betfair points balance
}

// AccountDetailsResponse - response from getAccountDetails
type AccountDetailsResponse struct {
	CurrencyCode  string  `json:"currencyCode"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	LocaleCode    string  `json:"localeCode"`
	Region        string  `json:"region"`
	Timezone      string  `json:"timezone"`
	DiscountRate  float64 `json:"discountRate"`
	PointsBalance int     `json:"pointsBalance"`
	CountryCode   string  `json:"countryCode"`
}

//
type CurrencyRate struct {
	CurrencyCode string  `json:"currencyCode"`
	Rate         float64 `json:"rate"`
}

//
type DeveloperApp struct {
	AppName     string                `json:"appName"`
	AppID       int64                 `json:"appID"`
	AppVersions []DeveloperAppVersion `json:"appVersions"`
}

//
type DeveloperAppVersion struct {
	Owner                string `json:"owner"`
	VersionID            int64  `json:"versionId"`
	Version              string `json:"version"`
	ApplicationKey       string `json:"applicationKey"`
	DelayData            bool   `json:"delayData"`
	SubscriptionRequired bool   `json:"subscriptionRequired"`
	OwnerManaged         bool   `json:"ownerManaged"`
	Active               bool   `json:"active"`
	VendorID             string `json:"vendorId"`
	VendorSecret         string `json:"vendorSecret"`
}

//
type ApplicationSubscription struct {
	SubscriptionToken    string    `json:"subscriptionToken"`
	ExpiryDateTime       Time      `json:"expiryDateTime"`
	ExpiredDateTime      Time      `json:"expiredDateTime"`
	CreatedDateTime      Time      `json:"createdDateTime"`
	ActivationDateTime   Time      `json:"activationDateTime"`
	CancellationDateTime Time      `json:"cancellationDateTime"`
	SubscriptionStatus   SubStatus `json:"subscriptionStatus"` //ALL, ACTIVATED, UNACTIVATED, CANCELLED, EXPIRED
	ClientReference      string    `json:"clientReference"`
	VendorClientID       string    `json:"vendorClientId"`
}

//
type TokenArg struct {
	ClientID     string    `json:"client_id"`
	Code         string    `json:"code,omitempty"`
	ClientSecret string    `json:"client_secret"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	GrantType    GrantType `json:"grant_type"`
}

//
type VendorAccessTokenInfo struct {
	AccessToken             string                  `json:"access_token"`
	ExpiresIn               int64                   `json:"expires_in"`
	RefreshToken            string                  `json:"refresh_token"`
	TokenType               TokenType               `json:"token_type"`
	ApplicationSubscription ApplicationSubscription `json:"application_subscription"`
}

// SubscriptionHistoryItem - response from getApplicationSubscriptionHistory
type SubscriptionHistoryItem struct {
	SubscriptionToken    string    `json:"subscriptionToken,omitempty"`    //Application key identifier
	ExpiryDateTime       Time      `json:"expiryDateTime,omitempty"`       //Subscription Expiry date
	ExpiredDateTime      Time      `json:"expiredDateTime,omitempty"`      //Subscription Expired date
	CreatedDateTime      Time      `json:"createdDateTime,omitempty"`      //Subscription Create date
	ActivationDateTime   Time      `json:"activationDateTime,omitempty"`   //Subscription Activation date
	CancellationDateTime Time      `json:"cancellationDateTime,omitempty"` //Subscription Cancelled date
	SubscriptionStatus   SubStatus `json:"subscriptionStatus,omitempty"`   //Subscription status {ALL, ACTIVATED, UNACTIVATED, CANCELLED, EXPIRED}
	ClientReference      string    `json:"clientReference,omitempty"`      //Client reference
}

//
type SubscriptionHistory []SubscriptionHistoryItem

//
func (sh SubscriptionHistory) GetBest() SubscriptionHistoryItem {

	if len(sh) == 0 {
		return SubscriptionHistoryItem{SubscriptionStatus: SubStatusUnactivated}
	}

	best := sh[0]
	for _, item := range sh[1:] {

		if (item.SubscriptionStatus == SubStatusActivated &&
			(best.SubscriptionStatus != SubStatusActivated || item.ExpiryDateTime > best.ExpiryDateTime)) ||
			(best.SubscriptionStatus != SubStatusActivated &&
				item.ExpiredDateTime+item.CancellationDateTime > best.ExpiredDateTime+best.CancellationDateTime) {

			best = item
		}
	}

	return best
}

//
func (sh SubscriptionHistory) Filter(fn func(i SubscriptionHistoryItem) bool) SubscriptionHistory {
	result := make(SubscriptionHistory, 0, len(sh))
	for _, i := range sh {
		if fn(i) {
			result = append(result, i)
		}
	}
	return result
}
