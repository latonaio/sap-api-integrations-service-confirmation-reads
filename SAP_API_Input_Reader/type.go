package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	ServiceConfirmation struct {
		ServiceConfirmation       string      `json:"document_no"`
		Plant                     string      `json:"deliver_to"`
		OrderQuantity             string      `json:"quantity"`
		PickedQuantity            string      `json:"picked_quantity"`
		NetPriceAmount            string      `json:"price"`
	    Batch                     string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                         string `json:"api_schema"`
	Accepter                        []string `json:"accepter"`
	MaterialCode                      string `json:"material_code"`
	Supplier                          string `json:"plant/supplier"`
	Stock                             string `json:"stock"`
	ServiceConfirmationType           string `json:"document_type"`
	ServiceConfirmationNo             string `json:"document_no"`
	ScheduleLineDeliveryDate          string `json:"planned_date"`
	ValidatedDate                     string `json:"validated_date"`
    Deleted                           bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string `json:"connection_key"`
	Result              bool   `json:"result"`
	RedisKey            string `json:"redis_key"`
	Filepath            string `json:"filepath"`
	ServiceConfirmation struct {
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
		PersonResponsible struct {
			PersonResponsible            string      `json:"PersonResponsible"`
			CustMgmtPartnerIsMainPartner bool        `json:"CustMgmtPartnerIsMainPartner"`
		} `json:"PersonResponsible"`
		ReferenceObject struct {
			ServiceReferenceEquipment    string      `json:"ServiceReferenceEquipment"`
			ServiceRefFunctionalLocation string      `json:"ServiceRefFunctionalLocation"`
			SrvcRefObjIsMainObject       bool        `json:"SrvcRefObjIsMainObject"`
		} `json:"ReferenceObject"`
		ServiceConfirmationItem struct {
			ServiceConfirmationItem       string `json:"ServiceConfirmationItem"`
			ServiceConfItemDescription    string `json:"ServiceConfItemDescription"`
			ServiceObjectType             string `json:"ServiceObjectType"`
			ServiceDocumentItemObjectType string `json:"ServiceDocumentItemObjectType"`
			Language                      string `json:"Language"`
			Product                       string `json:"Product"`
			Quantity                      string `json:"Quantity"`
			QuantityUnit                  string `json:"QuantityUnit"`
			ActualServiceDuration         string `json:"ActualServiceDuration"`
			ActualServiceDurationUnit     string `json:"ActualServiceDurationUnit"`
			ServiceConfItemCategory       string `json:"ServiceConfItemCategory"`
			ServiceConfItemIsCompleted    string `json:"ServiceConfItemIsCompleted"`
			ExecutingServiceEmployee      string `json:"ExecutingServiceEmployee"`
			ServicePerformer              string `json:"ServicePerformer"`
			BillableControl               string `json:"BillableControl"`
			ReferenceServiceOrder         string `json:"ReferenceServiceOrder"`
			ReferenceServiceOrderItem     string `json:"ReferenceServiceOrderItem"`
			ParentServiceConfItem         string `json:"ParentServiceConfItem"`
			ActualServiceStartDateTime    string `json:"ActualServiceStartDateTime"`
			ActualServiceEndDateTime      string `json:"ActualServiceEndDateTime"`
			ServicesRenderedDate          string `json:"ServicesRenderedDate"`
			TimeSheetOvertimeCategory     string `json:"TimeSheetOvertimeCategory"`
			RespyMgmtServiceTeam          string `json:"RespyMgmtServiceTeam"`
			RespyMgmtServiceTeamName      string `json:"RespyMgmtServiceTeamName"`
			ItemPricingElement            struct {
				ServiceConfirmation     string `json:"ServiceConfirmation"`
				ServiceConfirmationItem string `json:"ServiceConfirmationItem"`
				PricingProcedureStep    string `json:"PricingProcedureStep"`
				PricingProcedureCounter string `json:"PricingProcedureCounter"`
				ConditionType           string `json:"ConditionType"`
				ConditionRateValue      string `json:"ConditionRateValue"`
				ConditionCurrency       string `json:"ConditionCurrency"`
				ConditionQuantity       string `json:"ConditionQuantity"`
				ConditionQuantityUnit   string `json:"ConditionQuantityUnit"`
			} `json:"ItemPricingElement"`
		} `json:"ServiceConfirmationItem"`
	} `json:"ServiceConfirmation"`
	APISchema             string   `json:"api_schema"`
	Accepter              []string `json:"accepter"`
	ServiceConfirmationNo string   `json:"service_confirmation"`
	Deleted               bool     `json:"deleted"`
}
