package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-service-confirmation-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	ServiceConfirmation:            data.ServiceConfirmation,
	ServiceConfirmationType:        data.ServiceConfirmationType,
	ServiceConfirmationDescription: data.ServiceConfirmationDescription,
	ServiceObjectType:              data.ServiceObjectType,
	Language:                       data.Language,
	ServiceDocumentPriority:        data.ServiceDocumentPriority,
	RequestedServiceStartDateTime:  data.RequestedServiceStartDateTime,
	RequestedServiceEndDateTime:    data.RequestedServiceEndDateTime,
	PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
	CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
	ServiceConfirmationIsCompleted: data.ServiceConfirmationIsCompleted,
	ServiceConfirmationIsCanceled:  data.ServiceConfirmationIsCanceled,
	SalesOrganization:              data.SalesOrganization,
	DistributionChannel:            data.DistributionChannel,
	Division:                       data.Division,
	SalesOffice:                    data.SalesOffice,
	SalesGroup:                     data.SalesGroup,
	SoldToParty:                    data.SoldToParty,
	ShipToParty:                    data.ShipToParty,
	BillToParty:                    data.BillToParty,
	PayerParty:                     data.PayerParty,
	ContactPerson:                  data.ContactPerson,
	ReferenceServiceOrder:          data.ReferenceServiceOrder,
	ServiceConfirmationIsFinal:     data.ServiceConfirmationIsFinal,
	TransactionCurrency:            data.TransactionCurrency,
	RespyMgmtServiceTeam:           data.RespyMgmtServiceTeam,
	RespyMgmtServiceTeamName:       data.RespyMgmtServiceTeamName,
	ServiceOrganization:            data.ServiceOrganization,
	ToPersonResponsible:            data.ToPersonResponsible.Deferred.URI,
	ToReferenceObject:              data.ToReferenceObject.Deferred.URI,
	ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToPersonResponsible(raw []byte, l *logger.Logger) ([]ToPersonResponsible, error) {
	pm := &responses.ToPersonResponsible{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPersonResponsible. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toPersonResponsible := make([]ToPersonResponsible, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPersonResponsible = append(toPersonResponsible, ToPersonResponsible{
	ServiceConfirmation:          data.ServiceConfirmation,
	PersonResponsible:            data.PersonResponsible,
	CustMgmtPartnerIsMainPartner: data.CustMgmtPartnerIsMainPartner,
		})
	}

	return toPersonResponsible, nil
}

func ConvertToToReferenceObject(raw []byte, l *logger.Logger) ([]ToReferenceObject, error) {
	pm := &responses.ToReferenceObject{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToReferenceObject. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toReferenceObject := make([]ToReferenceObject, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toReferenceObject = append(toReferenceObject, ToReferenceObject{
	ServiceConfirmation:          data.ServiceConfirmation,
	ServiceReferenceEquipment:    data.ServiceReferenceEquipment,
	ServiceRefFunctionalLocation: data.ServiceRefFunctionalLocation,
	SrvcRefObjIsMainObject:       data.SrvcRefObjIsMainObject,
		})
	}

	return toReferenceObject, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	ServiceConfirmation:           data.ServiceConfirmation,
	ServiceConfirmationItem:       data.ServiceConfirmationItem,
	ServiceConfItemDescription:    data.ServiceConfItemDescription,
	ServiceObjectType:             data.ServiceObjectType,
	ServiceDocumentItemObjectType: data.ServiceDocumentItemObjectType,
	Language:                      data.Language,
	Product:                       data.Product,
	Quantity:                      data.Quantity,
	QuantityUnit:                  data.QuantityUnit,
	ActualServiceDuration:         data.ActualServiceDuration,
	ActualServiceDurationUnit:     data.ActualServiceDurationUnit,
	ServiceConfItemCategory:       data.ServiceConfItemCategory,
	ServiceConfItemIsCompleted:    data.ServiceConfItemIsCompleted,
	ExecutingServiceEmployee:      data.ExecutingServiceEmployee,
	ServicePerformer:              data.ServicePerformer,
	BillableControl:               data.BillableControl,
	ReferenceServiceOrder:         data.ReferenceServiceOrder,
	ReferenceServiceOrderItem:     data.ReferenceServiceOrderItem,
	ParentServiceConfItem:         data.ParentServiceConfItem,
	ActualServiceStartDateTime:    data.ActualServiceStartDateTime,
	ActualServiceEndDateTime:      data.ActualServiceEndDateTime,
	ServicesRenderedDate:          data.ServicesRenderedDate,
	TimeSheetOvertimeCategory:     data.TimeSheetOvertimeCategory,
	RespyMgmtServiceTeam:          data.RespyMgmtServiceTeam,
	RespyMgmtServiceTeamName:      data.RespyMgmtServiceTeamName,
	ToItemPricingElement:          data.ToItemPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
	ServiceConfirmation:     data.ServiceConfirmation,
	ServiceConfirmationItem: data.ServiceConfirmationItem,
	PricingProcedureStep:    data.PricingProcedureStep,
	PricingProcedureCounter: data.PricingProcedureCounter,
	ConditionType:           data.ConditionType,
	ConditionRateValue:      data.ConditionRateValue,
	ConditionCurrency:       data.ConditionCurrency,
	ConditionQuantity:       data.ConditionQuantity,
	ConditionQuantityUnit:   data.ConditionQuantityUnit,
		})
	}

	return toItemPricingElement, nil
}
