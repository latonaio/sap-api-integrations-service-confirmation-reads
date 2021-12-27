package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceConfirmation            string      `json:"ServiceConfirmation"`
			ServiceConfirmationType        string      `json:"ServiceConfirmationType"`
			ServiceConfirmationDescription string      `json:"ServiceConfirmationDescription"`
			ServiceObjectType              string      `json:"ServiceObjectType"`
			Language                       string      `json:"Language"`
			ServiceDocumentPriority        string      `json:"ServiceDocumentPriority"`
			RequestedServiceStartDateTime  string      `json:"RequestedServiceStartDateTime"`
			RequestedServiceEndDateTime    string      `json:"RequestedServiceEndDateTime"`
			PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
			ServiceConfirmationIsCompleted string      `json:"ServiceConfirmationIsCompleted"`
			ServiceConfirmationIsCanceled  bool        `json:"ServiceConfirmationIsCanceled"`
			SalesOrganization              string      `json:"SalesOrganization"`
			DistributionChannel            string      `json:"DistributionChannel"`
			Division                       string      `json:"Division"`
			SalesOffice                    string      `json:"SalesOffice"`
			SalesGroup                     string      `json:"SalesGroup"`
			SoldToParty                    string      `json:"SoldToParty"`
			ShipToParty                    string      `json:"ShipToParty"`
			BillToParty                    string      `json:"BillToParty"`
			PayerParty                     string      `json:"PayerParty"`
			ContactPerson                  string      `json:"ContactPerson"`
			ReferenceServiceOrder          string      `json:"ReferenceServiceOrder"`
			ServiceConfirmationIsFinal     string      `json:"ServiceConfirmationIsFinal"`
			TransactionCurrency            string      `json:"TransactionCurrency"`
			RespyMgmtServiceTeam           string      `json:"RespyMgmtServiceTeam"`
			RespyMgmtServiceTeamName       string      `json:"RespyMgmtServiceTeamName"`
			ServiceOrganization            string      `json:"ServiceOrganization"`
			ToPersonResponsible struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PersonResponsible"`
			ToReferenceObject struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ReferenceObject"`
			ToItem                         struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
