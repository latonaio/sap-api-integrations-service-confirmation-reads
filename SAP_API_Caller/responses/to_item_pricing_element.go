package responses

type ToItemPricingElement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceConfirmation     string `json:"ServiceConfirmation"`
			ServiceConfirmationItem string `json:"ServiceConfirmationItem"`
			PricingProcedureStep    string `json:"PricingProcedureStep"`
			PricingProcedureCounter string `json:"PricingProcedureCounter"`
			ConditionType           string `json:"ConditionType"`
			ConditionRateValue      string `json:"ConditionRateValue"`
			ConditionCurrency       string `json:"ConditionCurrency"`
			ConditionQuantity       string `json:"ConditionQuantity"`
			ConditionQuantityUnit   string `json:"ConditionQuantityUnit"`
		} `json:"results"`
	} `json:"d"`
}
