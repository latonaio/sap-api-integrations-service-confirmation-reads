package responses

type ToPersonResponsible struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ServiceConfirmation          string `json:"ServiceConfirmation"`
			PersonResponsible            string `json:"PersonResponsible"`
			CustMgmtPartnerIsMainPartner bool   `json:"CustMgmtPartnerIsMainPartner"`
		} `json:"results"`
	} `json:"d"`
}
