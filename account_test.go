package bfapi

import "testing"

func TestGetBest(t *testing.T) {

	items := SubscriptionHistory{
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 0, ExpiredDateTime: 5},
		{SubscriptionStatus: SubStatusCancelled, ExpiryDateTime: 0, CancellationDateTime: 6, ExpiredDateTime: 0},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 10},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 15},
		{SubscriptionStatus: SubStatusActivated, ExpiryDateTime: 20, CancellationDateTime: 0, ExpiredDateTime: 0},
	}

	item := items.GetBest()
	if item != items[4] {
		t.Fail()
	}

	items = SubscriptionHistory{
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 0, ExpiredDateTime: 5},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 15},
		{SubscriptionStatus: SubStatusCancelled, ExpiryDateTime: 0, CancellationDateTime: 6, ExpiredDateTime: 0},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 10},
	}

	item = items.GetBest()
	if item != items[1] {
		t.Fail()
	}

	items = SubscriptionHistory{
		{SubscriptionStatus: SubStatusActivated, ExpiryDateTime: 10, CancellationDateTime: 0, ExpiredDateTime: 0},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 0, ExpiredDateTime: 5},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 15},
		{SubscriptionStatus: SubStatusCancelled, ExpiryDateTime: 0, CancellationDateTime: 21, ExpiredDateTime: 0},
		{SubscriptionStatus: SubStatusExpired, ExpiryDateTime: 0, CancellationDateTime: 5, ExpiredDateTime: 10},
	}

	item = items.GetBest()
	if item != items[0] {
		t.Fail()
	}

}
