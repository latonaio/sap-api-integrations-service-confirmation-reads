package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceConfirmation           string      `json:"ServiceConfirmation"`
			ServiceConfirmationItem       string      `json:"ServiceConfirmationItem"`
			ServiceConfItemDescription    string      `json:"ServiceConfItemDescription"`
			ServiceObjectType             string      `json:"ServiceObjectType"`
			ServiceDocumentItemObjectType string      `json:"ServiceDocumentItemObjectType"`
			Language                      string      `json:"Language"`
			Product                       string      `json:"Product"`
			Quantity                      string      `json:"Quantity"`
			QuantityUnit                  string      `json:"QuantityUnit"`
			ActualServiceDuration         string      `json:"ActualServiceDuration"`
			ActualServiceDurationUnit     string      `json:"ActualServiceDurationUnit"`
			ServiceConfItemCategory       string      `json:"ServiceConfItemCategory"`
			ServiceConfItemIsCompleted    string      `json:"ServiceConfItemIsCompleted"`
			ExecutingServiceEmployee      string      `json:"ExecutingServiceEmployee"`
			ServicePerformer              string      `json:"ServicePerformer"`
			BillableControl               string      `json:"BillableControl"`
			ReferenceServiceOrder         string      `json:"ReferenceServiceOrder"`
			ReferenceServiceOrderItem     string      `json:"ReferenceServiceOrderItem"`
			ParentServiceConfItem         string      `json:"ParentServiceConfItem"`
			ActualServiceStartDateTime    string      `json:"ActualServiceStartDateTime"`
			ActualServiceEndDateTime      string      `json:"ActualServiceEndDateTime"`
			ServicesRenderedDate          string      `json:"ServicesRenderedDate"`
			TimeSheetOvertimeCategory     string      `json:"TimeSheetOvertimeCategory"`
			RespyMgmtServiceTeam          string      `json:"RespyMgmtServiceTeam"`
			RespyMgmtServiceTeamName      string      `json:"RespyMgmtServiceTeamName"`
			ToItemPricingElement              struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PricingElement"`
		} `json:"results"`
	} `json:"d"`
}
