package gobackpacktf

// GetMockIGetMarketPrices will return a mock of http://backpack.tf/api/IGetMarketPrices/v1/
func GetMockIGetMarketPrices() string {
	return `
    {
      "response": {
        "success": 1,
        "current_time": 1445863117,
        "items": {
          "AK-47 | Aquamarine Revenge (Battle-Scarred)": {
            "last_updated": 1445860816,
            "quantity": 85,
            "value": 1249
          },
          "AK-47 | Aquamarine Revenge (Factory New)": {
            "last_updated": 1445860816,
            "quantity": 27,
            "value": 5516
          },
          "AK-47 | Cartel (Well-Worn)": {
            "last_updated": 1445860816,
            "quantity": 53,
            "value": 306
          },
          "★ StatTrak™ Shadow Daggers | Urban Masked (Well-Worn)": {
            "last_updated": 1445862648,
            "quantity": 2,
            "value": 12690
          }
        }
      }
    }
    `
}

// GetMockKOIGetMarketPrices will return a mock of http://backpack.tf/api/IGetMarketPrices/v1/
// when the wrong APIkey is used
func GetMockKOIGetMarketPrices() string {
	return `
    {
      "response": {
        "success": 0,
        "message": "This API key is not valid."
      }
    }
    `
}

// GetMockBadJSONIGetMarketPrices will return a mock of http://backpack.tf/api/IGetMarketPrices/v1/
func GetMockBadJSONIGetMarketPrices() string {
	return `
      "response": {
        "success": 0,
        "message": "This API key is not valid."
      }
    }
    `
}
