package gobackpacktf

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetMarketPrices(t *testing.T) {
	apiKey := "123"
	appID := uint32(730)
	expectedValidItemPrices := map[string]ItemPrice{
		"AK-47 | Aquamarine Revenge (Factory New)":              ItemPrice{LastUpdated: int64(1445860816), Quantity: int(27), Value: int(5516)},
		"AK-47 | Cartel (Well-Worn)":                            ItemPrice{LastUpdated: int64(1445860816), Quantity: int(53), Value: int(306)},
		"★ StatTrak™ Shadow Daggers | Urban Masked (Well-Worn)": ItemPrice{LastUpdated: int64(1445862648), Quantity: int(2), Value: int(12690)},
		"AK-47 | Aquamarine Revenge (Battle-Scarred)":           ItemPrice{LastUpdated: int64(1445860816), Quantity: int(85), Value: int(1249)},
	}

	cases := []struct {
		apiKey             string
		mock               func() string
		expectedErr        error
		expectedItemPrices *map[string]ItemPrice
	}{
		{apiKey: apiKey, mock: GetMockIGetMarketPrices, expectedErr: nil, expectedItemPrices: &expectedValidItemPrices},
		{apiKey: apiKey, mock: GetMockKOIGetMarketPrices, expectedErr: errors.New("not nil"), expectedItemPrices: nil},
		{apiKey: apiKey, mock: GetMockBadJSONIGetMarketPrices, expectedErr: errors.New("not nil"), expectedItemPrices: nil},
		{apiKey: "", mock: GetMockBadJSONIGetMarketPrices, expectedErr: errors.New("not nil"), expectedItemPrices: nil},
	}

loopcases:
	for _, c := range cases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, c.mock())
		}))
		defer ts.Close()
		BackpacktfAPIURL = ts.URL
		itemPrices, err := GetMarketPrices(c.apiKey, appID)

		if c.expectedErr != nil {
			if err == nil {
				t.Errorf("GetMarketPrices(%s) didn't return an error whereas it should have.", c.mock())
			}
			continue loopcases
		}
		if !reflect.DeepEqual(c.expectedItemPrices, itemPrices) {
			t.Errorf("GetMarketPrices(%s): returned \n%v\ninstead of\n%v\n", c.mock(), itemPrices, c.expectedItemPrices)
			continue loopcases
		}
	}
}

func TestTimeOutGetMarketPrices(t *testing.T) {
	apiKey := "123"
	appID := uint32(730)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		time.Sleep(200 * time.Millisecond)
	}))
	defer ts.Close()
	ts.Config.WriteTimeout = 20 * time.Millisecond

	BackpacktfAPIURL = ts.URL
	_, err := GetMarketPrices(apiKey, appID)

	if err == nil {
		t.Errorf("GetMarketPrices(): didn't trigger an error whereas it should have")
	}
}
