package sap_api_output_formatter

type Header struct {
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
	ToPersonResponsible            string      `json:"to_PersonResponsible"`
	ToReferenceObject              string      `json:"to_ReferenceObject"`
	ToItem                         string      `json:"to_Item"`
}

type ToPersonResponsible struct {
	ServiceConfirmation          string `json:"ServiceConfirmation"`
	PersonResponsible            string `json:"PersonResponsible"`
	CustMgmtPartnerIsMainPartner bool   `json:"CustMgmtPartnerIsMainPartner"`
}

type ToReferenceObject struct {
	ServiceConfirmation          string `json:"ServiceConfirmation"`
	ServiceReferenceEquipment    string `json:"ServiceReferenceEquipment"`
	ServiceRefFunctionalLocation string `json:"ServiceRefFunctionalLocation"`
	SrvcRefObjIsMainObject       bool   `json:"SrvcRefObjIsMainObject"`
}

type ToItem struct {
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
	ToItemPricingElement          string      `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	ServiceConfirmation     string `json:"ServiceConfirmation"`
	ServiceConfirmationItem string `json:"ServiceConfirmationItem"`
	PricingProcedureStep    string `json:"PricingProcedureStep"`
	PricingProcedureCounter string `json:"PricingProcedureCounter"`
	ConditionType           string `json:"ConditionType"`
	ConditionRateValue      string `json:"ConditionRateValue"`
	ConditionCurrency       string `json:"ConditionCurrency"`
	ConditionQuantity       string `json:"ConditionQuantity"`
	ConditionQuantityUnit   string `json:"ConditionQuantityUnit"`
}
