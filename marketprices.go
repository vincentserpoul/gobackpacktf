package gobackpacktf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Result is the raw result of the json
type Result struct {
	Response `json:"response"`
}

// Response is the main body of the response
type Response struct {
	Success     int                  `json:"success"`
	Message     string               `json:"message"`
	CurrentTime int64                `json:"current_time"`
	Items       map[string]ItemPrice `json:"items"`
}

// ItemPrice are the items keyed by market hash names. quantity is the qty in the market
type ItemPrice struct {
	LastUpdated int64 `json:"last_updated"`
	Quantity    int   `json:"quantity"`
	Value       int   `json:"value"`
}

// BackpacktfAPIURLProduction is the existing API URL for backpacktf
const BackpacktfAPIURLProduction = "http://backpack.tf/api/IGetMarketPrices/v1"

// BackpacktfAPIURL is a global variable that contains the backpacktf URL
var BackpacktfAPIURL = BackpacktfAPIURLProduction

// GetMarketPrices will retrieve prices from the url
func GetMarketPrices(
	apiKey string,
	appID uint32,
) (*map[string]ItemPrice, error) {

	if apiKey == "" || appID == 0 {
		return nil, fmt.Errorf("gobackpacktf GetMarketPrices no parameter can be empty")
	}

	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("appid", strconv.FormatUint(uint64(appID), 10))

	// regular API string is http://backpack.tf/api/IGetMarketPrices/v1/
	resp, err := http.Get(BackpacktfAPIURL + "?" + querystring.Encode())

	if err != nil {
		return nil, fmt.Errorf("gobackpacktf GetMarketPrices http.Get: %v", err)
	}

	defer resp.Body.Close()

	mPrices := Result{}

	err = json.NewDecoder(resp.Body).Decode(&mPrices)
	if err != nil {
		return nil, fmt.Errorf("gobackpacktf GetMarketPrices Decode: %v", err)
	}

	if mPrices.Success != 1 {
		return nil, fmt.Errorf("gobackpacktf GetMarketPrices not successful: %s", mPrices.Message)
	}

	return &mPrices.Items, nil
}
